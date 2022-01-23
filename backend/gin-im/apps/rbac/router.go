package rbac

import (
	"gin-im/apps/rbac/controllers"
	"github.com/gin-gonic/gin"
)

var (
	permissionController = controllers.PermissionController{}
)

func Router(r *gin.Engine) {
	rbacRouter := r.Group("/rbac")
	// 权限
	rbacRouter.GET("/permission/*id", permissionController.GET)
	rbacRouter.POST("/permission", permissionController.POST)
	rbacRouter.PUT("/permission/:id", permissionController.PUT)
	rbacRouter.DELETE("/permission/:id", permissionController.DELETE)
}
