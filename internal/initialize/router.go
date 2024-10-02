package initialize

import (
	"fmt"
	c "starter/mod/internal/controllers"
	"starter/mod/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthMiddleware())
	r.GET("users/:uid", c.NewUserController().GetUserById)

	return r
}

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("Before AA")
		c.Next()
		fmt.Println("After AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("Before BB")
		c.Next()
		fmt.Println("After BB")
	}
}
