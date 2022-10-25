package infrastructure

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Kantaro0829/clean-architecture-in-go/interfaces/scrape"
	"github.com/PuerkitoBio/goquery"
)

type GoQueryHandler struct {
	query *goquery.Document
}

func NewGoQueryHandler() scrape.ScrapeHandler {

	res, err := http.Get("https://ejje.weblio.jp/content/eliminate")
	if err != nil {
		panic(err.Error)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	goQueryHandler := new(GoQueryHandler)
	goQueryHandler.query, _ = goquery.NewDocumentFromReader(res.Body)
	return goQueryHandler
}

func (handler *GoQueryHandler) Find(word string) string {
	baseUrl := "https://ejje.weblio.jp/content/"
	baseUrl += word

	res, err := http.Get(baseUrl)
	if err != nil {
		panic(err.Error)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	//fmt.Printf("url：%v", handler.query.Url.Path)
	handler.query, _ = goquery.NewDocumentFromReader(res.Body)
	fmt.Println(handler.query.Url)

	selection := handler.query.Find("div#summary")
	innerSeceltion := selection.Find("p")
	text := innerSeceltion.Text()
	result := strings.Replace(text, " ", "", -1)
	arr := strings.Split(result, "\n")
	var final string
	for i, s := range arr {
		if i == 4 {
			final = s
		}
		fmt.Printf("%s\n", s) // 赤 黄 青
		fmt.Println(i)
	}
	fmt.Println(len(arr)) // 3

	fmt.Printf("編集後のテキスト：%v", result)
	return final
}
