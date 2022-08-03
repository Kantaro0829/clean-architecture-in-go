package token

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type Tokenizer interface {
	New(domain.User) (domain.Token, error)
	Verify(domain.Token) error
}

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

func (tokenizer *TokenizerImpl) Verify(token domain.Token) error {
	tokenString := string(token)
	err := tokenizer.VerifyToken(tokenString)
	return err
}
