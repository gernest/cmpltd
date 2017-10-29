package models

type Command struct {
	Name        string
	Path        string
	SubCommands []SubCommand
	Flags       []Flag
}

type SubCommand struct {
	Command
}

type Flag struct {
	Name  string
	Alias string
}

type App struct {
	Name     string
	Commands []Command
	Flags    []Flag
}
