package controller

import (
	"fmt"
	"net/http"

	//"github.com/Kantaro0829/clean-architecture-in-go/interfaces/database"
	//"github.com/Kantaro0829/clean-architecture-in-go/interfaces/token"
	"github.com/Kantaro0829/clean-architecture-in-go/interfaces/scrape"

	"github.com/Kantaro0829/clean-architecture-in-go/usecase"
	"github.com/gin-gonic/gin"
)

type WordController struct {
	Interactor usecase.GetMeaningInteractor
}

func NewWordController(
	goQueryHandler scrape.ScrapeHandler,
	//sqlHandler database.SqlHandler,
	//以下追加
	//tokenHandler token.TokenHandler,
) *WordController {
	return &WordController{
		Interactor: usecase.GetMeaningInteractor{
			ScrapeRepository: &scrape.ScrapeRepository{
				ScrapeHandler: goQueryHandler,
			},
		},
	}
}

func (controller *WordController) GetMeaning(c *gin.Context) {
	englishword := c.Param("word")
	res := controller.Interactor.GetTranslatedMeaning(englishword)
	fmt.Println("取得した単語の意味:%v", res)

	c.JSON(http.StatusOK, gin.H{"message": res})
}
