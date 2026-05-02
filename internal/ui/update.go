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
		switch msg := msg.(type) {
		case []task.Item:
			m.list = msg
			return m, nil
		case tea.KeyPressMsg:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
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
				return m, textinput.Blink
			}
			m.list, _ = m.storage.Load()
		}
	case editView:
		switch msg := msg.(type) {
		case tea.KeyPressMsg:
			if msg.String() == "enter" {
				m.textInput.Blur()
				m.state = listView
			}
		}
		m.textInput, cmd = m.textInput.Update(msg)
	}
	return m, cmd
}
