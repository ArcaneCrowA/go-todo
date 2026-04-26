package main

import tea "charm.land/bubbletea/v2"

func main() {
	p := tea.NewProgram(newSimplePage())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

type simplePage struct {
	text string
}

func newSimplePage() simplePage {
	return simplePage{"Hello "}
}

func (s simplePage) Init() tea.Cmd { return nil }

func (s simplePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return s, tea.Quit
		}
	}
	return s, nil
}

func (s simplePage) View() tea.View {
	return tea.View{Content: s.text}
}
