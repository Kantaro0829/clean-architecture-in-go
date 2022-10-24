package scrape

import (
	"fmt"
)

type ScrapeRepository struct {
	ScrapeHandler
}

func (scrape *ScrapeRepository) FindMeaning(word string) string {

	meaning := scrape.Find(word)
	fmt.Println(meaning)
	return meaning
}
