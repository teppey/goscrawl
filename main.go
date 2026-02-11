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
		fmt.Println("gosketch is a tool for edit and run a Go code instantly like the \"Go Playgrond\".\n")
		fmt.Println("Usage:\n")
		fmt.Println("\tgosketch <command> [arguments]\n")
		fmt.Println("The commands are:\n")
		fmt.Println("\tclear  clear all sketches")
		fmt.Println("\tedit   edit a sketch")
		fmt.Println("\tlist   list sketches")
		fmt.Println("\tnew    create and edit a sketch and run it")
		fmt.Println("\trun    run a sketch")
		fmt.Println("\tshow   display a sketch\n")
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
