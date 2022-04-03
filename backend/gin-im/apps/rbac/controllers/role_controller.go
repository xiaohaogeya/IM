package controllers

import (
	"gin-im/apps/rbac/models"
	"gin-im/db"
	"github.com/gin-gonic/gin"
	"strings"
)

type RoleController struct {
	BaseController
}

func (c *RoleController) GET(ctx *gin.Context) {
	id := ctx.Param("id")
	var roleList []models.Role
	if id != "/" {
		id = strings.TrimLeft(id, "/")
		db.DB.Preload("Permissions").Where("id = ?", id).Find(&roleList)
	} else {
		db.DB.Find(&roleList)
	}
	c.Success(ctx, roleList)
}
