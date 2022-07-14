package database

import "github.com/Kantaro0829/clean-architecture-in-go/domain"

type SqlHandler interface {
	Create(object interface{})
	FindAll(object []domain.User) ([]domain.User, error)
	DeleteById(object interface{}, id string)
	UpdateById(object domain.User, name string)
}
