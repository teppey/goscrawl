package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func commandEdit(args []string) error {
	dir, err := baseDir()
	if err != nil {
		return err
	}

	path := filepath.Join(dir, "0.go")

	if err := edit(path); err != nil {
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

func edit(path string) error {
	cmd := exec.Command("vim", path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to edit file: %s: %w", cmd, err)
	}

	return nil
}

func goimports(path string) error {
	cmd := exec.Command("goimports", "-w", path)
	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		fmt.Print(string(out))
	}

	return err
}

func run(path string) error {
	cmd := exec.Command("go", "run", path)
	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		fmt.Print(string(out))
	}

	return err
}
