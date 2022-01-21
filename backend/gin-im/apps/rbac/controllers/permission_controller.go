package controllers

import (
	"gin-im/apps/rbac/models"
	"gin-im/db"
	"github.com/gin-gonic/gin"
)

type PermissionController struct {
	BaseController
}

func (c *PermissionController) Get(ctx *gin.Context) {
	var permissionList []models.Permission
	db.DB.Find(&permissionList)
	c.Success(ctx, gin.H{"permissions": permissionList})
}
