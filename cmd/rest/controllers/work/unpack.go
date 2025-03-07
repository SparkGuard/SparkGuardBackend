package work

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/bodgit/sevenzip"
	"github.com/nwaples/rardecode"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Распаковка ZIP-архива
func unzipZip(data []byte, destination string) error {
	reader := bytes.NewReader(data)

	// Создаем zip.Reader из буфера
	zipReader, err := zip.NewReader(reader, int64(len(data)))
	if err != nil {
		return err
	}

	// Распаковка файлов
	for _, file := range zipReader.File {
		err := extractZipFile(file, destination)
		if err != nil {
			return err
		}
	}

	return nil
}

func extractZipFile(file *zip.File, destination string) error {
	path := filepath.Join(destination, file.Name)

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

// Распаковка 7z-архива
func unzip7z(reader io.ReaderAt, destination string) error {
	archive, err := sevenzip.NewReader(reader, int64(reader.(interface{ Len() int }).Len()))
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
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		// Если это файл
		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
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

// Распаковка RAR-архива
func unzipRar(reader io.Reader, destination string) error {
	rarReader, err := rardecode.NewReader(reader, "")
	if err != nil {
		return err
	}

	for {
		header, err := rarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		path := filepath.Join(destination, header.Name)

		// Предотвращение выхода за пределы директории
		if !strings.HasPrefix(path, filepath.Clean(destination)+string(os.PathSeparator)) {
			return os.ErrInvalid
		}

		// Если это директория
		if header.IsDir {
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		// Если это файл
		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return err
		}

		if _, err := io.Copy(outFile, rarReader); err != nil {
			outFile.Close()
			return err
		}

		outFile.Close()
	}

	return nil
}
