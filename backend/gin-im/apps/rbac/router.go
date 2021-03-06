package rbac

import (
	"gin-im/apps/admin/middlewares"
	"gin-im/apps/rbac/controllers"
	"github.com/gin-gonic/gin"
)

var (
	permissionController     = controllers.PermissionController{}
	permissionTreeController = controllers.PermissionTreeController{}
	roleController           = controllers.RoleController{}
	authMiddleware           = middlewares.AuthMiddleware{}
)

func Router(r *gin.RouterGroup) {
	rbacRouter := r.Group("/rbac")
	rbacRouter.Use(authMiddleware.CheckAuth)
	// 权限
	rbacRouter.GET("/permission/*id", permissionController.GET)
	rbacRouter.POST("/permission", permissionController.POST)
	rbacRouter.PUT("/permission/:id", permissionController.PUT)
	rbacRouter.DELETE("/permission/:id", permissionController.DELETE)

	// 权限树
	rbacRouter.GET("/permission_tree/*id", permissionTreeController.GET)

	// 角色
	rbacRouter.GET("/role/*id", roleController.GET)
}
