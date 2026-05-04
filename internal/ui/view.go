package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m TodoList) View() tea.View {
	var b strings.Builder

	switch m.state {
	case editView:
		var c *tea.Cursor
		for i := range m.inputs {
			b.WriteString(m.inputs[i].View())
			if i < len(m.inputs)-1 {
				b.WriteRune('\n')
			}
		}
		style := blurredStyle
		if m.focusIndex == 2 {
			style = focusedStyle
		}
		fmt.Fprintf(&b, "\n%s %s\n", style.Render("Status:"), m.status)

		v := tea.NewView(b.String())
		v.Cursor = c
		return v

	default:
		b.WriteString(menuStyle.Render("\t\t\tTODO\n\n"))

		for i, item := range m.list {
			style := blurredStyle
			if m.cursor == i {
				style = focusedStyle
			}

			fmt.Fprintf(&b, "%s %s\n", style.Render(fmt.Sprintf("%2d: %s", i, item.Name)), item.Status)
		}

		b.WriteString(helpStyle.Render(`
			a : add new task
			e : edit task
			d : delete task
			q : exit`))

		return tea.NewView(b.String())
	}
}
