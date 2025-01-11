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
		printPrompt(strings.Join(args[1:], " "))
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
