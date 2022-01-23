package rbac

import (
	"gin-im/apps/rbac/controllers"
	"github.com/gin-gonic/gin"
)

var (
	permissionController = controllers.PermissionController{}
	permissionTreeController = controllers.PermissionTreeController{}
)

func Router(r *gin.Engine) {
	rbacRouter := r.Group("/rbac")
	// 权限
	rbacRouter.GET("/permission/*id", permissionController.GET)
	rbacRouter.POST("/permission", permissionController.POST)
	rbacRouter.PUT("/permission/:id", permissionController.PUT)
	rbacRouter.DELETE("/permission/:id", permissionController.DELETE)

	// 权限树
	rbacRouter.GET("/permission_tree/*id", permissionTreeController.GET)
}
