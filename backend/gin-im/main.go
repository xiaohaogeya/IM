package main

import (
	"gin-im/conf"
	"gin-im/routers"
	"gin-im/ws"
)

func init() {
	// 初始化文件配置
	err := conf.InitConfig()
	if err != nil {
		panic("初始化配置失败")
	}
}

func main() {

	// 开启ws协程
	go ws.Manager.Start()

	routers.Run()
}
