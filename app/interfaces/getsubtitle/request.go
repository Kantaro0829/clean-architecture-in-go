package getsubtitle

import (
	"encoding/json"

	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type ApiRequestRepository struct {
	ApiHandler
}

func (request *ApiRequestRepository) GetSubtitles(videoId string, url string) (domain.GetSubResp, error) {
	reqBody := &domain.GetSubReq{
		Body: videoId,
	}
	resBody := domain.GetSubResp{}

	resp, err := request.Post(reqBody, url)
	if err != nil {
		return resBody, err
	}
	json.Unmarshal(resp, &resBody)

	return resBody, nil
}
