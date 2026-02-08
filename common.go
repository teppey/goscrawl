package main

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
)

func baseDir() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", errors.New("failed to get current user")
	}

	return filepath.Join(os.TempDir(), "goplay_"+u.Username), nil
}
