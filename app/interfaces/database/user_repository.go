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
func (db *UserRepository) Store(u domain.User) error {

	err := db.Create(u)
	if err != nil {
		return err
	}
	return nil
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
	//db.UpdateById(u, name)
	return
}

func (db *UserRepository) GetMailNamePasswordByMail(mail string) (domain.User, error) {
	result, err := db.GetMailNamePasswordByMail(mail)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (db *UserRepository) UpdateByMail(user domain.User) error {

	err := db.UpdateName(user)

	if err != nil {
		return err
	}
	return nil
}

func (db *UserRepository) GetPassword(mail string) (string, error) {
	passwword, err := db.GetPasswordByMail(mail)

	if err != nil {
		return "", err
	}
	return passwword, nil

}

func (db *UserRepository) GetPasswordForUpdate(mail string) (domain.User, error) {
	user, err := db.GetPasswordAndId(mail)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (db *UserRepository) DeleteByMail(user domain.User) error {

	err := db.DeleteOne(user)

	if err != nil {
		return err
	}
	return nil
}
