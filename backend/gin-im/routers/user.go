package routers

import "gin-im/controllers"

var userController = controllers.UserController{}

func userRouter() {
	r := router()
	userRouter := r.Group("/user")
	userRouter.POST("/login", userController.Login)
}

