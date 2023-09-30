package main

type state interface {
	Init() state
}

type hello_state struct {
	buttons []string
	cursor  int
}

type set_task_state struct {
	buttons []string
	cursor  int
}

func (h hello_state) Init() state {
	return hello_state{
		buttons: []string{"Set task", "Statistics", "Quit"},
		cursor:  0,
	}
}

func (st set_task_state) Init() state {
	return set_task_state{
		buttons: []string{"Python", "C++", "Go", "Quit"},
		cursor:  0,
	}
}
