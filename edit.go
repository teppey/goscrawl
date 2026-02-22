package main

import (
	"path/filepath"
)

func editCommandEdit(args []string) error {
	dir, err := baseDir()
	if err != nil {
		return err
	}

	path := filepath.Join(dir, scribbleFile)

	if err := edit(path); err != nil {
		return err
	}

	if err := goModTidy(dir); err != nil {
		return err
	}

	if err := goimports(path); err != nil {
		return err
	}

	if err := run(path); err != nil {
		return err
	}

	return nil
}
