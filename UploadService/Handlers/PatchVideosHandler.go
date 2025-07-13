package handlers

import (
	"uploadservice/Models/PatchVideosModels"
)

type PatchVideosHandler struct{}

func (PatchVideosHandler) Handle(request models.PatchVideosRequest) (string, error) {
	return "Hello", nil
}
