package ui

import "github.com/ArcaneCrowA/go-todo/internal/task"

type Storage interface {
	Save(task task.Item) error
	Delete(task task.Item) error
	Load() ([]task.Item, error)
}
