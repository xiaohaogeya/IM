package routers

import (
	"gin-im/conf"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var engine *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	engine = gin.Default()
	// 跨域
	defaultConfig := cors.DefaultConfig()
	defaultConfig.AllowAllOrigins = true
	defaultConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	defaultConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	defaultConfig.ExposeHeaders = []string{"Content-Length"}
	engine.Use(cors.New(defaultConfig))
}

// router router
func router() *gin.Engine {
	return engine
}

var (
	g errgroup.Group
)

func Run() {

	wsRouter()
	userRouter()

	s := &http.Server{
		Addr:           ":" + conf.AppConfig.AppPort,
		Handler:        router(),
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
