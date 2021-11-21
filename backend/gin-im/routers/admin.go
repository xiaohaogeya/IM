package routers

import "gin-im/controllers"

func adminRouter() {
	r := router()
	adminRouter := r.Group("/admin")
	adminRouter.GET("/", controllers.Get)
}
