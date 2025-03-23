package main

import (
	"SparkGuardBackend/internal/db"
	"SparkGuardBackend/internal/repacker"
	"SparkGuardBackend/pkg/s3storage"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const path = "/Users/kerblif/Downloads/790617_1-Проект 3_2 - повторка-1571841"
const eventID = 4

func main() {
	if err := s3storage.Connect(os.Getenv("S3_ENDPOINT"), os.Getenv("S3_REGION"), os.Getenv("S3_BUCKET")); err != nil {
		panic(err)
	}

	err := processArchives(path)
	if err != nil {
		fmt.Printf("Error processing archives: %v\n", err)
	}
}

// processArchives ищет архивы в указанных директориях и вызывает upload для каждого архива.
func processArchives(rootPath string) error {
	// Рекурсивно проходим по всему дереву файлов/директорий.
	return filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Проверка, является ли найденный элемент файлом.
		if !info.IsDir() {
			// Проверяем, является ли файл архивом по расширению.
			if isArchive(path) {
				fmt.Printf("Found archive: %s\n", path) // Для отладки
				// Вызываем функцию загрузки для архива.
				if err = upload(path); err != nil {
					fmt.Printf("Error uploading file %s: %v\n", path, err)
					panic(err)
				}
			}
		}

		return nil
	})
}

// isArchive проверяет, является ли файл архивом по расширению.
func isArchive(filePath string) bool {
	// Список расширений архивных файлов.
	archiveExtensions := []string{".zip", ".tar", ".tar.gz", ".tgz", ".rar", ".7z"}

	for _, ext := range archiveExtensions {
		if strings.HasSuffix(strings.ToLower(filePath), ext) {
			return true
		}
	}

	return false
}

func upload(path string) error {
	parts := strings.Split(path, string(os.PathSeparator))

	workPath := parts[len(parts)-2]
	studentName := strings.Split(workPath, "_")[0]

	student, err := db.GetStudentByName(studentName)

	if err != nil {
		student = &db.Student{
			Name: studentName,
		}
		if err = db.CreateStudent(student); err != nil {
			panic(err)
		}
	}

	work := &db.Work{
		Time:      time.Now(),
		StudentID: student.ID,
		EventID:   eventID,
	}
	if err = db.CreateWork(work); err != nil {
		panic(err)
	}

	// Открываем файл
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	clear_file, err := repacker.Repack(file)

	if err != nil {
		return fmt.Errorf("failed to repack file: %v", err)
	}

	err = s3storage.UploadFileSafe(fmt.Sprintf("./%d/%d.zip", work.EventID, work.ID), clear_file)

	if err != nil {
		return fmt.Errorf("failed to upload file to S3: %v", err)
	}

	fmt.Println("File uploaded successfully")
	return nil
}
