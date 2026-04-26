package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/ArcaneCrowA/go-todo/internal/ui"
)

func main() {
	model := ui.TodoList{}

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
