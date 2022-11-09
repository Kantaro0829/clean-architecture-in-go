package getsubtitle

type ApiHandler interface {
	Post(body interface{}, url string) ([]byte, error)
}
