package database

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type SqlHandler interface {
	Create(user domain.User) error
	FindAll(object []domain.User) ([]domain.User, error)
	DeleteById(object interface{}, id string)
	UpdateName(user domain.User) error
	GetPasswordByMail(mail string) (string, error)
	GetMailNamePasswordByMail(mail string) (domain.User, error)
	GetPasswordAndId(mail string) (domain.User, error)
	DeleteOne(user domain.User) error
	//GetUserForUpdate(user domain.User) (domain.User, error)
}
