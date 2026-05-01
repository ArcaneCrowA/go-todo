package ui

import "github.com/ArcaneCrowA/go-todo/internal/task"

type Storage interface {
	Save(item task.Item) error
	Delete(item task.Item) error
	Load() ([]task.Item, error)
}
