package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ArcaneCrowA/go-todo/internal/task"
	"github.com/huandu/go-sqlbuilder"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteStore struct {
	db *sql.DB
}

func NewSqliteStore(path string) (*SqliteStore, func()) {
	db, err := sql.Open("sqlite3", path+"?_loc=UTC")
	if err != nil {
		log.Fatalf("couldn't open db connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("couldn't reach database: %v", err)
	}

	query := sqlbuilder.CreateTable("tasks").IfNotExists().
		Define("id", "INTEGER", "PRIMARY KEY", "NOT NULL").
		Define("name", "TEXT", "NOT NULL").
		Define("description", "TEXT", "NOT NULL").
		Define("status", "TEXT", "NOT NULL").
		Define("created", "DATETIME", "NOT NULL").
		Define("updated", "DATETIME", "NOT NULL").
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

func (s *SqliteStore) Load() ([]task.Item, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("*")
	sb.From("tasks")
	query, _ := sb.Build()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("error closing rows: %v", err)
		}
	}()

	var items []task.Item
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item task.Item
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
			&item.Status,
			&item.Created,
			&item.Updated,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *SqliteStore) Save(item task.Item) error {
	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("tasks")
	ib.Cols("name", "description", "status", "created", "updated")
	ib.Values(item.Name, item.Description, item.Status, item.Created, item.Updated)
	query, args := ib.Build()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	if affected, err := res.RowsAffected(); affected == 0 || err != nil {
		return fmt.Errorf("couldn't save: %w", err)
	}

	return nil
}

func (s *SqliteStore) Delete(item task.Item) error {
	return nil
}

func (s *SqliteStore) Edit(item task.Item) error {
	return nil
}
