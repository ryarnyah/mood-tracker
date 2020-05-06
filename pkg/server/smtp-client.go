package server

import (
	"database/sql"
	"fmt"
	"net/smtp"
	"sync"
	"time"

	"github.com/golang/glog"
)

const (
	pollingDuration = 10 * time.Second
)

type SmtpDbPoller struct {
	db *sql.DB

	host string
	port int

	senderEmail string
	externalURL string

	username string
	password string

	ticker *time.Ticker
}

func NewSmtpDbPoller(db *sql.DB, host string, port int, username, password, senderEmail, externalURL string) *SmtpDbPoller {
	return &SmtpDbPoller{
		db: db,

		host:     host,
		port:     port,
		username: username,
		password: password,

		senderEmail: senderEmail,
		externalURL: externalURL,

		ticker: time.NewTicker(pollingDuration),
	}
}

func (c *SmtpDbPoller) Start(wg *sync.WaitGroup, stopCh chan (interface{})) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-c.ticker.C:
				c.pollEmails()
				// Reset ticker (must be updated with Go 1.15)
				c.ticker.Stop()
				c.ticker = time.NewTicker(pollingDuration)
			case <-stopCh:
				c.ticker.Stop()
				return
			}
		}
	}()
}

func (c *SmtpDbPoller) pollEmails() {
	tx, err := c.db.Begin()
	if err != nil {
		glog.Error(err)
		return
	}
	defer tx.Rollback()

	rows, err := tx.Query(`SELECT ENTRY.ENTRY_ID, ENTRY.MOOD_ID, ENTRY.ENTRY_ACCESS_CODE, MAIL.EMAIL
           FROM MAIL JOIN ENTRY ON MAIL.ENTRY_ID = ENTRY.ENTRY_ID
           LIMIT 10`)
	if err != nil {
		glog.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var entryId int64
		var moodId int64
		var entryAccessCode string
		var email string
		err = rows.Scan(&entryId, &moodId, &entryAccessCode, &email)
		if err != nil {
			glog.Error(err)
			return
		}

		if c.host != "" {
			var auth smtp.Auth

			if c.username != "" && c.password != "" {
				auth = smtp.PlainAuth("", c.username, c.password, c.host)
			}

			to := []string{email}
			msg := []byte(
				fmt.Sprintf(
					"From: %s\r\n"+
						"To: %s\r\n"+
						"Subject: Your daily mood personal URL!\r\n"+
						"\r\n"+
						"Get Access to your daily mood URL here %s/entry/%d/%s.\r\n", c.senderEmail, email, c.externalURL, moodId, entryAccessCode),
			)

			err := smtp.SendMail(fmt.Sprintf("%s:%d", c.host, c.port), auth, c.senderEmail, to, msg)
			if err != nil {
				glog.Error(err)
				return
			}
		}

		_, err = tx.Exec(`DELETE FROM MAIL
                  WHERE MAIL.ENTRY_ID = ? AND MAIL.EMAIL = ?`, entryId, email)
		if err != nil {
			glog.Error(err)
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		glog.Error(err)
		return
	}
}
