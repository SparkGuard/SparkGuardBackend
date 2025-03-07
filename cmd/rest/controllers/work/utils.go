package work

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

func repack(data io.Reader) (io.Reader, error) {
	tempDir, err := os.MkdirTemp("", "unzip_*")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tempDir)

	// Разархивируем содержимое в новую директорию
	err = unzip(data, tempDir)
	if err != nil {
		return nil, err
	}

	// Удалим ненужные файлы и директории
	err = removeUnwantedFiles(tempDir)
	if err != nil {
		return nil, err
	}

	// Создадим новый zip-архив
	var buffer bytes.Buffer
	err = zipFiles(tempDir, &buffer)
	if err != nil {
		return nil, err
	}

	return &buffer, nil
}

func unzip(data io.Reader, destination string) error {
	// Буферизуем входные данные
	buffer := bytes.NewBuffer(nil)
	_, err := io.Copy(buffer, data)
	if err != nil {
		return err
	}

	// Определяем формат архива по сигнатуре
	archiveBytes := buffer.Bytes()

	if err = unzipZip(archiveBytes, destination); err == nil {
		return nil
	} else {
		log.Println("Not zip...")
	}

	if err = unzip7z(bytes.NewReader(archiveBytes), destination); err == nil {
		return nil
	} else {
		log.Println("Not 7z...")
	}

	if err = unzipRar(bytes.NewReader(archiveBytes), destination); err == nil {
		return nil
	} else {
		log.Println("Not rar...")
	}

	return errors.New("unsupported archive format")
}

func removeUnwantedFiles(root string) error {
	// Список файлов и директорий для удаления
	unwanted := []string{".git", ".idea", ".vs", ".DS_Store", "bin", "obj"}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Игнорируем ошибку отсутствия файла, так как он уже мог быть удален
			if os.IsNotExist(err) {
				return nil
			}
			return err
		}

		// Проверяем, содержит ли имя элемента что-то из unwanted
		for _, unwantedItem := range unwanted {
			if info.Name() == unwantedItem {
				// Удаляем файл или папку
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

	return err
}

func zipFiles(source string, writer io.Writer) error {
	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()

	err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Создаем путь в архиве относительно исходной директории
		relPath, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}

		// Игнорируем корневую директорию
		if relPath == "." {
			return nil
		}

		// Если это директория
		if info.IsDir() {
			_, err = zipWriter.Create(relPath + "/")
			return err
		}

		// Если это файл
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		writer, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, file)
		return err
	})

	return err
}
