package controllers

import (
	"gin-im/models"
	"gin-im/utils"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	user := &models.User{
		Id: 1,
	}
	token, _ := utils.GenerateToken(user)
	c.JSON(200, gin.H{
		"message": "pong",
		"data": map[string]interface{}{
			"id":    1,
			"name":  "小浩",
			"token": token,
		},
	})
}
