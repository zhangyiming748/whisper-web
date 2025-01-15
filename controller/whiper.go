package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

type WhisperController struct{}

// 结构体必须大写 否则找不到
type Whisper struct {
	Language string `json:"language"` // "English"
	Pattern  string `json:"pattern"`  // "mp4"
	Model    string `json:"model"`    // "large-v3"
}
type WhisperResponseBody struct {
	Language string `json:"language"` // "English"
	Pattern  string `json:"pattern"`  // "mp4"
	Model    string `json:"model"`    // "large-v3"
	Msg      string `json:"msg"`
}

/*
curl --location --request POST 'http://127.0.0.1:8193/api/v1/telegram/download' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \

	--data-raw '{
	    "urls": [
	        "string"
	    ],
	    "proxy": "string"
	}'
*/
func (y WhisperController) DownloadAll(ctx *gin.Context) {
	req := new(Whisper)
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//fmt.Printf("url = %s \nproxy = %s\n", req.URLs, req.Proxy)
	rep := WhisperResponseBody{
		Language: req.Language,
		Pattern:  req.Pattern,
		Model:    req.Model,
		Msg:      "已经开始转换",
	}
	log.Println("开始转换")
	//go logic.DownloadVideos(req.URLs, req.Proxy)
	ctx.JSON(200, rep)
}
