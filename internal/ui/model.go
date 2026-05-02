package ui

import (
	"log/slog"
	"os"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/ArcaneCrowA/go-todo/internal/task"
)

type sessionState int

const (
	listView sessionState = iota
	addView
	editView
)

type TodoList struct {
	list        []task.Item
	storage     Storage
	cursor      int
	state       sessionState
	inputs      []textinput.Model
	focusIndex  int
	status      string
	statusIndex int
	isAdding    bool
}

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
)

func New(storage Storage) TodoList {
	m := TodoList{
		storage: storage,
		state:   listView,
		inputs:  make([]textinput.Model, 2),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.CharLimit = 32

		s := t.Styles()
		s.Cursor.Color = lipgloss.Color("205")
		s.Focused.Prompt = focusedStyle

		t.SetStyles(s)

		switch i {
		case 0:
			t.Placeholder = "Title"
			t.Focus()
		case 1:
			t.Placeholder = "Description"
			t.Blur()
		}
		m.inputs[i] = t
	}

	return m
}

func (m TodoList) Init() tea.Cmd {
	return func() tea.Msg {
		list, err := m.storage.Load()
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}

		return list
	}
}
