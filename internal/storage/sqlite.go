package storage

import (
	"context"
	"database/sql"
	"log"
	"time"

	sb "github.com/huandu/go-sqlbuilder"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteStore struct {
	db *sql.DB
}

func NewSqliteStore(path string) (*SqliteStore, func()) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("couldn't open db connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("couldn't reach database: %v", err)
	}

	query := sb.CreateTable("tasks").IfNotExists().
		Define("id", "INTEGER", "PRIMARY KEY", "NOT NULL").
		Define("name", "TEXT", "NOT NULL").
		Define("description", "TEXT", "NOT NULL").
		Define("status", "TEXT", "NOT NULL").
		Define("created", "TEXT", "NOT NULL").
		Define("updated", "TEXT", "NOT NULL").
		String()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err = db.ExecContext(ctx, query)
	if err != nil {
		db.Close()
		log.Fatal("coudn't create table")
	}

	cleanup := func() {
		if err := db.Close(); err != nil {
			log.Printf("error closing database: %v", err)
		}
	}

	return &SqliteStore{db: db}, cleanup
}
