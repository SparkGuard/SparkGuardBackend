package repacker

import (
	"bytes"
	"errors"
	"io"
	"log"
)

// Unzip распаковывает архив в указанную директорию, поддерживая разные форматы.
func Unzip(data io.Reader, destination string) error {
	buffer := bytes.NewBuffer(nil)
	_, err := io.Copy(buffer, data)
	if err != nil {
		return err
	}

	if err = UnzipZip(bytes.NewReader(buffer.Bytes()), int64(buffer.Len()), destination); err == nil {
		return nil
	} else {
		log.Println("Not zip...")
	}

	if err = Unzip7z(bytes.NewReader(buffer.Bytes()), int64(buffer.Len()), destination); err == nil {
		return nil
	} else {
		log.Println("Not 7z...")
	}

	if err = UnzipRar(bytes.NewReader(buffer.Bytes()), int64(buffer.Len()), destination); err == nil {
		return nil
	} else {
		log.Println("Not rar...")
	}

	return errors.New("unsupported archive format")
}
