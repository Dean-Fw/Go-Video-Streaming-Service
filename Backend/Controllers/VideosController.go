package controllers

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type VideoController struct {
}

func (VideoController) GetVideos(w http.ResponseWriter, r *http.Request) {

	video := strings.Split(r.URL.Path, "/")[2]

	requestedBytes := r.Header.Get("Range")

	bytesRange := strings.Split(requestedBytes, "=")[1]
	bytesStartAndEnd := strings.Split(bytesRange, "-")

	intStart, err := strconv.Atoi(bytesStartAndEnd[0])

	if err != nil {
		fmt.Fprint(w, "Failed to read range header")
	}

	amountOfBytes := math.Pow(10, 6)

	file, err := GetFile(video)

	if err == os.ErrNotExist {
		fmt.Fprintf(w, "404 Could not find file: %s.mp4\n", video)
		return
	} else if err != nil {
		log.Fatalf("Failed to read file: %s", err.Error())
	}

	bytes := make([]byte, int(amountOfBytes))

	fileInfo, _ := file.Stat()

	file.ReadAt(bytes, int64(intStart))

	end := int(math.Min(float64(intStart)+amountOfBytes, float64(fileInfo.Size()-1)))

	w.Header().Add("Content-Type", "video/mp4")
	w.Header().Add("Content-Range", fmt.Sprintf("bytes %d-%d/%d", intStart, end, fileInfo.Size()))
	w.Header().Add("Accept-Ranges", "bytes")

	w.WriteHeader(206)

	w.Write(bytes)

	file.Close()

}

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
