package usecase

type ScrapeRepository interface {
	FindMeaning(word string) string
}
