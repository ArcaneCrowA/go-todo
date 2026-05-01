package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/ArcaneCrowA/go-todo/internal/storage"
	"github.com/ArcaneCrowA/go-todo/internal/ui"
)

func main() {
	jsonStorage := storage.NewJSONStore("items.json")
	model := ui.New(jsonStorage)

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
