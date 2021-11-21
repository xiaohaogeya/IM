package main

import (
	"gin-im/conf"
	"gin-im/routers"
)

func main() {
	// 初始化文件配置
	err := conf.InitConfig()
	if err != nil {
		panic("初始化配置失败")
	}

	routers.Run()
}
