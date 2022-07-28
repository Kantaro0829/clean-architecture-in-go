package usecase

import (
	"fmt"

	"github.com/Kantaro0829/clean-architecture-in-go/domain"
	//"golang.org/x/crypto/bcrypt"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) error {
	//interactor.UserRepository.Store(u)
	err := interactor.UserRepository.Store(u)
	return err
}

func (interactor *UserInteractor) GetInfo() ([]domain.User, error) {
	return interactor.UserRepository.Select()
}

func (interactor *UserInteractor) Delete(id string) {
	interactor.UserRepository.Delete(id)
}

func (interactor *UserInteractor) Update(u domain.User, name string) {
	interactor.UserRepository.Update(u, name)
}

func (interactor *UserInteractor) Login(mail string, password string) (bool, error) {
	regedPassword, err := interactor.UserRepository.GetPassword(mail)
	if err != nil {
		return false, err
	}
	//パスワード比較
	isValidated := ValitatePassword(regedPassword, password)

	return isValidated, nil
}

func (interactor *UserInteractor) UpdateUser(userJson domain.User) (string, error, bool) {
	mail := userJson.Mail
	password := userJson.Password
	name := userJson.Name
	user, err := interactor.UserRepository.GetPasswordForUpdate(mail)
	if err != nil {
		return "パスワードorメールアドレスの入力間違い", err, false
	}
	regedPassword := user.Password

	//パスワード比較
	isValidated := ValitatePassword(regedPassword, password)
	fmt.Println("パスワード比較完了")
	if isValidated != true {
		return "パスワードorメールアドレスの入力間違い", nil, false
	}

	user.Name = name
	fmt.Println(user)

	result := interactor.UserRepository.UpdateByMail(user)
	if result != nil {
		return "データ書き込み失敗", result, true
	}
	return "登録完了", nil, true
}

func (interactor *UserInteractor) DeleteByMail(userJson domain.User) (string, error, bool) {
	mail := userJson.Mail
	password := userJson.Password
	user, err := interactor.UserRepository.GetPasswordForUpdate(mail)

	if err != nil {
		return "パスワードorメールアドレスの入力間違い", err, false
	}

	regedPassword := user.Password
	isValidated := ValitatePassword(regedPassword, password)
	if isValidated != true {
		return "パスワードorメールアドレスの入力間違い", nil, false
	}

	result := interactor.UserRepository.DeleteByMail(user)
	if result != nil {
		return "データ削除失敗", result, true
	}
	return "削除完了", nil, true

}

func ValitatePassword(regedPassword string, password string) bool {
	if regedPassword == password {
		return true
	}
	return false
}
