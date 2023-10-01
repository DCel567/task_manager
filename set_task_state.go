package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type set_task_state struct {
	buttons []string
	cursor  int
}

func InitialSetTask() set_task_state {
	return set_task_state{
		buttons: []string{"Python", "C++", "Go", "Back"},
		cursor:  0,
	}
}

func (s set_task_state) Init() tea.Cmd {
	return nil
}

func (s set_task_state) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return s, tea.Quit

		case "up", "k":
			s.cursor--
			if s.cursor < 0 {
				s.cursor = len(s.buttons) - 1
			}

		case "down", "j":
			s.cursor = (s.cursor + 1) % len(s.buttons)

		case "enter", " ":
			switch s.cursor {
			case 0:

			case 1:

			case 2:

			case 3:
				// p := tea.NewProgram(InitialHello())
				// if _, err := p.Run(); err != nil {
				// 	fmt.Printf("Alas, there's been an error: %v", err)
				// 	os.Exit(1)
				// }
			}
		}
	}

	return s, nil
}

func (st set_task_state) View() string {
	s := "\nWhat you gonna do?\n\n"

	for i, choice := range st.buttons {
		cursor := " " //no cursor
		if st.cursor == i {
			cursor = ">" //cursor
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}
