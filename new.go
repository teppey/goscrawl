package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const template = `package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
`

func commandNew(args []string) error {
	dir, err := baseDir()
	if err != nil {
		return err
	}

	if err := os.Mkdir(dir, 0700); err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create directory: %s: %w", dir, err)
	}

	num := 0
	if err := rotate(dir, num); err != nil {
		return fmt.Errorf("failed to rotate: %s: %w", dir, err)
	}

	newPath := filepath.Join(dir, fmt.Sprintf("%d.go", num))
	file, err := os.Create(newPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %s: %w", newPath, err)
	}
	defer file.Close()

	io.WriteString(file, template)

	if err := edit(newPath); err != nil {
		return err
	}

	if err := goimports(newPath); err != nil {
		return err
	}

	if err := run(newPath); err != nil {
		return err
	}

	return nil
}

func rotate(dir string, n int) error {
	return nil

	//		const limit = 9
	//
	//		if n > limit {
	//			panic(fmt.Sprintf("over limit: dir=%s, limit=%d, n=%d", dir, limit, n))
	//		}
	//
	//		cur := filepath.Join(dir, fmt.Sprintf("%d.go", n))
	//		if n == limit {
	//			// TODO
	//		}
	//
	//		old := filepath.Join(dir, fmt.Sprintf("%d.go", n+1))
	//
	//		if _, err := os.Stat(old); errors.Is(err, os.ErrNotExist) {
	//			err := os.Rename(cur, old)
	//			if err != nil {
	//				return err
	//			}
	//		}
	//
	//		err := rotate(dir, n+1)
	//
	//		_ = cur
	//		_ = old
	//
	//		return nil
	//	}
	//
	//	func exists(path string) (bool, error) {
	//		_, err := os.Stat(path)
	//		if err != nil {
	//			if errors.Is(err, os.ErrNotExist) {
	//				return false, nil
	//			} else {
	//				return false, err
	//			}
	//
	//		} else {
	//			return true, nil
	//		}
}
