package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

func createFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

func writeFile(path string, content []byte) error {
	dir := filepath.Dir(path)
	if err := createFolder(dir); err != nil {
		return fmt.Errorf("gagal buat folder %s: %w", dir, err)
	}

	return os.WriteFile(path, content, 0644)
}
