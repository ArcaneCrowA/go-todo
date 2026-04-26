package storage

import "github.com/ArcaneCrowA/go-todo/internal/task"

type Storage interface {
	Save(task task.Task) error
	Load() ([]task.Task, error)
}
