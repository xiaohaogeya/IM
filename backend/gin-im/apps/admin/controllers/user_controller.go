package controllers

import (
	"gin-im/apps/admin/auth"
	"gin-im/apps/admin/models"
	"gin-im/db"
	"gin-im/utils"
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

	if user.ID == 0 || !user.ValidatePassword(form.Password) {
		c.Error(ctx, "400002")
		return
	}
	user.LoginAt = utils.NewTime().GetNow()
	db.DB.Save(user)

	userProfile := &models.UserProfile{UserId: user.ID}
	db.DB.Find(userProfile)

	token, _ := auth.GenerateToken(user.ID, user.UserName)
	c.Success(ctx, gin.H{
		"token":        token,
		"user_id":      user.ID,
		"user_name":    user.UserName,
		"user_profile": userProfile,
	})
}

// Register 注册
func (c *UserController) Register(ctx *gin.Context) {
	form := struct {
		UserName string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
		NickName string `form:"nick_name"`
		Mobile   string `form:"mobile"`
		Email    string `form:"email"`
	}{}

	if err := ctx.ShouldBind(&form); err != nil {
		c.Error(ctx, "400001")
		return
	}
	validate := utils.Validator{}
	if form.Mobile != "" && !validate.ValidateMobile(form.Mobile) {
		c.Error(ctx, "400006")
		return
	}

	if form.Email != "" && !validate.ValidateEmail(form.Email) {
		c.Error(ctx, "400007")
		return
	}

	user := &models.User{
		UserName: form.UserName,
	}
	db.DB.Find(user)

	if user.ID != 0 {
		c.Error(ctx, "400004")
		return
	}

	password := form.Password
	password = user.MakePassword(password)
	user.Password = password
	db.DB.Create(user)

	if user.ID == 0 {
		c.Error(ctx, "400005")
		return
	}

	userProfile := models.UserProfile{
		UserId:   user.ID,
		NickName: form.NickName,
		Mobile:   form.Mobile,
		Email:    form.Email,
	}
	db.DB.Create(&userProfile)
	c.Success(ctx, gin.H{"user_id": user.ID})
}

// Profile 用户详情
func (c *UserController) Profile(ctx *gin.Context) {
	userId := c.UserId(ctx)

	userProfile := &models.UserProfile{UserId: userId}
	db.DB.Joins("User").First(userProfile)

	c.Success(ctx, userProfile)
}
