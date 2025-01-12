package commands

type CommandInfo struct {
	Name string
	Type string
	Description string
}

var CommandRegistry = map[string]CommandInfo{
	"exit": {Name: "exit", Type: "shell builtin", Description: "Exit the shell"},
	"echo": {Name: "echo", Type: "shell builtin", Description: "Print the arguments to the standard output"},
	"type": {Name: "type", Type: "shell builtin", Description: "Display information about command type"},
}