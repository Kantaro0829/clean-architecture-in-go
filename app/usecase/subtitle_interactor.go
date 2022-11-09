package usecase

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type SubtitleInteractor struct {
	SubtitleRepository SubtitleRepository
}

func (interactor *SubtitleInteractor) GetYoutubeSubtitle(videoId string) (domain.GetSubResp, error) {
	url := "http://get-subtitle-service:5001/get_subtittle"
	subtitles, err := interactor.SubtitleRepository.GetSubtitles(videoId, url)
	if err != nil {
		return domain.GetSubResp{}, nil
	}
	return subtitles, nil
}
