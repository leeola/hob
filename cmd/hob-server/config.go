package main

type Config struct {
	BindAddr string
	Events   map[string]string
	Actions  Actions
}

type Actions struct {
	Subprocs []Subproc
}

type Subproc struct {
	Action string
	Bin    string
	Args   []string
}
