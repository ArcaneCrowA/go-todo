package ui

import (
	"log/slog"
	"os"

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
			m.inputs[0].SetValue(m.list[m.cursor].Name)
			m.inputs[1].SetValue(m.list[m.cursor].Description)
			m.status = m.list[m.cursor].Status
			m.statusIndex = 0
			m.focusIndex = 0
			*cmd = textinput.Blink
		}
		m.list, _ = m.storage.Load()
	}
}

func (m *TodoList) editUpdate(msg tea.Msg, cmd *tea.Cmd) {
	item := m.list[m.cursor]

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			*cmd = tea.Quit
			return
		case "h", "left":
			if m.focusIndex == 2 {
				m.statusIndex = (m.statusIndex - 1 + task.NumStatuses) % task.NumStatuses
				m.status = task.Statuses[m.statusIndex]
			}
		case "l", "right":
			if m.focusIndex == 2 {
				m.statusIndex = (m.statusIndex + 1) % task.NumStatuses
				m.status = task.Statuses[m.statusIndex]
			}
		case "tab", "shift+tab":
			if msg.String() == "tab" {
				m.focusIndex++
			} else {
				m.focusIndex--
			}

			if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs) - 1
			}
			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range m.inputs {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus()
					continue
				}
				m.inputs[i].Blur()
			}

		case "enter":
			item.Name = m.inputs[0].Value()
			item.Description = m.inputs[1].Value()
			item.Status = task.Statuses[m.statusIndex]
			if err := m.storage.Edit(item); err != nil {
				slog.Error(err.Error())
				os.Exit(1)
			}
			for i := range m.inputs {
				m.inputs[i].Reset()
			}
			m.state = listView

		}
		m.list, _ = m.storage.Load()
		*cmd = m.updateInputs(msg)
	}
}

func (m *TodoList) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
