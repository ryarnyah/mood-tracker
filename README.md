# Mood-Tracker [![Build Status](https://travis-ci.org/ryarnyah/mood-tracker.svg?branch=master)](https://travis-ci.org/ryarnyah/mood-tracker) #

Some web-grpc mood tracker to track your team mood anonymously :D

## Installation ##

#### Binaries ####

- **linux** [amd64](https://github.com/ryarnyah/mood-tracker/releases/download/0.0.5/mood-tracker-linux-amd64) [386](https://github.com/ryarnyah/mood-tracker/releases/download/0.0.5/mood-tracker-linux-386)

```bash
sudo curl -L https://github.com/ryarnyah/mood-tracker/releases/download/0.0.5/mood-tracker-linux-amd64 -o /usr/local/bin/mood-tracker && sudo chmod +x /usr/local/bin/mood-tracker
```

#### Via Go ####

```bash
$ go get github.com/ryarnyah/mood-tracker
```

#### From Source ####

```bash
$ mkdir -p $GOPATH/src/github.com/ryarnyah
$ git clone https://github.com/ryarnyah/mood-tracker $GOPATH/src/github.com/ryarnyah/mood-tracker
$ cd !$
$ make
```

#### Running with Docker ####

```bash
docker run ryarnyah/mood-tracker-linux-amd64:0.0.5 <option>
```

## Usage ##

```bash

___  ___                _      _____              _
|  \/  |               | |    |_   _|            | |
| .  . | ___   ___   __| |______| |_ __ __ _  ___| | _____ _ __
| |\/| |/ _ \ / _ \ / _` |______| | '__/ _` |/ __| |/ / _ \ '__|
| |  | | (_) | (_) | (_| |      | | | | (_| | (__|   <  __/ |
\_|  |_/\___/ \___/ \__,_|      \_/_|  \__,_|\___|_|\_\___|_|

 Get your mood ready.
 Version: 
 Build: 
  -alsologtostderr
    	log to standard error as well as files
  -enable-tls
    	Use TLS - required for HTTP2.
  -host string
    	Server host. (default "localhost:8090")
  -log_backtrace_at value
    	when logging hits line file:N, emit a stack trace
  -log_dir string
    	If non-empty, write log files in this directory
  -logtostderr
    	log to standard error instead of files
  -profiling-enable
    	Enable profiling
  -profiling-host string
    	HTTP profiling host:port (default "localhost:6060")
  -smt-sender-email string
    	SMTP email from (default "root@localhost")
  -smtp-external-url string
    	Public URL or mood tracker (default "localhost:8090")
  -smtp-host string
    	SMTP Server host
  -smtp-password string
    	SMTP password (if needed)
  -smtp-port int
    	SMTP Server host port (default 587)
  -smtp-username string
    	SMTP username (if needed)
  -stderrthreshold value
    	logs at or above this threshold go to stderr
  -tls-cert-file string
    	Path to the CRT/PEM file. (default "server.crt")
  -tls-key-file string
    	Path to the private key file. (default "server.key")
  -v value
    	log level for V logs
  -version
    	Print version
  -vmodule value
    	comma-separated list of pattern=N settings for file-filtered logging
```
