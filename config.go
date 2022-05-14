package main

type Config struct {
	Servers     []*Server
	WaitAddress []string
}

type Server struct {
	Directory      string
	Exe            string
	Environment    []string
	Args           []string
	FaildSecond    int
	UpdateFileName string
}
