package main

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/google/uuid"

	proto "github.com/ryarnyah/mood-tracker/proto"
)

type moodServer struct {
	db sql.DB
}

func (m * moodServer) AddEntry(ctx context.Context, request *proto.AddEntryRequest) (*proto.AddEntryResponse, error) {
	// First find by accessCode

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	updateEntry, err := tx.Prepare("UPDATE ENTRY SET RECORD = ?, COMMENT = ? WHERE ACCESS_CODE = ? AND REF_ACCESS_CODE = ? AND RECORD IS NULL AND COMMENT IS NULL")
	if err != nil {
		return nil, err
	}
	_, err = updateEntry.Exec(request.GetEntry().GetRecord(), request.GetEntry().GetComment(), request.GetAccessCode(), request.GetRefAccessCode())
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &proto.AddEntryResponse{}, nil
}
func (m * moodServer) ListEntries(ctx context.Context, request *proto.ListEntryRequest) (*proto.ListEntryResponse, error) {
	return nil, nil
}
func (m * moodServer) CreateMood(ctx context.Context, request *proto.CreateMoodRequest) (*proto.CreateMoodResponse,  error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	moodUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	// Create mood entry
	moodStmt, err := tx.Prepare("INSERT INTO MOOD (ACCESS_CODE, TITLE, CONTENT, NUMBER_OF_RECORDS_NEEDED) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer moodStmt.Close()

	_, err = moodStmt.Exec(moodUUID.String(), request.GetTitle(), request.GetContent(), request.GetNumberOfRecordsNeeded())
	if err != nil {
		return nil, err
	}

	// Create entry entries
	entryStmt, err := tx.Prepare("INSERT INTO ENTRY (ACCESS_CODE, REF_ACCESS_CODE) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}

	recordsAccessCodes := make([]string, request.GetNumberOfRecordsNeeded())
	var i uint32
	for i = 0; i < request.GetNumberOfRecordsNeeded(); i++ {
		entryUUID, err := uuid.NewUUID()
		if err != nil {
			return nil, err
		}
		_, err = entryStmt.Exec(moodUUID.String(), entryUUID.String())
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
		AccessCode: moodUUID.String(),
		RecordsAccessCodes: recordsAccessCodes,
	}, nil
}

func main() {
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"sqlite3", driver,
	)
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange{
		panic(err)
	}

}
