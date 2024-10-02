package controllers

import (
	"starter/mod/internal/services"
	"starter/mod/package/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	// uid := c.Param("uid")
	response.SuccessResponse(c, 20001, []string{"user1", "user2"})
}
