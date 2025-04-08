package work

import (
	"SparkGuardBackend/internal/db"
	"archive/zip"
	"bytes"
	"fmt"
	"log"
)

// createAdoptionsArchive создает ZIP-архив с сегментами.
func createAdoptionsArchive(adoptions []*db.Adoption, processor *SegmentProcessor) (*bytes.Buffer, error) {
	outputZipBuffer := new(bytes.Buffer)
	zipWriter := zip.NewWriter(outputZipBuffer)
	defer zipWriter.Close()

	processedCount := 0
	for _, adoption := range adoptions {
		data, err := processor.GetSegmentData(adoption)
		if err != nil {
			log.Printf("WARN: Failed to get segment data for adoption %d: %v. Skipping.", adoption.ID, err)
			continue // Пропускаем этот сегмент, но продолжаем с другими
		}

		// Пишем в выходной ZIP
		fileWriter, err := zipWriter.Create(data.Filename)
		if err != nil {
			log.Printf("ERROR: Failed creating file '%s' in output zip for adoption %d: %v", data.Filename, adoption.ID, err)
			// Можно решить остановить процесс или продолжить
			continue
		}

		_, err = fileWriter.Write([]byte(data.Metadata))
		if err != nil {
			log.Printf("ERROR: Failed writing metadata to '%s' for adoption %d: %v", data.Filename, adoption.ID, err)
			continue
		}

		_, err = fileWriter.Write(data.Content)
		if err != nil {
			log.Printf("ERROR: Failed writing content to '%s' for adoption %d: %v", data.Filename, adoption.ID, err)
			continue
		}
		processedCount++
	}

	log.Printf("Processed %d adoption segments for the archive.", processedCount)

	// Закрываем zipWriter (через defer) перед возвратом буфера
	err := zipWriter.Close() // Важно проверить ошибку при закрытии
	if err != nil {
		return nil, fmt.Errorf("failed to finalize output zip archive: %w", err)
	}

	return outputZipBuffer, nil
}

// createEmptyZip создает пустой zip-архив (для случаев, когда нет сегментов).
func createEmptyZip() *bytes.Buffer {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	// Важно закрыть Writer, чтобы записались корректные заголовки пустого архива
	w.Close()
	return buf
}
