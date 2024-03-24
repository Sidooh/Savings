package utils

import (
	"log/slog"
	"os"
	"path/filepath"
)

func GetFile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}
	//defer file.Close()

	return file
}

func CreateFile(path string) *os.File {
	err := os.MkdirAll(filepath.Dir(path), 0660)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	file, err := os.Create(path)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}
	defer file.Close()

	return file
}
