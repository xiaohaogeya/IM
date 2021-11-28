package controllers

import (
	"gin-im/auth"
	"gin-im/db"
	"gin-im/models"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

// Login 登录
func (c *UserController) Login(ctx *gin.Context) {
	form := struct {
		UserName string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}{}

	if err := ctx.ShouldBind(&form); err != nil {
		c.Error(ctx, "400001")
		return
	}

	user := &models.User{
		UserName: form.UserName,
	}
	db.DB.Find(user)

	if user.ID == 0 || !user.ValidatePassword(form.Password){
		c.Error(ctx, "400002")
		return
	}

	token, _ := auth.GenerateToken(user)
	data := make(map[string]interface{})
	data["token"] = token
	data["user"] = user
	c.Success(ctx, data)
}
