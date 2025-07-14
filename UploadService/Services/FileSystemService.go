package services

import (
	"fmt"
	"io"
	"os"
	"uploadservice/Errors/Services/FileSystemService"
	"uploadservice/Models/Services/FileSystemService"
)

type FileSystemService struct{}

func (FileSystemService) CreateNewFile(FileNameHash string, content []byte) error {
	err := os.WriteFile(fmt.Sprintf("../Videos/%s.mp4", FileNameHash), content, 0644)

	if err != nil {
		return err
	}

	return nil
}

func (FileSystemService) GetFile(fileNameHash string) (filesystemModels.GetFilesResponseModel, error) {
	directory := "../Videos"

	file, err := os.OpenFile(fmt.Sprintf("%s/%s.mp4", directory, fileNameHash), os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		if os.IsNotExist(err) {
			return filesystemModels.GetFilesResponseModel{},
				filesystemServiceErrors.FileNotFoundError{File: fmt.Sprintf("%s.mp4", fileNameHash)}
		} else {
			return filesystemModels.GetFilesResponseModel{},
				err
		}
	}

	fileInfo, err := file.Stat()

	if err != nil {
		return filesystemModels.GetFilesResponseModel{},
			err
	}

	return filesystemModels.GetFilesResponseModel{
		FileWriter: file,
		FileSize:   fileInfo.Size(),
	}, nil
}

func (FileSystemService) UpdateFile(file io.WriteCloser, content []byte) error {
	_, err := file.Write(content)

	if err != nil {
		fmt.Print(err.Error())
	}
	return nil
}
