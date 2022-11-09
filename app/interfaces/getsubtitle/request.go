package getsubtitle

import (
	"encoding/json"
	"fmt"

	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type ApiRequestRepository struct {
	ApiHandler
}

func (request *ApiRequestRepository) GetSubtitles(videoId string, url string) (domain.GetSubResp, error) {
	reqBody := &domain.GetSubReq{
		VideoId: videoId,
	}

	resp, err := request.Post(reqBody, url)
	resBody := domain.GetSubResp{}
	if err != nil {
		return resBody, err
	}

	json.Unmarshal(resp, &resBody)
	fmt.Printf("受け取った字幕%v", resBody)

	return resBody, nil
}
