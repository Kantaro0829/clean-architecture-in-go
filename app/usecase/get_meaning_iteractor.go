package usecase

type GetMeaningInteractor struct {
	ScrapeRepository ScrapeRepository
	// //以下試しに追加
	// //Tokenizer token.Tokenizer
	// Tokenizer Tokenizer
}

func (interactor *GetMeaningInteractor) GetTranslatedMeaning(word string) string {

	meaning := interactor.ScrapeRepository.FindMeaning(word)
	return meaning
}
