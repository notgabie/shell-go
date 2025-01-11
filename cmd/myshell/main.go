package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	prompt = "$ "
)

type CommandInfo struct {
	Name string
	Type string
	Description string
}

var commands = map[string]CommandInfo{
	"exit": {Name: "exit", Type: "builtin", Description: "Exit the shell"},
	"echo": {Name: "echo", Type: "builtin", Description: "Print the arguments to the standard output"},
	"type": {Name: "type", Type: "builtin", Description: "Display information about command type"},
}

func printPrompt(input ...interface{}) {
	if len(input) == 0 {
		fmt.Print(prompt)
		return
	} 
	fmt.Println(prompt + fmt.Sprint(input...))
}

func executeCommand(command string) bool {
	args := strings.Fields(command)

	if len(args) == 0 {
		return true
	}

	switch args[0] {
	case "exit":
		exitCode := 0
		if len(args) > 1 {
			if code, err := strconv.Atoi(args[1]); err == nil {
				exitCode = code
			}
	}
	os.Exit(exitCode)
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))

	case "type":
		if len(args) < 2 {
			fmt.Println("type: missing argument")
		} else {
			command := args[1]
			if query, ok := commands[command]; ok {
				fmt.Println(query.Name + " is a shell " + query.Type)
			} else {
				fmt.Println(command + ": not found")
			}
		}

	case "man":
		if len(args) < 2 {
			fmt.Println("man: missing argument")
		} else {
			command := args[1]
			if query, ok := commands[command]; ok {
				fmt.Println(query.Name + ": " + query.Description)
			} else {
				fmt.Println(command + ": not found")
			}
		}

	default:
		fmt.Println(args[0] + ": command not found")
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		printPrompt()
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			continue
		}

		trimmedCommand := strings.TrimSpace(command)
		executeCommand(trimmedCommand)
	}
}
