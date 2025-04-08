package work

import (
	"SparkGuardBackend/internal/db"
	"archive/zip"
	"bytes"
	"fmt"
	"log"
)

// createAdoptionsArchive создает ZIP-архив с файлами заимствований.
// Каждый файл в архиве соответствует заимствованию workId, а содержимое
// файла включает связанные заимствования.
func createAdoptionsArchive(adoptions []*db.Adoption, workId uint64, processor *SegmentProcessor) (*bytes.Buffer, error) {
	// Группируем заимствования по их ID.
	adoptionGroups := make(map[uint64][]*db.Adoption)

	// Разделяем заимствования для создания отдельных файлов.
	for _, adoption := range adoptions {
		if adoption.WorkID != workId {
			continue
		}

		relatedAdoptions, err := findRelatedAdoptions(adoptions, adoption)

		if err != nil {
			log.Printf("ERROR: Failed to get related adoptions for adoption %d: %v", adoption.ID, err)
			continue
		}

		adoptionGroups[adoption.ID] = relatedAdoptions
	}

	// Создаем ZIP-архив.
	outputZipBuffer := new(bytes.Buffer)
	zipWriter := zip.NewWriter(outputZipBuffer)
	defer zipWriter.Close()

	// Обрабатываем каждую группу заимствований.
	for adoptionID, group := range adoptionGroups {
		err := writeAdoptionGroupToZip(zipWriter, adoptionID, group, processor)
		if err != nil {
			log.Printf("ERROR: Failed to write adoption group %d: %v", adoptionID, err)
			continue
		}
	}

	// Закрываем ZIP и возвращаем готовый архив.
	err := zipWriter.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to finalize zip archive: %w", err)
	}

	return outputZipBuffer, nil
}

// findRelatedAdoptions ищет заимствования, связанные с конкретным workId.
func findRelatedAdoptions(allAdoptions []*db.Adoption, curAdoption *db.Adoption) ([]*db.Adoption, error) {
	refersTo := curAdoption.RefersTo

	var adoptionID uint64 = 0
	if refersTo != nil {
		adoptionID = *refersTo
	}

	var relatedAdoptions = []*db.Adoption{curAdoption}
	for _, adoption := range allAdoptions {
		if (adoption.RefersTo != nil && *adoption.RefersTo == curAdoption.ID) ||
			adoption.ID == adoptionID {
			relatedAdoptions = append(relatedAdoptions, adoption)
		}
	}

	return relatedAdoptions, nil
}

// writeAdoptionGroupToZip записывает группу заимствований в файл внутри ZIP-архива.
func writeAdoptionGroupToZip(zipWriter *zip.Writer, adoptionID uint64, adoptions []*db.Adoption, processor *SegmentProcessor) error {
	// Генерируем имя файла для группы.
	filename := fmt.Sprintf("adoption_%d.txt", adoptionID)
	fileWriter, err := zipWriter.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file in zip archive: %w", err)
	}

	// Пишем данные всех заимствований группы в файл.
	for _, adoption := range adoptions {
		data, err := processor.GetSegmentData(adoption)
		if err != nil {
			log.Printf("WARN: Failed to get segment data for adoption %d: %v. Skipping.", adoption.ID, err)
			continue
		}

		// Записываем метаданные и контент в файл.
		content := fmt.Sprintf("Filename: %s\nMetadata: %s\nContent:\n%s\n\n",
			data.Filename, data.Metadata, string(data.Content))
		_, err = fileWriter.Write([]byte(content))
		if err != nil {
			log.Printf("ERROR: Failed to write content for adoption %d: %v", adoption.ID, err)
			continue
		}
	}

	return nil
}

// createEmptyZip создает пустой zip-архив (для случаев, когда нет сегментов).
func createEmptyZip() *bytes.Buffer {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	// Важно закрыть Writer, чтобы записались корректные заголовки пустого архива
	w.Close()
	return buf
}
