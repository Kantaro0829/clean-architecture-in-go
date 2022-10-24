package infrastructure

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type GoQueryHandler struct {
	query *goquery.Document
}

// func NewGoQueryHandler() scrape.ScrapeHandler {

// 	res, err := http.Get("https://ejje.weblio.jp/content/eliminate")
// 	if err != nil {
// 		panic(err.Error)
// 	}
// 	defer res.Body.Close()
// 	if res.StatusCode != 200 {
// 		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
// 	}
// 	goQueryHandler := new(GoQueryHandler)
// 	goQueryHandler.query, _ = goquery.NewDocumentFromReader(res.Body)
// 	return goQueryHandler.query
// }

func (handler *GoQueryHandler) Find(selector string) string {
	baseUrl := "https://ejje.weblio.jp/content/eliminate"
	res, err := http.Get(baseUrl)
	if err != nil {
		panic(err.Error)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	selection := handler.query.Find("div#summary")
	innerSeceltion := selection.Find("p")
	text := innerSeceltion.Text()
	return text
}
