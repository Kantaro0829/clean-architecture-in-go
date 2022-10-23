package token

type TokenHandler interface {
	Generate(int, string, string) (string, error)
	VerifyToken(string) (int, error)
	//VerifyToken(string) error
}
