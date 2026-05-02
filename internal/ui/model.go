package ui

import (
	"log/slog"
	"os"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/ArcaneCrowA/go-todo/internal/task"
)

type sessionState int

const (
	listView sessionState = iota
	addView
	editView
)

type TodoList struct {
	list      []task.Item
	storage   Storage
	cursor    int
	state     sessionState
	textInput textinput.Model
}

func New(storage Storage) TodoList {
	ti := textinput.New()
	ti.Placeholder = "Edit task..."
	ti.CharLimit = 100
	ti.SetWidth(40)
	return TodoList{
		storage:   storage,
		state:     listView,
		textInput: ti,
	}
}

func (m TodoList) Init() tea.Cmd {
	return func() tea.Msg {
		list, err := m.storage.Load()
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}

		return list
	}
}
