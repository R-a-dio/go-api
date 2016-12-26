package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Database *sqlx.DB
var (
	lastPlayedStmt *sqlx.Stmt
	queueStmt      *sqlx.Stmt
	mainStmt       *sqlx.Stmt
)

func openDatabase(dsn string) error {
	var err error
	Database, err = sqlx.Open("mysql", dsn)

	lastPlayedStmt, err = Database.Preparex("SELECT esong.meta AS meta, UNIX_TIMESTAMP(eplay.dt) AS time FROM `eplay` LEFT JOIN `esong` ON eplay.isong = esong.id ORDER BY eplay.dt DESC LIMIT 5;")
	if err != nil {
		return err
	}

	queueStmt, err = Database.Preparex("SELECT meta AS meta, UNIX_TIMESTAMP(time) AS time, type FROM `queue` ORDER BY `time` ASC LIMIT 5;")
	if err != nil {
		return err
	}

	mainStmt, err = Database.Preparex(`SELECT np, listeners, bitrate, djid,
        isafkstream, isstreamdesk, start_time, end_time, lastset, trackid, 
        thread, requesting, djs.djname, djtext, djimage, visible, css, djcolor,
        theme_id, priority, role  FROM streamstatus JOIN djs ON djs.id = djid LIMIT 1;`)
	if err != nil {
		return err
	}

	return nil
}
