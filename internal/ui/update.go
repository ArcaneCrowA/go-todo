package ui

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/ArcaneCrowA/go-todo/internal/task"
)

func (m TodoList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.state {
	case listView:
		m.listUpdate(msg, &cmd)
	case editView:
		m.editUpdate(msg, &cmd)
	}
	return m, cmd
}

func (m *TodoList) listUpdate(msg tea.Msg, cmd *tea.Cmd) {
	switch msg := msg.(type) {
	case []task.Item:
		m.list = msg
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			*cmd = tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.list)-1 {
				m.cursor++
			}
		case "e":
			m.state = editView
			m.textInput.Focus()
			*cmd = textinput.Blink
		}
		m.list, _ = m.storage.Load()
	}
}

func (m *TodoList) editUpdate(msg tea.Msg, cmd *tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if msg.String() == "enter" {
			m.textInput.Blur()
			m.state = listView
		}
	}
	m.textInput, *cmd = m.textInput.Update(msg)
}
