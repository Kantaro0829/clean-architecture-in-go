package scrape

type ScrapeHandler interface {
	Find(word string) string
}
