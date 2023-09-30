package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// type model struct {
// 	choices  []string
// 	cursor   int
// 	selected map[int]struct{}
// }

type model struct {
	choices  []string
	cursor   int
	selected map[int]bool
}

func initialModel(variants []string) model {
	return model{
		choices:  variants,
		selected: make(map[int]bool, len(variants)),
		//selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}

		case "down", "j":
			m.cursor = (m.cursor + 1) % len(m.choices)

		case "enter", " ":
			res := m.selected[m.cursor]
			if res {
				//delete(m.selected, m.cursor)
				m.selected[m.cursor] = false
			} else {
				//m.selected[m.cursor] = struct{}{}
				m.selected[m.cursor] = true
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "What should we buy at the market?\n\n"

	for i, choice := range m.choices {
		cursor := " " //no cursor
		if m.cursor == i {
			cursor = ">" //cursor
		}

		checked := " " //not selected
		if res := m.selected[i]; res {
			checked = "x" //selected
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}

func main() {
	p := tea.NewProgram(initialModel([]string{"Buy carrots", "Buy celery", "Buy kohlrabi"}))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
