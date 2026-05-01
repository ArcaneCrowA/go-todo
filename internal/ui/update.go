package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/ArcaneCrowA/go-todo/internal/task"
)

func (t TodoList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case []task.Item:
		t.list = msg
		return t, nil
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return t, tea.Quit
		}
	}
	return t, nil
}
