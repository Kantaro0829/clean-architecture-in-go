package infrastructure

import (
	"net/http"

	controllers "github.com/Kantaro0829/clean-architecture-in-go/interfaces/api"
	"github.com/gin-gonic/gin"
)

func Init() {
	// Echo instance
	router := gin.Default()
	userController := controllers.NewUserController(NewSqlHandler())

	router.GET("/users", func(c *gin.Context) {
		users := userController.GetUser()
		//c.Bind(&users)
		c.JSON(http.StatusOK, gin.H{"a": users})
		return
	})

	router.POST("/users", func(c *gin.Context) {
		userController.Create(c)
		c.JSON(http.StatusOK, gin.H{"message": "data was inserted"})
		return
	})

	router.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		userController.Delete(id)
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
		return
	})

	router.Run(":3000")
}
