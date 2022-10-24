package scrape

type ScrapeHandler interface {
	Find(selector string) string
}
