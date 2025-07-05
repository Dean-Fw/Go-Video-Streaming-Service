package main

import (
	"os"
)

func GetFile(fileName string) (*os.File, error) {
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

	return nil, os.ErrNotExist
}
