package ui

import (
	"log/slog"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/ArcaneCrowA/go-todo/internal/task"
)

type TodoList struct {
	list    []task.Item
	storage Storage
	cursor  int
}

func New(storage Storage) TodoList {
	return TodoList{
		storage: storage,
	}
}

func (t TodoList) Init() tea.Cmd {
	return func() tea.Msg {
		list, err := t.storage.Load()
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}

		return list
	}
}
