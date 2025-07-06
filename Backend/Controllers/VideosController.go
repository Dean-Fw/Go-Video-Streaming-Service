package controllers

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"videoservice/Interfaces"
)

type VideoController struct {
	FileSystemService interfaces.IFileSystemService
}

func (vc VideoController) GetVideos(w http.ResponseWriter, r *http.Request) {

	video := strings.Split(r.URL.Path, "/")[2]

	intStart, err := getRequestedRange(r.Header)

	amountOfBytes := math.Pow(10, 6)

	file, err := vc.FileSystemService.GetFile(video)

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

func getRequestedRange(header http.Header) (int, error) {
	requestedBytes := header.Get("Range")

	bytesRange := strings.Split(requestedBytes, "=")[1]
	bytesStartAndEnd := strings.Split(bytesRange, "-")

	start, err := strconv.Atoi(bytesStartAndEnd[0])

	if err != nil {
		return -1, err
	}

	return start, nil
}
