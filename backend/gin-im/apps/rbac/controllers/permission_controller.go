package controllers

import (
	"gin-im/apps/rbac/models"
	"gin-im/db"
	"github.com/gin-gonic/gin"
	"strings"
)

type PermissionController struct {
	BaseController
}

type PermissionTreeController struct {
	BaseController
}

func (c *PermissionController) GET(ctx *gin.Context) {
	id := ctx.Param("id")
	var permissionList []models.Permission
	if id != "/" {
		id = strings.TrimLeft(id, "/")
		db.DB.Where("id = ?", id).Find(&permissionList)
	} else {
		db.DB.Find(&permissionList)
	}
	c.Success(ctx, permissionList)
}

func (c *PermissionController) POST(ctx *gin.Context) {
	permission := &models.Permission{}
	if err := ctx.ShouldBind(permission); err != nil {
		c.Error(ctx, "400001")
		return
	}

	tx := db.DB.Create(permission)
	if tx.Error != nil {
		c.Error(ctx, "400008")
		return
	}
	c.Success(ctx, permission)
}

func (c *PermissionController) PUT(ctx *gin.Context) {
	id := ctx.Param("id")
	data := make(map[string]interface{})
	err := ctx.BindJSON(&data)
	if err != nil {
		c.Error(ctx, "400009")
		return
	}

	permission := &models.Permission{}
	if tx := db.DB.Model(&models.Permission{}).Where("id = ?", id).Updates(data).First(permission); tx.Error != nil {
		c.Error(ctx, "400008")
		return
	}
	c.Success(ctx, permission)
}

func (c *PermissionController) DELETE(ctx *gin.Context) {
	id := ctx.Param("id")
	if tx := db.DB.Model(&models.Permission{}).Where("id = ?", id).Delete(models.Permission{}); tx.Error != nil {
		c.Error(ctx, "400008")
		return
	}
	c.Success(ctx, "")
}

// GET 默认过滤掉parent_id为空
func (c *PermissionTreeController) GET(ctx *gin.Context) {
	id := ctx.Param("id")
	var permissionList []models.Permission
	if id != "/" {
		id = strings.TrimLeft(id, "/")
		db.DB.Preload("Children").Where("parent_id is null").Where("id = ?", id).Find(&permissionList)
	} else {
		db.DB.Preload("Children").Where("parent_id is null").Find(&permissionList)
	}
	c.Success(ctx, permissionList)
}
