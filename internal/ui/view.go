package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (t TodoList) View() tea.View {
	b := strings.Builder{}

	b.WriteString("to do list\n\n")

	for i, item := range t.list {
		cursor := " " // no cursor
		if t.cursor == i {
			cursor = ">" // cursor!
		}
		fmt.Fprintf(&b, "%s %s\n", cursor, item.Name)
	}

	return tea.NewView(b.String())
}
