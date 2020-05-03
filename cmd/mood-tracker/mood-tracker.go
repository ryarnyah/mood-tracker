package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rakyll/statik/fs"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_ratelimit "github.com/ryarnyah/mood-tracker/pkg/grpc_ratelimit"
	"github.com/ryarnyah/mood-tracker/pkg/server"
	proto "github.com/ryarnyah/mood-tracker/proto"
	"github.com/ryarnyah/mood-tracker/version"

	migrate_statik "github.com/ryarnyah/mood-tracker/pkg/migrate/statik"
	_ "github.com/ryarnyah/mood-tracker/statik"
	_ "github.com/ryarnyah/mood-tracker/statik_migrations"
)

var (
	enableTLS       = flag.Bool("enable-tls", false, "Use TLS - required for HTTP2.")
	tlsCertFilePath = flag.String("tls-cert-file", "server.crt", "Path to the CRT/PEM file.")
	tlsKeyFilePath  = flag.String("tls-key-file", "server.key", "Path to the private key file.")
	host            = flag.String("host", "localhost:8090", "Server host.")

	profiling     = flag.Bool("profiling-enable", false, "Enable profiling")
	profilingHost = flag.String("profiling-host", "localhost:6060", "HTTP profiling host:port")
	v             = flag.Bool("version", false, "Print version")
)

const (
	// BANNER for usage.
	BANNER = `
 Get your mood ready.
 Version: %s
 Build: %s
`
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, BANNER, version.VERSION, version.GITCOMMIT)
		flag.PrintDefaults()
	}
	flag.Parse()

	if *v {
		fmt.Printf(BANNER, version.VERSION, version.GITCOMMIT)
		return
	}

	if *profiling {
		go func() {
			logrus.Error(http.ListenAndServe(*profilingHost, http.DefaultServeMux))
		}()
	}

	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		glog.Fatalf("unable to open db %v", err)
	}
	defer db.Close()

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		glog.Fatalf("unable to migrate db %v", err)
	}
	statikMigrationFS, err := fs.NewWithNamespace("migrations")
	if err != nil {
		glog.Fatal(err)
	}
	d, err := migrate_statik.WithInstance(statikMigrationFS, "/")
	if err != nil {
		glog.Fatal(err)
	}
	m, err := migrate.NewWithInstance(
		"statik",
		d,
		"sqlite3", driver,
	)
	if err != nil {
		glog.Fatalf("unable to migrate db %v", err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		glog.Fatalf("unable to migrate db %v", err)
	}

	limiter := grpc_ratelimit.NewGRPCLimiter(30, 10)

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
			grpc_ratelimit.StreamServerInterceptor(limiter),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
			grpc_ratelimit.UnaryServerInterceptor(limiter),
		)),
	)
	proto.RegisterMoodServer(grpcServer, server.NewMoodServer(db))

	wrappedServer := grpcweb.WrapServer(grpcServer)
	handler := http.StripPrefix("/grpc/", wrappedServer)

	statikFS, err := fs.NewWithNamespace("public")
	if err != nil {
		glog.Fatal(err)
	}

	staticHandler := http.FileServer(statikFS)

	r := mux.NewRouter()
	r.PathPrefix("/grpc/").Handler(handler)

	/* Tips to default on index.html for vuejs app */
	err = fs.Walk(statikFS, "/", func(path string, info os.FileInfo, err error) error {
		r.Path(path).Handler(staticHandler)
		return nil
	})
	if err != nil {
		glog.Fatal(err)
	}

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := statikFS.Open("/index.html")
		if err != nil {
			glog.Error(err)
			staticHandler.ServeHTTP(w, r)
			return
		}
		http.ServeContent(w, r, "index.html", time.Now(), f)
	})

	httpServer := http.Server{
		Addr:    *host,
		Handler: r,
	}

	if *enableTLS {
		if err := httpServer.ListenAndServeTLS(*tlsCertFilePath, *tlsKeyFilePath); err != nil {
			glog.Fatalf("failed starting http2 server: %v", err)
		}
	} else {
		if err := httpServer.ListenAndServe(); err != nil {
			glog.Fatalf("failed starting http server: %v", err)
		}
	}
}
