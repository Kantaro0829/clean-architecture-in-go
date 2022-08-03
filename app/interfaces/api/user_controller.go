package controller

import (
	"fmt"
	"net/http"

	"github.com/Kantaro0829/clean-architecture-in-go/domain"
	"github.com/Kantaro0829/clean-architecture-in-go/interfaces/database"
	"github.com/Kantaro0829/clean-architecture-in-go/interfaces/token"
	"github.com/Kantaro0829/clean-architecture-in-go/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(
	sqlHandler database.SqlHandler,
	//以下追加
	tokenHandler token.TokenHandler,
) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
			//以下追加
			Tokenizer: &token.TokenizerImpl{
				TokenHandler: tokenHandler,
			},
		},
	}
}

//ルーティングのハンドラ
func (controller *UserController) Create(c *gin.Context) {
	var userJson domain.User
	//上で宣言した構造体にJsonをバインド。エラーならエラー処理を返す
	if err := c.ShouldBindJSON(&userJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mail, name, password := userJson.Mail, userJson.Name, userJson.Password

	user := domain.User{}
	user.Mail = mail
	user.Name = name
	user.Password = password
	fmt.Printf("json中身%v", user)
	//u := domain.User{}
	//c.Bind(&u)
	//controller.Interactor.Add(u)
	err := controller.Interactor.Add(user)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "データ登録完了"})

	return
}

func (controller *UserController) GetUser(c *gin.Context) {
	res, err := controller.Interactor.GetInfo()
	if err != nil {
		//エラーハンドリング
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (controller *UserController) Delete(id string) {
	controller.Interactor.Delete(id)
}

func (controller *UserController) Update(c *gin.Context) {
	user := domain.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.Interactor.Update(user, user.Name)
}

func (controller *UserController) UpdateByMail(user domain.User) (string, error, bool) {
	message, err, isValidated := controller.Interactor.UpdateUser(user)
	if err != nil {
		return message, err, isValidated
	}
	return message, nil, isValidated
}

func (controller *UserController) DeleteByMail(user domain.User) (string, error, bool) {
	message, err, isValidated := controller.Interactor.DeleteByMail(user)
	if err != nil {
		return message, err, isValidated
	}
	return message, nil, isValidated
}

func (controller *UserController) Login(mail string, password string) (domain.Token, bool, error) {
	//以下Token追加
	token, result, err := controller.Interactor.Login(mail, password)

	if err != nil {
		//以下エラー時Tokenの代わりに空文字を返す処理追加
		return "", false, err
	}
	fmt.Println(token)
	//戻り値にtoken追加
	return token, result, nil
}
