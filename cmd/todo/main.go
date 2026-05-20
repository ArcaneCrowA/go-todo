package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
	"github.com/ArcaneCrowA/go-todo/internal/storage"
	"github.com/ArcaneCrowA/go-todo/internal/ui"
)

func main() {
	// store := storage.NewJSONStore("items.json")
	store, cleanup := storage.NewSqliteStore("items.db")
	defer cleanup()

	model := ui.New(store)

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
