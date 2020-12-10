package server

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"time"

	"github.com/golang/protobuf/ptypes"
	proto "github.com/ryarnyah/mood-tracker/proto"
	"github.com/segmentio/ksuid"
)

type moodServer struct {
	db *sql.DB
}

func NewMoodServer(db *sql.DB) proto.MoodServer {
	return &moodServer{
		db: db,
	}
}

func checkSignedData(signKey, data, expectedSignature string) error {
	// Check sign key
	randomKey, err := hex.DecodeString(signKey)
	if err != nil {
		return err
	}
	signature, err := hex.DecodeString(expectedSignature)
	if err != nil {
		return err
	}
	moodSigner := hmac.New(sha256.New, randomKey)
	_, err = moodSigner.Write([]byte(data))
	if err != nil {
		return err
	}
	moodIdSigned := moodSigner.Sum(nil)

	if !hmac.Equal(moodIdSigned, signature) {
		return errors.New("signature mismatch")
	}

	return nil
}

func (m *moodServer) GetMoodFromEntry(ctx context.Context, request *proto.GetMoodFromEntryRequest) (*proto.GetMoodFromEntryResponse, error) {
	var title string
	var content string
	var signKey string

	err := m.db.QueryRowContext(ctx, `SELECT MOOD.SIGN_KEY
          FROM MOOD
          WHERE MOOD.MOOD_ID = ?`, request.GetMoodId()).Scan(&signKey)
	if err != nil {
		return nil, err
	}

	if err = checkSignedData(signKey, request.GetMoodId()+request.GetEntryId(), request.GetEntrySignature()); err != nil {
		return nil, err
	}

	now := time.Now()
	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, now.Location())

	err = m.db.QueryRowContext(ctx, `SELECT MOOD.TITLE, MOOD.CONTENT, MOOD.SIGN_KEY
          FROM MOOD JOIN ENTRY ON MOOD.MOOD_ID = ENTRY.MOOD_ID
          LEFT JOIN RECORD ON ENTRY.ENTRY_ID = RECORD.ENTRY_ID
          WHERE ENTRY.MOOD_ID = ? AND ENTRY.ENTRY_ID = ? AND (RECORD.ENTRY_ID IS NULL OR RECORD.RECORD_DATETIME <> ?)`,
		request.GetMoodId(), request.GetEntryId(), today).Scan(&title, &content, &signKey)

	if err != nil {
		return nil, err
	}

	return &proto.GetMoodFromEntryResponse{
		Title:   title,
		Content: content,
	}, nil
}

func (m *moodServer) AddEntry(ctx context.Context, request *proto.AddEntryRequest) (*proto.AddEntryResponse, error) {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	now := time.Now()
	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, now.Location())

	var signKey string
	err = m.db.QueryRowContext(ctx, `SELECT MOOD.SIGN_KEY
          FROM MOOD
          WHERE MOOD.MOOD_ID = ?`, request.GetMoodId()).Scan(&signKey)
	if err != nil {
		return nil, err
	}

	if err = checkSignedData(signKey, request.GetMoodId()+request.GetEntryId(), request.GetEntrySignature()); err != nil {
		return nil, err
	}

	var entryID string
	err = tx.QueryRowContext(ctx, `SELECT ENTRY.ENTRY_ID
          FROM ENTRY LEFT JOIN RECORD ON ENTRY.ENTRY_ID = RECORD.ENTRY_ID
          WHERE ENTRY.MOOD_ID = ? AND ENTRY.ENTRY_ID = ? AND (RECORD.ENTRY_ID IS NULL OR RECORD.RECORD_DATETIME <> ?)`, request.GetMoodId(), request.GetEntryId(), today).Scan(&entryID)
	if err == sql.ErrNoRows {
		return nil, errors.New("entry-id or mood-id is invalid or expired")
	} else if err != nil {
		return nil, err
	}

	updateEntry, err := tx.PrepareContext(ctx, "INSERT INTO RECORD (ENTRY_ID, RECORD, COMMENT, RECORD_DATETIME) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer updateEntry.Close()
	_, err = updateEntry.ExecContext(ctx, entryID, request.GetEntry().GetRecord(), request.GetEntry().GetComment(), today)
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
	var signKey string
	err := m.db.QueryRowContext(ctx, `SELECT MOOD.SIGN_KEY
          FROM MOOD
          WHERE MOOD.MOOD_ID = ?`, request.GetMoodId()).Scan(&signKey)
	if err != nil {
		return nil, err
	}

	if err = checkSignedData(signKey, request.GetMoodId(), request.GetMoodSignature()); err != nil {
		return nil, err
	}

	rows, err := m.db.QueryContext(ctx, `SELECT RECORD.RECORD, RECORD.COMMENT, RECORD.RECORD_DATETIME
          FROM RECORD JOIN ENTRY ON RECORD.ENTRY_ID = ENTRY.ENTRY_ID
          JOIN MOOD ON MOOD.MOOD_ID = ENTRY.MOOD_ID
          WHERE MOOD.MOOD_ID = ?
          ORDER BY RECORD.RECORD_DATETIME DESC`, request.GetMoodId())
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
	statRows, err := m.db.QueryContext(ctx, `SELECT RECORD.RECORD, RECORD.RECORD_DATETIME, COUNT(1)
          FROM ENTRY
          JOIN RECORD ON RECORD.ENTRY_ID = ENTRY.ENTRY_ID
          JOIN MOOD ON MOOD.MOOD_ID = ENTRY.MOOD_ID
          WHERE MOOD.MOOD_ID = ? AND RECORD.RECORD_DATETIME > date('now', '-7 day')
          GROUP BY RECORD.RECORD, RECORD.RECORD_DATETIME
          ORDER BY RECORD.RECORD, RECORD.RECORD_DATETIME`, request.GetMoodId())
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
          WHERE MOOD.MOOD_ID = ?`, request.GetMoodId()).Scan(&title, &content)

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
	if request.GetNumberOfRecordsNeeded() == uint32(0) {
		if len(request.GetEmails()) == 0 {
			return nil, errors.New("number of records needed or emails is mandatory")
		}
	}

	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	randomKey := make([]byte, 256)
	_, err = rand.Read(randomKey)
	if err != nil {
		return nil, err
	}
	randomKeyhex := hex.EncodeToString(randomKey)

	moodUUID := ksuid.New()

	moodSigner := hmac.New(sha256.New, randomKey)
	_, err = moodSigner.Write([]byte(moodUUID.String()))
	if err != nil {
		return nil, err
	}
	moodSignature := hex.EncodeToString(moodSigner.Sum(nil))

	// Create mood entry
	moodStmt, err := tx.PrepareContext(ctx, "INSERT INTO MOOD (MOOD_ID, TITLE, CONTENT, SIGN_KEY) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer moodStmt.Close()

	_, err = moodStmt.ExecContext(ctx, moodUUID.String(), request.GetTitle(), request.GetContent(), randomKeyhex)
	if err != nil {
		return nil, err
	}

	// Create entry entries
	entryStmt, err := tx.PrepareContext(ctx, "INSERT INTO ENTRY (ENTRY_ID, MOOD_ID) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	defer entryStmt.Close()

	recordsIds := []*proto.EntrySigned{}
	var i uint32
	for i = 0; i < request.GetNumberOfRecordsNeeded(); i++ {
		entryUUID := ksuid.New()
		// Sign entry
		signer := hmac.New(sha256.New, randomKey)
		_, err = signer.Write([]byte(moodUUID.String() + entryUUID.String()))
		if err != nil {
			return nil, err
		}
		entrySignature := hex.EncodeToString(signer.Sum(nil))
		_, err = entryStmt.ExecContext(ctx, entryUUID.String(), moodUUID)
		if err != nil {
			return nil, err
		}

		recordsIds = append(recordsIds, &proto.EntrySigned{
			EntryId:        entryUUID.String(),
			EntrySignature: entrySignature,
		})
	}

	mailStmt, err := tx.PrepareContext(ctx, "INSERT INTO MAIL (ENTRY_ID, EMAIL) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	defer mailStmt.Close()
	for _, email := range request.GetEmails() {
		entryUUID := ksuid.New()
		_, err := entryStmt.ExecContext(ctx, entryUUID.String(), moodUUID)
		if err != nil {
			return nil, err
		}
		_, err = mailStmt.ExecContext(ctx, entryUUID, email)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &proto.CreateMoodResponse{
		MoodId:        moodUUID.String(),
		MoodSignature: moodSignature,
		EntriesIds:    recordsIds,
	}, nil
}
