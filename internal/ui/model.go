package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/ArcaneCrowA/go-todo/internal/task"
)

type TodoList struct {
	list []task.Item
}

func (t TodoList) Init() tea.Cmd {
	return func() tea.Msg {
		return t.list
	}
}
