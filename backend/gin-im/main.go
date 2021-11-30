package main

import (
	"gin-im/routers"
	"gin-im/ws"
)

func main() {

	// 开启ws协程
	go ws.Manager.Start()

	routers.Run()
}
