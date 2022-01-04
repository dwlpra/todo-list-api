package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@/ade_dwi_putra")
	if err != nil {
		return nil, err
	}
	// See "Important settings" section.
	db.SetConnMaxIdleTime(time.Minute * 5)
	db.SetConnMaxLifetime(time.Minute * 60)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)

	return db, nil
}
