package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

const schema = `CREATE TABLE scheduler (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					date CHAR(8) NOT NULL DEFAULT "",
					title VARCHAR(256) NOT NULL DEFAULT "",
					comment TEXT NOT NULL DEFAULT "",
					repeat VARCHAR(128) NOT NULL DEFAULT ""
);

CREATE INDEX idx_scheduler_date ON scheduler (date);`

var db *sql.DB
var dbFile = "scheduler.db"

func Init() error {
	var install bool

	_, err := os.Stat(dbFile)
	if err != nil {
		install = true
	}

	envDbFile := os.Getenv("TODO_DBFILE")
	if len(envDbFile) > 0 {
		dbFile = envDbFile
	}

	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}

	if install {
		_, err = db.Exec(schema)
		if err != nil {
			return err
		}
	}

	return nil
}

func Close() error {
	return db.Close()
}
