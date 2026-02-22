package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func clearCommand(args []string) error {
	dir, err := baseDir()
	if err != nil {
		return err
	}

	names := []string{scribbleFile, scribbleOldFile, "go.mod", "go.sum"}
	for _, name := range names {
		path := filepath.Join(dir, name)
		if exists(path) {
			if err := os.Remove(path); err != nil {
				return fmt.Errorf("failed to remove file: %s: %w", path, err)
			}
		}
	}

	return nil
}
