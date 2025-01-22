package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func ExecuteCommand(command string) bool {
	args := strings.Fields(command)

	if len(args) == 0 {
		return true
	}

	if strings.HasPrefix(args[1], "'") && strings.HasSuffix(args[len(args)-1], "'") {
		quotedString := command[strings.Index(command, "'"):]
		quotedString = strings.TrimPrefix(quotedString, "'")
		quotedString = strings.TrimSuffix(quotedString, "'")
		quotedString = strings.Replace(quotedString, "'", "", -1)
		args = append(args[:1], quotedString)
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
			}
		}

	case "man":
		if len(args) < 2 {
			fmt.Println("man: missing argument")
		} else {
			command := args[1]
			if query, ok := CommandRegistry[command]; ok {
				fmt.Println(query.Name + ": " + query.Description)  
			}
		}

	case "pwd":
		if wd, err := os.Getwd(); err == nil {
			fmt.Println(wd)
		} else {
			fmt.Println("pwd: error getting working directory:", err)
		}

	case "cd":
		changeDirectory(args[1])

	default:
		runExecutable(args[0], args[1:])
	}
	return true
}

func findExecutable(query string) string {
	if path, err := exec.LookPath(query); err == nil {
		return path
	} else {
		fmt.Println(query + ": not found")
	}
	return ""
}

func runExecutable(file string, args []string) {
	executable := findExecutable(file)
	dirs := strings.Split(executable, "/")
	executable = dirs[len(dirs)-1]

	cmd := exec.Command(executable, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func changeDirectory(dir string) bool {
    // If dir variable is not provided it defaults to the shell's HOME environment variable. Alternatively if the user provides a tilde (~), that represents the HOME directory
    if dir == "" || dir == "~" {
        dir = os.Getenv("HOME")
    }

    // If the dir variable is not an absolute path, store the value of CDPATH in the cdPath variable
    if !strings.HasPrefix(dir, "/") {
        cdPath := os.Getenv("CDPATH")
        // If the CDPATH environment variable is not empty, split the variable by colon and iterate over the paths to look for the dir variable
        if cdPath != "" {
            paths := strings.Split(cdPath, ":")
            for _, path := range paths {
                // When CDPATH includes an empty string, it means the current directory, so we set the path variable to "."
                if path == ""{
                    path = "."
				}
                // Join the path and dir variables to create a full path
                fullPath := filepath.Join(path, dir)
				// Clean the path to remove any redundant separators
                fullPath = filepath.Clean(fullPath)
                //fmt.Println("Checking path:", fullPath)
                if _, err := os.Stat(fullPath); err == nil {
                    dir = fullPath
                    break
                }
            }
        }
    }

    // Change directory
    oldPwd, _ := os.Getwd()
    if err := os.Chdir(dir); err != nil {
        fmt.Println("cd: " + dir + ": No such file or directory")
        return false
    }

    // Update PWD and OLDPWD environment variables
    newPwd, _ := os.Getwd()
    os.Setenv("PWD", newPwd)
    os.Setenv("OLDPWD", oldPwd)

    // Print the new directory if a non-empty CDPATH directory was used
    if cdPath := os.Getenv("CDPATH"); cdPath != "" && !strings.HasPrefix(dir, "/") {
        fmt.Println(newPwd)
    }

    return true
}