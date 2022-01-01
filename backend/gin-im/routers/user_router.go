package routers

import (
	controllers2 "gin-im/app/controllers"
	middlewares2 "gin-im/app/middlewares"
)

var (
	userController = controllers2.UserController{}
	authMiddleware = middlewares2.AuthMiddleware{}
)

func userRouter() {
	r := router()
	userRouter := r.Group("/user")
	userRouter.POST("/login", userController.Login)
	userRouter.POST("/register", userController.Register)
	userRouter.Use(authMiddleware.CheckAuth)
	{
		userRouter.GET("/profile", userController.Profile)
	}

}
