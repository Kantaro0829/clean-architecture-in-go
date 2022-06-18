package database

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type UserRepository struct {
	SqlHandler
}

//DB操作のための関数
//同階層の./interface/database/user_repogitory.goから呼び出している
func (db *UserRepository) Store(u domain.User) {
	db.Create(&u)
}

func (db *UserRepository) Select() []domain.User {
	user := []domain.User{}
	db.FindAll(&user)
	return user
}
func (db *UserRepository) Delete(id string) {
	user := []domain.User{}
	db.DeleteById(&user, id)
}
func (db *UserRepository) Update(u domain.User, name string) {
	db.UpdateById(u, name)
}
