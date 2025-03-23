package repacker

import (
	"os"
	"path/filepath"
)

// RemoveUnwantedFiles удаляет ненужные файлы и директории из указанного пути.
func RemoveUnwantedFiles(root string) error {
	unwanted := []string{".git", ".idea", ".vs", ".DS_Store", "bin", "obj"}

	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if os.IsNotExist(err) {
				return nil
			}
			return err
		}

		for _, unwantedItem := range unwanted {
			if info.Name() == unwantedItem {
				if info.IsDir() {
					if err := os.RemoveAll(path); err != nil && !os.IsNotExist(err) {
						return err
					}
				} else {
					if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
						return err
					}
				}
				break
			}
		}

		return nil
	})
}
