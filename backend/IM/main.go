package main

import (
	_ "IM/routers"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// 日志配置
func logConfig() {
	_ = logs.SetLogger(logs.AdapterMultiFile,
		`{"filename":"logs/IM.log","separate":["error", "info"],"maxdays":30,"color":true}`,
	)
	logs.EnableFuncCallDepth(true)

	f := &logs.PatternLogFormatter{
		Pattern:    "%F:%n|%w%t>> %m",
		WhenFormat: "2006-01-02",
	}
	logs.RegisterFormatter("pattern", f)
	// 全局配置
	_ = logs.SetGlobalFormatter("pattern")

	// 为了提升性能, 可以设置异步输出:
	logs.Async()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	logConfig()
	beego.Run()
}
