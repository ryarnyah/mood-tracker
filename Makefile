# Set an output prefix, which is the local directory if not specified
PREFIX?=$(shell pwd)

# Setup name variables for the package/tool
BASE_REPOSITORY := ryarnyah
BINARIES := mood-tracker
BASE_PKG := github.com/$(BASE_REPOSITORY)
BASE_BINARIES := ./cmd

# Set any default go build tags
BUILDTAGS := 

# Set the build dir, where built cross-compiled binaries will be output
BASE_BUILD_DIR := build
BUILDDIR := ${PREFIX}/$(BASE_BUILD_DIR)

# Populate version variables
# Add to compile time flags
VERSION := $(shell cat VERSION.txt)
GITCOMMIT := $(shell git rev-parse --short HEAD)
GITUNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)
ifneq ($(GITUNTRACKEDCHANGES),)
	GITCOMMIT := $(GITCOMMIT)-dirty
endif
CTIMEVAR=-X $(BASE_PKG)/$(1)/version.GITCOMMIT=$(GITCOMMIT) -X $(BASE_PKG)/$(1)/version.VERSION=$(VERSION)
GO_LDFLAGS=-ldflags "-w $(call CTIMEVAR,$(1))"
GO_LDFLAGS_STATIC=-ldflags "-w $(call CTIMEVAR,$(1)) -extldflags -static"

# List the GOOS and GOARCH to build
GOOSARCHES = linux/amd64 linux/386

all: clean build fmt lint test vet ## Runs a clean, build, fmt, lint, test, staticcheck, vet (staticcheck disable)

.PHONY: build
build: $(BINARIES) ## Builds a dynamic executable or package

$(BINARIES): VERSION.txt protoc yarn-build statik
	@echo "+ $@"
	GO111MODULE=on CGO_ENABLED=1 go build -tags "$(BUILDTAGS)" $(call GO_LDFLAGS,$@) -o $@ $(BASE_BINARIES)/$@

.PHONY: static
static: ## Builds a static executable
	@echo "+ $@"
	$(foreach BINARY,$(BINARIES),GO111MODULE=on CGO_ENABLED=1 go build -tags "$(BUILDTAGS) static_build" $(call GO_LDFLAGS_STATIC $(BINARY)) -o $(BINARY) $(BASE_BINARIES)/$(BINARY))

.PHONY: fmt
fmt: ## Verifies all files have men `gofmt`ed
	@echo "+ $@"
	@gofmt -s -l . | grep -v '.pb.go:' | grep -v vendor | tee /dev/stderr

.PHONY: lint
lint: ## Verifies `golint` passes
	@echo "+ $@"
	@golint ./... | grep -v '.pb.go:' | grep -v vendor | tee /dev/stderr

.PHONY: test
test: ## Runs the go tests
	@echo "+ $@"
	@go test -v -tags "$(BUILDTAGS) cgo" $(shell go list ./... | grep -v vendor)

.PHONY: vet
vet: ## Verifies `go vet` passes
	@echo "+ $@"
	@go vet $(shell go list ./... | grep -v vendor) | grep -v '.pb.go:' | tee /dev/stderr

.PHONY: staticcheck
staticcheck: ## Verifies `staticcheck` passes
	@echo "+ $@"
	@staticcheck $(shell go list ./... | grep -v vendor) | grep -v '.pb.go:' | tee /dev/stderr

.PHONY: cover
cover: ## Runs go test with coverage
	@echo "" > coverage.txt
	@for d in $(shell go list ./... | grep -v vendor); do \
		go test -race -coverprofile=profile.out -covermode=atomic "$$d"; \
		if [ -f profile.out ]; then \
			cat profile.out >> coverage.txt; \
			rm profile.out; \
		fi; \
	done;

define buildpretty
mkdir -p $(BUILDDIR)/$(1)/$(2);
GOOS=$(1) GOARCH=$(2) GO111MODULE=on CGO_ENABLED=1 go build \
	 -o $(BUILDDIR)/$(1)/$(2)/$(3) \
	 -a -tags "$(BUILDTAGS) static_build netgo" \
	 -installsuffix netgo $(call GO_LDFLAGS_STATIC,$3) $(BASE_BINARIES)/$(3);
md5sum $(BUILDDIR)/$(1)/$(2)/$(3) > $(BUILDDIR)/$(1)/$(2)/$(3).md5;
sha256sum $(BUILDDIR)/$(1)/$(2)/$(3) > $(BUILDDIR)/$(1)/$(2)/$(3).sha256;
endef

.PHONY: cross
cross: VERSION.txt protoc yarn-build statik ## Builds the cross-compiled binaries, creating a clean directory structure (eg. GOOS/GOARCH/binary)
	@echo "+ $@"
	$(foreach BINARY,$(BINARIES), $(foreach GOOSARCH,$(GOOSARCHES), $(call buildpretty,$(subst /,,$(dir $(GOOSARCH))),$(notdir $(GOOSARCH)),$(BINARY))))

define buildrelease
GOOS=$(1) GOARCH=$(2) GO111MODULE=on CGO_ENABLED=1 go build \
	 -o $(BUILDDIR)/$(3)-$(1)-$(2) \
	 -a -tags "$(BUILDTAGS) static_build netgo" \
	 -installsuffix netgo $(call GO_LDFLAGS_STATIC,$3) $(BASE_BINARIES)/$(3);
md5sum $(BUILDDIR)/$(3)-$(1)-$(2) > $(BUILDDIR)/$(3)-$(1)-$(2).md5;
sha256sum $(BUILDDIR)/$(3)-$(1)-$(2) > $(BUILDDIR)/$(3)-$(1)-$(2).sha256;
endef

.PHONY: release
release: VERSION.txt protoc yarn-build statik ## Builds the cross-compiled binaries, naming them in such a way for release (eg. binary-GOOS-GOARCH)
	@echo "+ $@"
	$(foreach BINARY,$(BINARIES), $(foreach GOOSARCH,$(GOOSARCHES), $(call buildrelease,$(subst /,,$(dir $(GOOSARCH))),$(notdir $(GOOSARCH)),$(BINARY))))

.PHONY: protoc
protoc: yarn
	protoc -I proto/ \
	--proto_path=${GOPATH}/src \
	--go_out=plugins=grpc:proto \
	--plugin=protoc-gen-ts=./mood-tracker-client/node_modules/.bin/protoc-gen-ts \
	--ts_out=service=grpc-web:mood-tracker-client/src/proto \
	--js_out=import_style=commonjs,binary:mood-tracker-client/src/proto \
	--govalidators_out=proto \
	proto/mood.proto

	grep -v 'validators_validator_pb' ./mood-tracker-client/src/proto/mood_pb.d.ts > ./mood-tracker-client/src/proto/mood_pb.d.ts.tmp
	mv ./mood-tracker-client/src/proto/mood_pb.d.ts.tmp ./mood-tracker-client/src/proto/mood_pb.d.ts
	grep -v 'validators_validator_pb' ./mood-tracker-client/src/proto/mood_pb.js > ./mood-tracker-client/src/proto/mood_pb.js.tmp
	mv ./mood-tracker-client/src/proto/mood_pb.js.tmp ./mood-tracker-client/src/proto/mood_pb.js
	printf '/* eslint-disable */\n//@ts-nocheck\n' | cat - ./mood-tracker-client/src/proto/mood_pb.js > ./mood-tracker-client/src/proto/mood_pb.js.tmp
	mv ./mood-tracker-client/src/proto/mood_pb.js.tmp ./mood-tracker-client/src/proto/mood_pb.js


.PHONE: install-yarn
install-yarn:
	npm install -g yarn

.PHONY: yarn
yarn: install-yarn
	yarn --cwd ./mood-tracker-client install

.PHONY: yarn-build
yarn-build: yarn
	yarn --cwd ./mood-tracker-client build

.PHONY: statik
statik:
	statik -f -ns public -p statik -src=./mood-tracker-client/dist
	statik -f -ns migrations -p statik_migrations -src=./migrations

.PHONY: bump-version
BUMP := patch
bump-version: ## Bump the version in the version file. Set BUMP to [ patch | major | minor ]
	$(eval NEW_VERSION = $(shell sembump --kind $(BUMP) $(VERSION)))
	@echo "Bumping VERSION.txt from $(VERSION) to $(NEW_VERSION)"
	echo $(NEW_VERSION) > VERSION.txt
	@echo "Updating links to download binaries in README.md"
	sed -e s/$(VERSION)/$(NEW_VERSION)/g README.md > README.md.tmp
	mv README.md.tmp README.md
	git add VERSION.txt README.md
	git commit -vsm "Bump version to $(NEW_VERSION)"
	@echo "Run make tag to create and push the tag for new version $(NEW_VERSION)"

.PHONY: tag
tag: ## Create a new git tag to prepare to build a release
	git tag -sa $(VERSION) -m "$(VERSION)"
	@echo "Run git push origin $(VERSION) to push your new tag to GitHub and trigger a travis build."

.PHONY: AUTHORS
AUTHORS:
	@$(file >$@,# This file lists all individuals having contributed content to the repository.)
	@$(file >>$@,# For how it is generated, see `make AUTHORS`.)
	@echo "$(shell git log --format='\n%aN <%aE>' | LC_ALL=C.UTF-8 sort -uf)" > $@
	git add AUTHORS
	git commit -vsm "Updated AUTHORS"

.PHONY: clean
clean: ## Cleanup any build binaries or packages
	@echo "+ $@"
	$(foreach BINARY,$(BINARIES), $(RM) $(BINARY))
	$(RM) -r $(BUILDDIR)
	$(RM) -r mood-tracker-client/dist

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: docker
docker: ## Build docker images for GOOS/GOARCH/BINARIES
	$(foreach BINARY,$(BINARIES), $(foreach GOOSARCH,$(GOOSARCHES), $(call builddocker,$(subst /,,$(dir $(GOOSARCH))),$(notdir $(GOOSARCH)),$(BINARY))))

.PHONY: docker-deploy
docker-deploy: ## Deploy docker images to hub
	$(foreach BINARY,$(BINARIES), $(foreach GOOSARCH,$(GOOSARCHES), $(call deploydocker,$(subst /,,$(dir $(GOOSARCH))),$(notdir $(GOOSARCH)),$(BINARY))))

define builddocker
docker build --build-arg BINARY_PATH=$(BASE_BUILD_DIR)/$(3)-$(1)-$(2) --build-arg BINARY_NAME=$(3) -t $(BASE_REPOSITORY)/$(3)-$(1)-$(2):$(VERSION) -f deploy/Dockerfile .;
endef

define deploydocker
docker push $(BASE_REPOSITORY)/$(3)-$(1)-$(2):$(VERSION);
endef

.PHONY: dev-dependencies
dev-dependencies: install-yarn ## Install all dev dependencies
	@GO111MODULE=off go get -v -u github.com/jessfraz/junk/sembump
	@GO111MODULE=off go get -v -u honnef.co/go/tools/cmd/staticcheck
	@GO111MODULE=off go get -v -u golang.org/x/lint/golint
	@GO111MODULE=off go get -v -u github.com/golang/protobuf/protoc-gen-go
	@GO111MODULE=off go get -v -u github.com/mwitkow/go-proto-validators
	@GO111MODULE=off go get -v -u github.com/rakyll/statik
	@GO111MODULE=off go get -v -u github.com/mwitkow/go-proto-validators/protoc-gen-govalidators