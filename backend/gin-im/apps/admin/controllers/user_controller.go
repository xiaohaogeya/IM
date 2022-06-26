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
	db.DB.Where(user).First(user)

	if user.ID == 0 || !user.ValidatePassword(form.Password) {
		c.Error(ctx, "400002")
		return
	}
	user.LoginAt = utils.NewTime().GetNow()
	db.DB.Save(user)

	token, _ := auth.GenerateToken(user.ID, user.UserName)
	c.Success(ctx, gin.H{
		"token":        token,
		"user_id":      user.ID,
		"user_name":    user.UserName,
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
	db.DB.Where(user).First(user)

	if user.ID != 0 {
		c.Error(ctx, "400004")
		return
	}

	password := form.Password
	password = user.MakePassword(password)
	user.Password = password

	// 事务
	tx := db.DB.Begin()
	if err := tx.Create(user).Error; err != nil {
		// 返回任何错误都会回滚事务
		tx.Rollback()
		c.Error(ctx, "400005")
		return
	}

	userProfile := &models.UserProfile{
		UserId:   user.ID,
		NickName: form.NickName,
		Mobile:   form.Mobile,
		Email:    form.Email,
	}
	if err := tx.Create(userProfile).Error; err != nil {
		// 返回任何错误都会回滚事务
		tx.Rollback()
		c.Error(ctx, "400005")
		return
	}
	// 提交事务
	tx.Commit()
	c.Success(ctx, gin.H{"user_id": user.ID})
}

// Profile 用户详情
func (c *UserController) Profile(ctx *gin.Context) {
	userId := c.UserId(ctx)

	userProfile := &models.UserProfile{UserId: userId}
	db.DB.Joins("User").Where(userProfile).First(userProfile)

	c.Success(ctx, userProfile)
}
