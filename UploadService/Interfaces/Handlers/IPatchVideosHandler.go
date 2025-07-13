package handlerInterfaces

import models "uploadservice/Models/PatchVideosModels"

type IPatchVideosHandler interface {
	Handle(request models.PatchVideosRequest) (string, error)
}
