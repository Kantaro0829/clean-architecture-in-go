package domain

type GetSubReq struct {
	VideoId string `json:"video_id"`
}

type GetSubResp struct {
	Status    uint8  `json:"status"`
	Subtittle string `json:"subtittle"`
}
