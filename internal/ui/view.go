package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m TodoList) View() tea.View {
	switch m.state {
	case addView:

	case editView:
		return tea.NewView(fmt.Sprintf(
			"Enter new value:\n\n%s\n\n(press enter to save)",
			m.textInput.View(),
		))
	default:
		b := strings.Builder{}

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
