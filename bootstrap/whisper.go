package bootstrap

import (
	"whisper/controller"

	"github.com/gin-gonic/gin"
)

func InitYtdlp(engine *gin.Engine) {
	routeGroup := engine.Group("/api")
	{
		c := new(controller.WhisperController)
		//routeGroup.GET("/v1/s1/gethello", c.GetHello)
		routeGroup.POST("/v1/whisper/download", c.DownloadAll)
	}
}
