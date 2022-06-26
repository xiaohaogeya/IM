package ws

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup)  {
	wsRouter := r.Group("/ws")
	wsRouter.GET("/", WsHandler)
}