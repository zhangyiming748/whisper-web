package main

import (
	"log"
	"net/http"
	"time"
	"whisper/bootstrap"
	"whisper/util"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func testResponse(c *gin.Context) {
	c.JSON(http.StatusGatewayTimeout, gin.H{
		"code": http.StatusGatewayTimeout,
		"msg":  "timeout",
	})
}

func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(3000*time.Millisecond),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(testResponse),
	)
}
func init() {
	util.SetLog("gin.log")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
func main() {
	// gin服务
	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	engine.Use(timeoutMiddleware())
	bootstrap.InitHello(engine)
	bootstrap.InitYtdlp(engine)
	// 启动http服务
	engine.Run(":9003")
}
