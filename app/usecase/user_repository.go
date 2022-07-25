package usecase

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type UserRepository interface {
	Store(domain.User) error
	Select() ([]domain.User, error)
	Delete(id string)
	DeleteByMail(user domain.User) error
	Update(u domain.User, name string)
	UpdateByMail(user domain.User) error
	GetPassword(mail string) (string, error)
	GetPasswordForUpdate(mail string) (domain.User, error)
	//GetUserForUpdate(user domain.User) (domain.User, error)
}
