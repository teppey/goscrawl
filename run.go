package main

import (
	"path/filepath"
)

func runCommand(args []string) error {
	dir, err := baseDir()
	if err != nil {
		return err
	}

	if err := goModInit(dir); err != nil {
		return err
	}

	if err := goModTidy(dir); err != nil {
		return err
	}

	path := filepath.Join(dir, scribbleFile)
	if err := goimports(path); err != nil {
		return err
	}

	if err := run(path); err != nil {
		return err
	}

	return nil
}
