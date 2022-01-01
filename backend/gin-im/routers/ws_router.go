package routers

import "gin-im/ws"

func wsRouter()  {
	r := router()
	wsRouter := r.Group("/ws")
	wsRouter.GET("/", ws.WsHandler)
}