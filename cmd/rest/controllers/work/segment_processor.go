package work

import (
	"SparkGuardBackend/internal/db"
	"SparkGuardBackend/pkg/s3storage"
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"
)

// segmentData содержит метаданные и контент извлеченного сегмента.
type segmentData struct {
	Metadata string
	Content  []byte
	Filename string // Имя файла для записи в итоговый ZIP
}

// SegmentProcessor управляет загрузкой архивов из S3 и извлечением сегментов.
type SegmentProcessor struct {
	s3Cache map[uint64]*bytes.Buffer // Кэш для скачанных архивов (WorkID -> zipData)
}

// NewSegmentProcessor создает новый процессор сегментов.
func NewSegmentProcessor() *SegmentProcessor {
	return &SegmentProcessor{
		s3Cache: make(map[uint64]*bytes.Buffer),
	}
}

// GetSegmentData извлекает данные конкретного сегмента по записи Adoption.
// Возвращает метаданные, контент сегмента, имя файла для ZIP и ошибку.
func (sp *SegmentProcessor) GetSegmentData(adoption *db.Adoption) (*segmentData, error) {
	if adoption.Path == nil || adoption.PartOffset == nil || adoption.PartSize == nil || *adoption.PartSize == 0 {
		return nil, fmt.Errorf("adoption ID %d: missing path, segment info, or zero size", adoption.ID)
	}

	segmentWorkID := adoption.WorkID
	sourceZipData, err := sp.getSourceZipData(segmentWorkID)
	if err != nil {
		return nil, fmt.Errorf("adoption ID %d: failed to get source zip for work %d: %w", adoption.ID, segmentWorkID, err)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(sourceZipData.Bytes()), int64(sourceZipData.Len()))
	if err != nil {
		return nil, fmt.Errorf("adoption ID %d: failed to read source archive for work %d: %w", adoption.ID, segmentWorkID, err)
	}

	*adoption.Path = strings.ReplaceAll(*adoption.Path, "\\", "/")

	for _, file := range zipReader.File {
		if strings.ReplaceAll(file.Name, "\\", "/") == *adoption.Path {
			if file.FileInfo().IsDir() {
				return nil, fmt.Errorf("adoption ID %d: path '%s' is a directory", adoption.ID, *adoption.Path)
			}

			content, err := extractSegmentFromFile(file, *adoption.PartOffset, *adoption.PartSize)
			if err != nil {
				return nil, fmt.Errorf("adoption ID %d: failed to extract segment from '%s': %w", adoption.ID, *adoption.Path, err)
			}

			meta := formatMetadata(adoption)
			filename := generateOutputFilename(adoption)

			return &segmentData{
				Metadata: meta,
				Content:  content,
				Filename: filename,
			}, nil
		}
	}

	return nil, fmt.Errorf("adoption ID %d: file '%s' not found in archive for work %d", adoption.ID, *adoption.Path, segmentWorkID)
}

// getSourceZipData получает данные ZIP-архива из кэша или скачивает из S3.
func (sp *SegmentProcessor) getSourceZipData(workID uint64) (*bytes.Buffer, error) {
	if cachedData, found := sp.s3Cache[workID]; found {
		return cachedData, nil
	}

	workData, err := db.GetWork(uint(workID))
	if err != nil {
		return nil, fmt.Errorf("work %d not found in DB: %w", workID, err)
	}

	s3Key := fmt.Sprintf("./%d/%d.zip", workData.EventID, workData.ID)
	downloadedData, err := s3storage.DownloadFileToMemory(s3Key)
	if err != nil {
		return nil, fmt.Errorf("failed to download archive %s from S3: %w", s3Key, err)
	}

	sp.s3Cache[workID] = downloadedData
	return downloadedData, nil
}

// extractSegmentFromFile извлекает конкретный сегмент из zip.File.
func extractSegmentFromFile(file *zip.File, charOffset uint64, charSize uint64) ([]byte, error) {
	if charSize == 0 {
		return []byte{}, nil
	}

	rc, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("could not open file in zip: %w", err)
	}

	defer rc.Close()

	fileBytes, err := io.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("could not read file content: %w", err)
	}

	if !utf8.Valid(fileBytes) {
		return nil, errors.New("file content is not valid UTF-8")
	}

	fileRunes := []rune(string(fileBytes))
	totalRunes := len(fileRunes)

	startCharIndex := int(charOffset)
	numCharsToRead := int(charSize)

	if startCharIndex < 0 {
		return nil, fmt.Errorf("invalid character offset: %d must be non-negative", startCharIndex)
	}

	if startCharIndex >= totalRunes {
		return nil, fmt.Errorf("invalid character offset: %d is beyond file length %d", startCharIndex, totalRunes)
	}

	if numCharsToRead <= 0 {
		return []byte{}, nil
	}

	endCharIndex := startCharIndex + numCharsToRead
	if endCharIndex > totalRunes {
		endCharIndex = totalRunes
		numCharsToRead = endCharIndex - startCharIndex
	}

	if numCharsToRead <= 0 {
		return []byte{}, nil
	}

	segmentRunes := fileRunes[startCharIndex:endCharIndex]

	segmentBytes := []byte(string(segmentRunes))

	return segmentBytes, nil
}

// formatMetadata форматирует метаданные для записи в файл.
func formatMetadata(adoption *db.Adoption) string {
	return fmt.Sprintf(
		"--- Adoption Metadata ---\n"+
			"Adoption ID: %d\n"+
			"Original Work ID: %d\n"+
			"Original File Path: %s\n"+
			"Segment Offset: %d\n"+
			"Segment Size: %d\n"+
			"Verdict: %s\n"+
			"AI Generated: %t\n"+
			"Similarity Score: %.4f\n"+
			"Refers To Adoption ID: %s\n"+
			"Description: %s\n"+
			"--- Segment Content Starts Below ---\n\n",
		adoption.ID,
		adoption.WorkID,
		*adoption.Path,
		*adoption.PartOffset,
		*adoption.PartSize,
		adoption.Verdict,
		adoption.IsAIGenerated,
		adoption.SimilarityScore,
		formatOptionalUint64(adoption.RefersTo),
		adoption.Description,
	)
}

// generateOutputFilename создает имя файла для итогового ZIP.
func generateOutputFilename(adoption *db.Adoption) string {
	safePath := strings.ReplaceAll(*adoption.Path, "/", "_")
	safePath = strings.ReplaceAll(safePath, "\\", "_")

	return fmt.Sprintf("adoption_%d_work_%d_%s.txt", adoption.ID, adoption.WorkID, safePath)
}

// Вспомогательная функция для форматирования опционального uint64*
func formatOptionalUint64(ptr *uint64) string {
	if ptr == nil {
		return "None"
	}
	return strconv.FormatUint(*ptr, 10)
}
