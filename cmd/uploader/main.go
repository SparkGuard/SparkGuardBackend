package main

import (
	"SparkGuardBackend/internal/db"
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const path = "/Users/kerblif/Downloads/790617_1-Проект 3_2-1561132"
const url = "http://localhost:8080/works/%d/upload"
const eventID = 2

func main() {
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

	// Создаем Buffer и экземпляр Writer для multipart/form-data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Добавляем файл в форму
	part, err := writer.CreateFormFile("file", path)
	if err != nil {
		return fmt.Errorf("failed to create form file: %v", err)
	}

	// Копируем содержимое файла в форму
	if _, err := io.Copy(part, file); err != nil {
		return fmt.Errorf("failed to copy file content: %v", err)
	}

	// Завершаем запись данных
	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close writer: %v", err)
	}

	// Создаем запрос
	req, err := http.NewRequest("PUT", fmt.Sprintf(url, work.ID), body)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("accept", "application/json")

	// Отправляем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем ответ
	if resp.StatusCode != http.StatusOK {
		tmp, _ := io.ReadAll(resp.Body)
		log.Println(string(tmp))
		return fmt.Errorf("request failed with status: %s", resp.Status)
	}

	fmt.Println("File uploaded successfully")
	return nil
}
