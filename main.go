package main

import (
	"fmt"
	"os"
)

type exitCode int

const (
	exitOK    exitCode = 0
	exitError exitCode = 1
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("error: no command")
		os.Exit(int(exitError))
	}

	command := os.Args[1]
	subArgs := os.Args[2:]
	var err error
	switch command {
	case "clear":
		err = commandClear(subArgs)
	case "edit":
		err = commandEdit(subArgs)
	case "list":
		err = commandList(subArgs)
	case "new":
		err = commandNew(subArgs)
	case "run":
		err = commandRun(subArgs)
	case "show":
		err = commandShow(subArgs)
	default:
		err = fmt.Errorf("unknown command: %s", command)
	}

	code := exitOK
	if err != nil {
		fmt.Printf("error: %s", err)
		code = exitError
	}

	os.Exit(int(code))
}
