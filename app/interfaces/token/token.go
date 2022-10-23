package token

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

// type Tokenizer interface {
// 	New(domain.User) (domain.Token, error)
// 	Verify(domain.Token) error
// }

type TokenizerImpl struct {
	TokenHandler
}

func (tokenizer *TokenizerImpl) New(user domain.User) (domain.Token, error) {
	var tokenString domain.Token
	generatedToken, err := tokenizer.Generate(user.ID, user.Name, user.Mail)
	if err != nil {
		return tokenString, err
	}
	tokenString = domain.Token(generatedToken)
	return tokenString, nil
}

//戻り値error から(int error)に変更
func (tokenizer *TokenizerImpl) Verify(token domain.Token) (int, error) {
	tokenString := string(token)
	id, err := tokenizer.VerifyToken(tokenString)
	return id, err
}
