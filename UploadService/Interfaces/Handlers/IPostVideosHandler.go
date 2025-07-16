package handlerInterfaces

import postVideosModels "uploadservice/Models/PostVideosModels"

type IPostVideosHandler interface {
	Handle(request postVideosModels.PostVideosRequest) error
}
