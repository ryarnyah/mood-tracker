package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang/glog"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rakyll/statik/fs"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	proto "github.com/ryarnyah/mood-tracker/proto"
	"github.com/ryarnyah/mood-tracker/version"

	_ "github.com/ryarnyah/mood-tracker/statik"
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

type moodServer struct {
	db *sql.DB
}

func (m *moodServer) GetMoodFromEntry(ctx context.Context, request *proto.GetMoodFromEntryRequest) (*proto.GetMoodFromEntryResponse, error) {
	var title string
	var content string

	err := m.db.QueryRowContext(ctx, `SELECT MOOD.TITLE, MOOD.CONTENT
          FROM MOOD JOIN ENTRY ON MOOD.MOOD_ID = ENTRY.MOOD_ID
          LEFT JOIN RECORD ON ENTRY.ENTRY_ID = RECORD.ENTRY_ID
          WHERE ENTRY.MOOD_ID = ? AND ENTRY.ENTRY_ACCESS_CODE = ? AND RECORD.ENTRY_ID IS NULL`, request.GetMoodId(), request.GetEntryAccessCode()).Scan(&title, &content)

	if err != nil {
		return nil, err
	}

	return &proto.GetMoodFromEntryResponse{
		Title:   title,
		Content: content,
	}, nil
}

func (m *moodServer) AddEntry(ctx context.Context, request *proto.AddEntryRequest) (*proto.AddEntryResponse, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var entryID int
	err = tx.QueryRowContext(ctx, `SELECT ENTRY.ENTRY_ID
          FROM ENTRY LEFT JOIN RECORD ON ENTRY.ENTRY_ID = RECORD.ENTRY_ID
          WHERE ENTRY.MOOD_ID = ? AND ENTRY.ENTRY_ACCESS_CODE = ? AND RECORD.ENTRY_ID IS NULL`, request.GetMoodId(), request.GetEntryAccessCode()).Scan(&entryID)
	if err == sql.ErrNoRows {
		return nil, errors.New("access-code or mood-id is invalid or expired")
	} else if err != nil {
		return nil, err
	}

	updateEntry, err := tx.Prepare("INSERT INTO RECORD (ENTRY_ID, RECORD, COMMENT) VALUES (?, ?, ?)")
	if err != nil {
		return nil, err
	}
	_, err = updateEntry.Exec(entryID, request.GetEntry().GetRecord(), request.GetEntry().GetComment())
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &proto.AddEntryResponse{}, nil
}
func (m *moodServer) GetMood(ctx context.Context, request *proto.GetMoodRequest) (*proto.GetMoodResponse, error) {
	rows, err := m.db.Query(`SELECT RECORD.RECORD, RECORD.COMMENT
          FROM RECORD JOIN ENTRY ON RECORD.ENTRY_ID = ENTRY.ENTRY_ID
          JOIN MOOD ON MOOD.MOOD_ID = ENTRY.MOOD_ID
          WHERE MOOD.MOOD_ID = ? AND MOOD.MOOD_ACCESS_CODE = ?`, request.GetMoodId(), request.GetMoodAccessCode())
	if err != nil {
		return nil, err
	}
	entries := make([]*proto.Entry, 0)
	for rows.Next() {
		var record uint32
		var comment string
		err = rows.Scan(&record, &comment)
		if err != nil {
			return nil, err
		}
		entries = append(entries, &proto.Entry{
			Record:  record,
			Comment: comment,
		})
	}
	return &proto.GetMoodResponse{
		Entries: entries,
	}, nil
}

func (m *moodServer) CreateMood(ctx context.Context, request *proto.CreateMoodRequest) (*proto.CreateMoodResponse, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	moodUUID := uuid.New()

	// Create mood entry
	moodStmt, err := tx.Prepare("INSERT INTO MOOD (MOOD_ACCESS_CODE, TITLE, CONTENT) VALUES (?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer moodStmt.Close()

	r, err := moodStmt.Exec(moodUUID.String(), request.GetTitle(), request.GetContent())
	if err != nil {
		return nil, err
	}

	moodID, err := r.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Create entry entries
	entryStmt, err := tx.Prepare("INSERT INTO ENTRY (ENTRY_ACCESS_CODE, MOOD_ID) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}

	recordsAccessCodes := []string{}
	var i uint32
	for i = 0; i < request.GetNumberOfRecordsNeeded(); i++ {
		entryUUID := uuid.New()
		_, err = entryStmt.Exec(entryUUID.String(), moodID)
		if err != nil {
			return nil, err
		}
		recordsAccessCodes = append(recordsAccessCodes, entryUUID.String())
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &proto.CreateMoodResponse{
		MoodId:             moodID,
		MoodAccessCode:     moodUUID.String(),
		EntriesAccessCodes: recordsAccessCodes,
	}, nil
}

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
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"sqlite3", driver,
	)
	if err != nil {
		glog.Fatalf("unable to migrate db %v", err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		glog.Fatalf("unable to migrate db %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	proto.RegisterMoodServer(grpcServer, &moodServer{
		db: db,
	})

	wrappedServer := grpcweb.WrapServer(grpcServer)
	handler := http.StripPrefix("/grpc/", wrappedServer)

	statikFS, err := fs.New()
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
