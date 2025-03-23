package repacker

import (
	"bytes"
	"io"
	"os"
)

// Repack принимает архив, распаковывает, удаляет ненужные файлы, и создает новый архив.
func Repack(data io.Reader) (io.Reader, error) {
	tempDir, err := os.MkdirTemp("", "unzip_*")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tempDir)

	// Разархивируем содержимое в новую директорию
	err = Unzip(data, tempDir)
	if err != nil {
		return nil, err
	}

	// Удаляем ненужные файлы и директории
	err = RemoveUnwantedFiles(tempDir)
	if err != nil {
		return nil, err
	}

	// Создаем новый zip-архив
	var buffer bytes.Buffer
	err = ZipFiles(tempDir, &buffer)
	if err != nil {
		return nil, err
	}

	return &buffer, nil
}
