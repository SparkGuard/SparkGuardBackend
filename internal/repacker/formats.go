package repacker

import (
	"archive/zip"
	"fmt"
	"github.com/bodgit/sevenzip"
	"github.com/nwaples/rardecode"
	"golang.org/x/text/encoding/charmap"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// decodeFileName декодирует имя файла из CP437 в UTF-8.
func decodeFileName(name string) (string, error) {
	decoder := charmap.CodePage437.NewDecoder() // CP437 → UTF-8
	decodedName, err := decoder.String(name)
	if err != nil {
		return "", fmt.Errorf("failed to decode file name: %s, error: %v", name, err)
	}
	return decodedName, nil
}

func extractZipFile(file *zip.File, destination string) error {
	name, err := decodeFileName(file.Name)
	if err != nil {
		return err
	}

	path := filepath.Join(destination, name)

	// Предотвращение выхода за пределы директории
	if !strings.HasPrefix(path, filepath.Clean(destination)+string(os.PathSeparator)) {
		return os.ErrInvalid
	}

	// Если это директория
	if file.FileInfo().IsDir() {
		return os.MkdirAll(path, os.ModePerm)
	}

	// Если это файл
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}
	defer outFile.Close()

	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	_, err = io.Copy(outFile, rc)
	return err
}

// UnzipZip распаковывает zip-архив.
func UnzipZip(data io.ReaderAt, size int64, destination string) error {
	// Создаем zip.Reader из буфера
	zipReader, err := zip.NewReader(data, size)
	if err != nil {
		return err
	}

	// Распаковка файлов
	for _, file := range zipReader.File {
		err = extractZipFile(file, destination)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

// Unzip7z распаковывает 7z-архив.
func Unzip7z(reader io.ReaderAt, size int64, destination string) error {
	archive, err := sevenzip.NewReader(reader, size)
	if err != nil {
		return err
	}

	for _, file := range archive.File {
		fmt.Printf("Extracting: %s\n", file.Name)

		path := filepath.Join(destination, file.Name)

		// Предотвращение выхода за пределы директории
		if !strings.HasPrefix(filepath.Clean(path), filepath.Clean(destination)+string(os.PathSeparator)) {
			return os.ErrInvalid
		}

		// Если это директория
		if file.FileInfo().IsDir() {
			if err = os.MkdirAll(path, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		// Если это файл
		if err = os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		rc, err := file.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		if _, err = io.Copy(outFile, rc); err != nil {
			rc.Close()
			outFile.Close()
			return err
		}

		rc.Close()
		outFile.Close()
	}
	return nil
}

// UnzipRar распаковывает rar-архив.
func UnzipRar(reader io.Reader, _ int64, destination string) error {
	// Создаем RAR-декодер
	rarReader, err := rardecode.NewReader(reader, "")
	if err != nil {
		return fmt.Errorf("failed to initialize RAR reader: %w", err)
	}

	for {
		// Получаем заголовок файла из архива
		header, err := rarReader.Next()
		if err == io.EOF {
			break // Конец архива
		}
		if err != nil {
			// Если ошибка — логируем и продолжаем
			fmt.Printf("Skipping file due to error: %v\n", err)
			continue
		}

		path := filepath.Join(destination, header.Name)

		// Предотвращение выхода за пределы директории
		if !strings.HasPrefix(path, filepath.Clean(destination)+string(os.PathSeparator)) {
			fmt.Printf("Skipping invalid file path: %s\n", path)
			continue
		}

		// Если это директория, создаем ее
		if header.IsDir {
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				fmt.Printf("Error creating directory %s: %v\n", path, err)
			}
			continue
		}

		// Если это файл, создаем необходимые директории
		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			fmt.Printf("Error creating parent directories for %s: %v\n", path, err)
			continue
		}

		// Создаем файл для записи
		outFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", path, err)
			continue
		}

		// Копируем содержимое файла
		if _, err := io.Copy(outFile, rarReader); err != nil {
			fmt.Printf("Error writing file %s: %v\n", path, err)
			outFile.Close()
			continue
		}

		// Закрываем файл
		outFile.Close()
	}

	return nil
}
