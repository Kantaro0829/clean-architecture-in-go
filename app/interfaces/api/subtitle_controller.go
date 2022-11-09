package controller

import (
	"fmt"
	"net/http"

	//"github.com/Kantaro0829/clean-architecture-in-go/interfaces/database"
	//"github.com/Kantaro0829/clean-architecture-in-go/interfaces/token"
	"github.com/Kantaro0829/clean-architecture-in-go/interfaces/getsubtitle"

	"github.com/Kantaro0829/clean-architecture-in-go/usecase"
	"github.com/gin-gonic/gin"
)

type SubtitleController struct {
	Interactor usecase.SubtitleInteractor
}

func NewSubtitleController(
	apiHandler getsubtitle.ApiHandler,
) *SubtitleController {
	return &SubtitleController{
		Interactor: usecase.SubtitleInteractor{
			SubtitleRepository: &getsubtitle.ApiRequestRepository{
				ApiHandler: apiHandler,
			},
		},
	}
}

func (controller *SubtitleController) GetSubtitle(c *gin.Context) {
	videoId := c.Param("videoId")
	res, err := controller.Interactor.GetYoutubeSubtitle(videoId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "字幕を取得できませんでした！"})
	}
	fmt.Printf("videoId : %v", videoId)
	fmt.Printf("取得した字幕:%v", res)

	c.JSON(http.StatusOK, res)
}
