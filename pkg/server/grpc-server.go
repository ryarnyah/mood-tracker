package server

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/golang/protobuf/ptypes"
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

	now := time.Now()
	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, now.Location())

	err := m.db.QueryRowContext(ctx, `SELECT MOOD.TITLE, MOOD.CONTENT
          FROM MOOD JOIN ENTRY ON MOOD.MOOD_ID = ENTRY.MOOD_ID
          LEFT JOIN RECORD ON ENTRY.ENTRY_ID = RECORD.ENTRY_ID
          WHERE ENTRY.MOOD_ID = ? AND ENTRY.ENTRY_ACCESS_CODE = ? AND (RECORD.ENTRY_ID IS NULL OR RECORD.RECORD_DATETIME <> ?)`, request.GetMoodId(), request.GetEntryAccessCode(), today).Scan(&title, &content)

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

	now := time.Now()
	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, now.Location())

	var entryID int
	err = tx.QueryRowContext(ctx, `SELECT ENTRY.ENTRY_ID
          FROM ENTRY LEFT JOIN RECORD ON ENTRY.ENTRY_ID = RECORD.ENTRY_ID
          WHERE ENTRY.MOOD_ID = ? AND ENTRY.ENTRY_ACCESS_CODE = ? AND (RECORD.ENTRY_ID IS NULL OR RECORD.RECORD_DATETIME <> ?)`, request.GetMoodId(), request.GetEntryAccessCode(), today).Scan(&entryID)
	if err == sql.ErrNoRows {
		return nil, errors.New("access-code or mood-id is invalid or expired")
	} else if err != nil {
		return nil, err
	}

	updateEntry, err := tx.Prepare("INSERT INTO RECORD (ENTRY_ID, RECORD, COMMENT, RECORD_DATETIME) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer updateEntry.Close()
	_, err = updateEntry.Exec(entryID, request.GetEntry().GetRecord(), request.GetEntry().GetComment(), today)
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
	rows, err := m.db.Query(`SELECT RECORD.RECORD, RECORD.COMMENT, RECORD.RECORD_DATETIME
          FROM RECORD JOIN ENTRY ON RECORD.ENTRY_ID = ENTRY.ENTRY_ID
          JOIN MOOD ON MOOD.MOOD_ID = ENTRY.MOOD_ID
          WHERE MOOD.MOOD_ID = ? AND MOOD.MOOD_ACCESS_CODE = ?
          ORDER BY RECORD.RECORD_DATETIME DESC`, request.GetMoodId(), request.GetMoodAccessCode())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	entries := make([]*proto.EntryWithDate, 0)
	for rows.Next() {
		var day time.Time
		var record uint32
		var comment string
		err = rows.Scan(&record, &comment, &day)
		if err != nil {
			return nil, err
		}
		t, err := ptypes.TimestampProto(day)
		if err != nil {
			return nil, err
		}
		entries = append(entries, &proto.EntryWithDate{
			Record:      record,
			Comment:     comment,
			RecordEntry: t,
		})
	}
	statRows, err := m.db.Query(`SELECT RECORD.RECORD, RECORD.RECORD_DATETIME, COUNT(1)
          FROM ENTRY
          JOIN RECORD ON RECORD.ENTRY_ID = ENTRY.ENTRY_ID
          JOIN MOOD ON MOOD.MOOD_ID = ENTRY.MOOD_ID
          WHERE MOOD.MOOD_ID = ? AND MOOD.MOOD_ACCESS_CODE = ? AND RECORD.RECORD_DATETIME > date('now', '-7 day')
          GROUP BY RECORD.RECORD, RECORD.RECORD_DATETIME
          ORDER BY RECORD.RECORD, RECORD.RECORD_DATETIME`, request.GetMoodId(), request.GetMoodAccessCode())
	if err != nil {
		return nil, err
	}
	defer statRows.Close()

	stats := make(map[uint32]*proto.MoodStat)
	for statRows.Next() {
		var record uint32
		var day time.Time
		var count int64
		err = statRows.Scan(&record, &day, &count)
		if err != nil {
			return nil, err
		}
		if _, ok := stats[record]; !ok {
			stats[record] = &proto.MoodStat{
				Record: record,
			}
		}

		t, err := ptypes.TimestampProto(day)
		if err != nil {
			return nil, err
		}
		stats[record].RecordStats = append(stats[record].RecordStats, &proto.RecordStat{
			RecordEntry: t,
			Count:       count,
		})
	}

	var title string
	var content string

	err = m.db.QueryRowContext(ctx, `SELECT MOOD.TITLE, MOOD.CONTENT
          FROM MOOD
          WHERE MOOD.MOOD_ID = ? AND MOOD.MOOD_ACCESS_CODE = ?`, request.GetMoodId(), request.GetMoodAccessCode()).Scan(&title, &content)

	if err != nil {
		return nil, err
	}

	moodStats := make([]*proto.MoodStat, 0)
	for _, s := range stats {
		moodStats = append(moodStats, s)
	}

	return &proto.GetMoodResponse{
		Title:   title,
		Content: content,
		Entries: entries,
		Stats:   moodStats,
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
