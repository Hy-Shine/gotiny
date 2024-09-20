package image

import (
	"encoding/base64"
	"io"
	"os"
)

func Write(reader io.Reader, storePath string) error {
	f, err := os.Create(storePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.ReadFrom(reader)
	return err
}

func WriteString(content, filePath string) error {
	imageBytes := []byte(content)
	err := os.WriteFile(filePath, imageBytes, 0o755)
	return err
}

func WriteBytes(imageBytes []byte, filePath string) error {
	err := os.WriteFile(filePath, imageBytes, 0o755)
	return err
}

func ToBase64(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// read file content
	imageData, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	// encode to base64
	base64String := base64.StdEncoding.EncodeToString(imageData)
	return base64String, nil
}
