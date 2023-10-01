package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type hello_state struct {
	buttons []string
	cursor  int
}

func InitialHello() hello_state {
	return hello_state{
		buttons: []string{"Set task", "Statistics", "Quit"},
		cursor:  0,
	}
}

func (h hello_state) Init() tea.Cmd {
	return nil
}

func (h hello_state) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return h, tea.Quit

		case "up", "k":
			h.cursor--
			if h.cursor < 0 {
				h.cursor = len(h.buttons) - 1
			}

		case "down", "j":
			h.cursor = (h.cursor + 1) % len(h.buttons)

		case "enter", " ":
			switch h.cursor {
			case 0:
				p := tea.NewProgram(InitialSetTask(), tea.WithAltScreen())
				if _, err := p.Run(); err != nil {
					fmt.Printf("Alas, there's been an error: %v", err)
					os.Exit(1)
				}
			case 1:

			case 2:
				return h, tea.Quit
			}
		}
	}

	return h, nil
}

func (h hello_state) View() string {
	s := "\nEgor's Task Manager v1.0\n\n"

	for i, button := range h.buttons {
		cursor := " " //no cursor
		if h.cursor == i {
			cursor = ">" //cursor
		}

		s += fmt.Sprintf("%s %s\n", cursor, button)
	}

	s += "\nPress q to quit.\n"

	return s
}
