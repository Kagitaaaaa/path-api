package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type TipeUpload string

const (
	TipeUploadPostingan      TipeUpload = "uploads/postingan/"
	TipeUploadProfilePicture TipeUpload = "uploads/pfp/"
)

func WriteFile(tipeUpload TipeUpload, file *multipart.FileHeader) (string, error) {
	dirPath := filepath.Dir(string(tipeUpload))
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(filepath.Join(dirPath, file.Filename))
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return file.Filename, nil
}
