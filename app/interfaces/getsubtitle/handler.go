package getsubtitle

type ApiHandler interface {
	Post(reqBody interface{}, url string) ([]byte, error)
}
