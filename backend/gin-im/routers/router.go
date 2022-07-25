package routers

import (
	"gin-im/apps/admin"
	"gin-im/apps/rbac"
	"gin-im/apps/ws"
	"gin-im/conf"
	"gin-im/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var engine *gin.Engine

func init() {
	gin.SetMode(conf.AppConfig.AppModel)
	engine = gin.Default()
	// 跨域
	defaultConfig := cors.DefaultConfig()
	defaultConfig.AllowAllOrigins = true
	defaultConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	defaultConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	defaultConfig.ExposeHeaders = []string{"Content-Length"}
	engine.Use(cors.New(defaultConfig))
}

var (
	g errgroup.Group
)

func router() {
	// 没有路由即 404返回
	engine.NoRoute(func(g *gin.Context) {
		errMsg := utils.ErrMsg{}
		code := "400000"
		msg := errMsg.String(code)
		g.AbortWithStatusJSON(http.StatusNotFound, gin.H{"msg": msg, "code": code})
	})
	api := engine.Group("/api")

	ws.Router(api)
	admin.Router(api)
	rbac.Router(api)
}

func Run() {

	router()

	s := &http.Server{
		Addr:           ":" + conf.AppConfig.AppPort,
		Handler:        engine,
		ReadTimeout:    time.Duration(conf.AppConfig.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(conf.AppConfig.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	g.Go(func() error {
		return s.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
