package helper

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// SaveFile menyimpan file upload dan mengembalikan path relatifnya
func SaveFile(fileHeader *multipart.FileHeader, folder string) (string, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	filename := fileHeader.Filename
	savePath := filepath.Join("uploads", folder, filename)

	if err := os.MkdirAll(filepath.Dir(savePath), 0755); err != nil {
		return "", err
	}

	dst, err := os.Create(savePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return "/" + savePath, nil
}
