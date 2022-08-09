package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/Kantaro0829/clean-architecture-in-go/domain"
	controllers "github.com/Kantaro0829/clean-architecture-in-go/interfaces/api"
	"github.com/gin-gonic/gin"
)

func Init() {
	// Echo instance
	router := gin.Default()
	userController := controllers.NewUserController(
		NewSqlHandler(),
		NewTokenHandler(), //jwt用のDI
	)

	router.GET("/users", func(c *gin.Context) {

		userController.GetUser(c)

		return
	})

	router.POST("/users", func(c *gin.Context) {
		userController.Create(c) //それぞれのルーティングごと関数を呼び出す
		// c.JSON(http.StatusOK, gin.H{"message": "data was inserted"})
		return
	})

	router.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		userController.Delete(id) //それぞれのルーティングごと関数を呼び出す
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
		return
	})

	router.POST("/users/update", func(c *gin.Context) {
		var userJson domain.User
		//上で宣言した構造体にJsonをバインド。エラーならエラー処理を返す
		if err := c.ShouldBindJSON(&userJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		message, err, isValidated := userController.UpdateByMail(userJson)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": message})
			return
		}
		if isValidated != true {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": message})
			return
		}
		//userController.Update(c)
		c.JSON(http.StatusOK, gin.H{"message": message})
		return
	})

	router.POST("users/login", func(c *gin.Context) {
		var userJson domain.User
		//上で宣言した構造体にJsonをバインド。エラーならエラー処理を返す
		if err := c.ShouldBindJSON(&userJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		mail, password := userJson.Mail, userJson.Password
		//以下Tokenを追加で受け取る
		token, result, err := userController.Login(mail, password)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "dbサーバーのエラー"})
			return
		}

		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "JWTtokenの発行失敗"})
			return
		}

		if result {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "ログイン完了", "token": token})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "パスワードかメールアドレスが違います"})
		return
	})

	router.POST("users/delete", func(c *gin.Context) {
		var userJson domain.User
		if err := c.ShouldBindJSON(&userJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		message, err, isValidated := userController.DeleteByMail(userJson)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": message})
			return
		}
		if isValidated != true {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": message})
			return
		}
		//userController.Update(c)
		c.JSON(http.StatusOK, gin.H{"message": message})
		return
	})

	// inUserRouter := router.Group("/users",
	// 	func(c *gin.Context) { userController.Authenticate(c) },
	// )

	router.GET("users/authenticate", func(c *gin.Context) {
		// token := domain.Token
		//token = c.Request.Header["Authorization"][0]
		result := userController.Authenticate(c)
		//result := userController.Authenticate(token)
		if result != nil {
			fmt.Printf(":エラー内容：%v", result)
			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"status":  http.StatusUnauthorized,
					"message": "JWT認証失敗",
				})
			return
		}
		c.JSON(
			http.StatusAccepted,
			gin.H{
				"statsu":  http.StatusAccepted,
				"message": "JWT認証成功",
			})
		return
	})

	router.Run(":3000")
}
