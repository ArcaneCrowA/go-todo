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
		fmt.Fprintf(&b, "%d. %s\n", i+1, item.Name)
	}

	return tea.NewView(b.String())
}
