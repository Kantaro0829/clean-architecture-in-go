package usecase

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type SubtitleRepository interface {
	GetSubtitles(video_id string, url string) (domain.GetSubResp, error)
}
