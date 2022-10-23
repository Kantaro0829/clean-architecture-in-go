package usecase

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type Tokenizer interface {
	New(domain.User) (domain.Token, error)
	//Verify(domain.Token) error
	Verify(domain.Token) (int, error)
}
