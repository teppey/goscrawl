package main

import (
	"fmt"
	"os"
)

func listCommand(args []string) error {
	dir, err := baseDir()
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fmt.Println(entry.Name())
	}

	return nil
}
