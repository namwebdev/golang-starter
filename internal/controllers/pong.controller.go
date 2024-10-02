package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Bong(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong " + name,
		"uid":     uid,
	})
}
