package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"videoservice/Errors/HandlerErrors"
	"videoservice/Errors/ServiceErrors"
	"videoservice/Interfaces"
)

type GetVideosHandler struct {
	FileSystemService interfaces.IFileSystemService
}

func (vh GetVideosHandler) Handle(w http.ResponseWriter, r *http.Request) {
	video := r.PathValue("name") // Does this create dependency on the implementation of the route?

	byteRangeStart, err := getRequestedByteRange(r.Header)

	if err != nil {
		switch err.(type) {
		case handlererrors.RangeHeaderMalformed:
			http.Error(w, err.Error(), 400)
			return

		default:
			http.Error(w, "Failed to process request", 500)
			return
		}
	}

	amountOfBytes := math.Pow(10, 6)

	file, err := vh.FileSystemService.GetFile(video)

	if err != nil {
		switch err.(type) {
		case serviceerrors.NotFoundError:
			http.Error(w, err.Error(), 404)
			return

		default:
			http.Error(w, fmt.Sprintf("Failed to read file %s.mp4", video), 500)
			return
		}
	}

	bytes := make([]byte, int(amountOfBytes))

	fileInfo, err := file.Stat()

	if err != nil {
		http.Error(w, "Cannot get information about the requested file", 500)
		return
	}

	file.ReadAt(bytes, int64(byteRangeStart))

	end := int(math.Min(float64(byteRangeStart)+amountOfBytes, float64(fileInfo.Size()-1)))

	writeResponse(w, bytes, byteRangeStart, end, fileInfo.Size())

	file.Close()

}

func getRequestedByteRange(header http.Header) (int, error) {
	requestedBytes := header.Get("Range")

	if requestedBytes == "" {
		return 0, handlererrors.RangeHeaderMalformed{
			Message: "Range Header is missing",
		}
	}

	if !strings.Contains(requestedBytes, "bytes=") {
		return 0, handlererrors.RangeHeaderMalformed{
			Message: "Range Header does not specify bytes",
		}
	}

	byteRange := strings.TrimLeft(requestedBytes, "bytes=")

	bytesRangeStringSlice := strings.Split(byteRange, "-")
	bytesRangeIntSlice, err := strconv.Atoi(bytesRangeStringSlice[0])

	if err != nil {
		return 0, handlererrors.RangeHeaderMalformed{
			Message: "First part of range header was not parsable as an integer",
		}
	}

	return bytesRangeIntSlice, nil
}

func writeResponse(w http.ResponseWriter, bytes []byte, contentStart int, contentEnd int, fileSize int64) {
	w.Header().Add("Content-Type", "video/mp4")
	w.Header().Add("Content-Range", fmt.Sprintf("bytes %d-%d/%d", contentStart, contentEnd, fileSize))
	w.Header().Add("Accept-Ranges", "bytes")

	w.WriteHeader(206)

	w.Write(bytes)
}
