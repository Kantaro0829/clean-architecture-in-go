package database

import (
	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type UserRepository struct {
	SqlHandler
}

//DB操作のための関数
//同階層の./interface/database/user_repogitory.goから呼び出している

//本来ここで詳細な操作をするべき？ 構造体の指定などもここ？
func (db *UserRepository) Store(u domain.User) {
	db.Create(&u)
}

func (db *UserRepository) Select() ([]domain.User, error) {
	user := []domain.User{}
	users, err := db.FindAll(user)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (db *UserRepository) Delete(id string) {
	user := []domain.User{}
	db.DeleteById(&user, id)
}
func (db *UserRepository) Update(u domain.User, name string) {
	db.UpdateById(u, name)
}
