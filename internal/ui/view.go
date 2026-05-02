package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m TodoList) View() tea.View {
	var b strings.Builder

	switch m.state {
	case addView:
	case editView:
		var c *tea.Cursor
		for i := range m.inputs {
			b.WriteString(m.inputs[i].View())
			if i < len(m.inputs)-1 {
				b.WriteRune('\n')
			}
		}
		v := tea.NewView(b.String())
		v.Cursor = c
		return v

	default:
		b.WriteString("ToDo\n\n")

		for i, item := range m.list {
			cursor := " " // no cursor
			if m.cursor == i {
				cursor = ">" // cursor!
			}
			fmt.Fprintf(&b, "%s %s\n", cursor, item.Name)
		}

		b.WriteString("\nctrl+c or q to exit")

		return tea.NewView(b.String())
	}
	return tea.View{}
}
