package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseConn() (*sql.DB, error) {
	connString := "root:password@tcp(127.0.0.1:3306)/class_roster?parseTime=true"

	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("sql.Open error:", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	return db, nil
}
