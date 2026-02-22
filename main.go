package main

import (
	"fmt"
	"os"
)

const (
	exitOK = 0
	exitError = 1
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("goscribble is a tool for edit and run Go code instantly on console.\n\n")
		fmt.Print("Usage:\n\n")
		fmt.Print("\tgoscribble <command> [arguments]\n\n")
		fmt.Print("The commands are:\n\n")
		fmt.Println("\tclear  clear all files")
		fmt.Println("\tdir    print directory path")
		fmt.Println("\tedit   edit and run file")
		fmt.Println("\tlist   list all files")
		fmt.Println("\tnew    create and edit file and run it")
		fmt.Println("\trun    run file")
		fmt.Print("\tshow   display file\n\n")
		os.Exit(exitError)
	}

	command := os.Args[1]
	subArgs := os.Args[2:]
	var err error
	switch command {
	case "clear":
		err = clearCommand(subArgs)
	case "dir":
		err = dirCommand(subArgs)
	case "edit":
		err = editCommandEdit(subArgs)
	case "list":
		err = listCommand(subArgs)
	case "new":
		err = newCommand(subArgs)
	case "run":
		err = runCommand(subArgs)
	case "show":
		err = showCommand(subArgs)
	default:
		err = fmt.Errorf("unknown command: %s", command)
	}

	code := exitOK
	if err != nil {
		fmt.Printf("error: %s", err)
		code = exitError
	}

	os.Exit(code)
}
