package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	filename "github.com/keepeye/logrus-filename"
	"github.com/shoppehub/conf"
	"github.com/sirupsen/logrus"
)

var Port = 4000

func New() *gin.Engine {

	filenameHook := filename.NewHook()
	logrus.AddHook(filenameHook)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)

	logrus.Debug("init server")

	conf.Init("")

	// 使用默认中间件（logger和recovery）
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.GET("/", RenderTemplate)
	r.GET("/:module", RenderTemplate)
	r.GET("/:module/:page", RenderTemplate)
	// r.GET("/mod/:name", server.RenderTemplate)

	r.GET("/assets/*path", func(c *gin.Context) {
		ProxyHandler(c.Writer, c.Request)
	})

	port := conf.GetInt("port")
	if port != 0 {
		Port = int(port)
	}
	logrus.Info("start server on " + fmt.Sprint(port))

	return r
}
