package services

import (
	"fmt"
	"os"
)

type FileSystemService struct{}

func (FileSystemService) CreateNewFile(FileNameHash string, content []byte) error {
	fmt.Printf("%s", FileNameHash)
	err := os.WriteFile(fmt.Sprintf("../Videos/%s.mp4", FileNameHash), content, 0644)

	if err != nil {
		return err
	}

	return nil
}
