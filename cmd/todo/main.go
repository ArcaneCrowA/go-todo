package main

import tea "charm.land/bubbletea/v2"

func main() {
	p := tea.NewProgram()
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
