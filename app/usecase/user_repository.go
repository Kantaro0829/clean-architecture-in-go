package usecase

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type UserRepository interface {
	Store(domain.User)
	Select() []domain.User
	Delete(id string)
}
