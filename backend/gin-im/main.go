package main

import (
	ws2 "gin-im/apps/ws"
	"gin-im/routers"
)

func main() {

	// 开启ws协程
	go ws2.Manager.Start()

	routers.Run()
}
