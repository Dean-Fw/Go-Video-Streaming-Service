package services

import (
	"fmt"
	"io"
	"os"
	"uploadservice/Errors/Services/FileSystemService"
)

type FileSystemService struct{}

func (FileSystemService) CreateNewFile(FileNameHash string, content []byte) error {
	err := os.WriteFile(fmt.Sprintf("../Videos/%s.mp4", FileNameHash), content, 0644)

	if err != nil {
		return err
	}

	return nil
}

func (FileSystemService) GetFile(fileNameHash string) (io.WriteCloser, error) {
	directory := "../Videos"

	file, err := os.OpenFile(fmt.Sprintf("%s/%s.mp4", directory, fileNameHash), os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		if os.IsNotExist(err) {
			return nil, filesystemServiceErrors.FileNotFoundError{File: fmt.Sprintf("%s.mp4", fileNameHash)}
		} else {
			return nil, err
		}
	}

	return file, nil
}

func (FileSystemService) UpdateFile(file io.WriteCloser, content []byte) error {
	_, err := file.Write(content)

	if err != nil {
		fmt.Print(err.Error())
	}
	return nil
}
