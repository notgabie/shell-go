package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	commands "github.com/codecrafters-io/shell-starter-go/cmd/myshell/commands"
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
		commands.ExecuteCommand(trimmedCommand)
	}
}
