package main

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	mainUpdateTickrate = time.Second * 2
	listUpdateTickrate = time.Second * 10
)

func startUpdater() {
	go mainUpdater()
	go listUpdater()
}

func mainUpdater() {
	ticker := time.NewTicker(mainUpdateTickrate)
	defer ticker.Stop()

	for range ticker.C {
		var a API
		err := mainStmt.QueryRowx().StructScan(&a)
		if err != nil {
			log.Printf("database error: %s", err)
			continue
		}

		apiMain.Store(a)
	}
}

func listUpdater() {
	ticker := time.NewTicker(listUpdateTickrate)
	defer ticker.Stop()

	for range ticker.C {
		queryLists()
	}
}

func queryLists() {
	tx, err := Database.Beginx()
	if err != nil {
		log.Printf("database transaction error: %s", err)
		return
	}
	defer tx.Commit()

	queue, err := queryListStmt(tx.Stmtx(queueStmt))
	if err != nil {
		log.Printf("database error when retrieving queue: %s", err)
		return
	}

	lp, err := queryListStmt(tx.Stmtx(lastPlayedStmt))
	if err != nil {
		log.Printf("database error when retrieving last played: %s", err)
		return
	}

	apiQueue.Store(queue)
	apiLastPlayed.Store(lp)
}

func queryListStmt(stmt *sqlx.Stmt) ([]ListEntryAPI, error) {
	rows, err := stmt.Queryx()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []ListEntryAPI
	for rows.Next() {
		var l ListEntryAPI
		err = rows.StructScan(&l)
		if err != nil {
			return nil, err
		}

		l.Time = formatTimeAgo(l.Timestamp)
		res = append(res, l)
	}

	return res, nil
}

var timeagoFormat = `<time class="timeago" datetime="2006-01-02T15:04:05-0700">15:04:05</time>`

func formatTimeAgo(unix int64) string {
	return time.Unix(unix, 0).Format(timeagoFormat)
}
