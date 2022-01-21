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
	rbacRouter.GET("/permission", permissionController.Get)
}
