package infrastructure

import (
	"fmt"
	"net/http"

	controllers "github.com/Kantaro0829/clean-architecture-in-go/interfaces/api"
	"github.com/gin-gonic/gin"
)

func Init() {
	// Echo instance
	router := gin.Default()
	userController := controllers.NewUserController(NewSqlHandler())

	router.GET("/users", func(c *gin.Context) {
		//users, err := userController.GetUser() //それぞれのルーティングごと関数を呼び出す
		userController.GetUser(c)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"message": err.})
		// }
		//c.Bind(&users)
		//.JSON(http.StatusOK, gin.H{"a": users})
		return
	})

	router.POST("/users", func(c *gin.Context) {
		userController.Create(c) //それぞれのルーティングごと関数を呼び出す
		c.JSON(http.StatusOK, gin.H{"message": "data was inserted"})
		return
	})

	router.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		userController.Delete(id) //それぞれのルーティングごと関数を呼び出す
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
		return
	})

	router.POST("/users/update/:id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println(id)
		fmt.Println("update関数呼び出し")
		userController.Update(c)
		c.JSON(http.StatusOK, gin.H{"message": "data was updated"})
	})

	router.Run(":3000")
}
