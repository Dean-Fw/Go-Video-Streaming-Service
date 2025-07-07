package services

import (
	"fmt"
	"os"
	errors "videoservice/Errors/ServiceErrors"
)

type FileSystemService struct{}

func (FileSystemService) GetFile(fileName string) (*os.File, error) {
	files, err := os.ReadDir("../Videos")

	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.Name() == fileName+".mp4" {
			pFile, err := os.Open("../Videos/" + file.Name())

			if err != nil {
				return nil, err
			}

			return pFile, nil
		}

	}

	return nil, errors.NotFoundError{
		Message: fmt.Sprintf("File %s.mp4 was not found", fileName),
	}
}
