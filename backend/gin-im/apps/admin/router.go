package admin

import (
	"gin-im/apps/admin/controllers"
	"gin-im/apps/admin/middlewares"
	"github.com/gin-gonic/gin"
)

var (
	userController = controllers.UserController{}
	authMiddleware = middlewares.AuthMiddleware{}
)

func Router(r *gin.Engine) {
	userRouter := r.Group("/user")
	userRouter.POST("/login", userController.Login)
	userRouter.POST("/register", userController.Register)
	userRouter.Use(authMiddleware.CheckAuth)
	{
		userRouter.GET("/profile", userController.Profile)
	}

}
