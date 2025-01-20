package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ExecuteCommand(command string) bool {
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
			if query, ok := CommandRegistry[command]; ok {
				fmt.Println(query.Name + " is a " + query.Type)
			} else if file := findExecutable(command); file != "" {
				fmt.Println(command + " is " + file) 
			} else {
				fmt.Println(command + ": not found")
			}
		}

	case "man":
		if len(args) < 2 {
			fmt.Println("man: missing argument")
		} else {
			command := args[1]
			if query, ok := CommandRegistry[command]; ok {
				fmt.Println(query.Name + ": " + query.Description)  
			} else {
				fmt.Println(command + ": not found")
			}
		}

	default:
		runExecutable(args[0], args[1:])
	}
	return true
}

func findExecutable(query string) string {
	if path, err := exec.LookPath(query); err == nil {
		return path
	}
	return ""
}

func runExecutable(file string, args []string) {
	executable := findExecutable(file)
	if executable == "" {
		fmt.Println(file + ": command not found")
		
	cmd := exec.Command(executable, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
	}
}
}