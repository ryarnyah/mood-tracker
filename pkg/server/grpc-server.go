package server

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	proto "github.com/ryarnyah/mood-tracker/proto"
)

type moodServer struct {
	db *sql.DB
}

func NewMoodServer(db *sql.DB) proto.MoodServer {
	return &moodServer{
		db: db,
	}
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
	defer updateEntry.Close()
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
	defer rows.Close()
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
	statRows, err := m.db.Query(`SELECT ifnull(RECORD.RECORD, 0), COUNT(1)
          FROM ENTRY LEFT JOIN RECORD ON RECORD.ENTRY_ID = ENTRY.ENTRY_ID
          JOIN MOOD ON MOOD.MOOD_ID = ENTRY.MOOD_ID
          WHERE MOOD.MOOD_ID = ? AND MOOD.MOOD_ACCESS_CODE = ?
          GROUP BY RECORD.RECORD`, request.GetMoodId(), request.GetMoodAccessCode())
	if err != nil {
		return nil, err
	}
	defer statRows.Close()

	stats := make(map[uint32]int64)
	for statRows.Next() {
		var record uint32
		var count int64
		err = statRows.Scan(&record, &count)
		if err != nil {
			return nil, err
		}
		stats[record] = count
	}

	var title string
	var content string

	err = m.db.QueryRowContext(ctx, `SELECT MOOD.TITLE, MOOD.CONTENT
          FROM MOOD
          WHERE MOOD.MOOD_ID = ? AND MOOD.MOOD_ACCESS_CODE = ?`, request.GetMoodId(), request.GetMoodAccessCode()).Scan(&title, &content)

	if err != nil {
		return nil, err
	}

	return &proto.GetMoodResponse{
		Title:   title,
		Content: content,
		Entries: entries,
		Stats:   stats,
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
	defer entryStmt.Close()

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
