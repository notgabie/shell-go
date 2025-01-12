package commands

import (
	"fmt"
	"os"
	"strconv"
	"path/filepath"
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
			} else if file := isExecutable(command); file != "" {
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
		fmt.Println(args[0] + ": command not found")
	}
	return true
}

func isExecutable(query string) string {
	path := os.Getenv("PATH")
	dirs := strings.Split(path, ":")

	for _, dir := range dirs {
		file := filepath.Join(dir, query)
		if _, err := os.Stat(file); err == nil {
			return file
		}
	}
	return ""
}