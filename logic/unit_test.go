package logic

import (
	"testing"
	"whisper/util"
)

func init() {
	util.SetLog("telegram.log")
}
func TestDownloadVideo(t *testing.T) {
	DownloadVideo("https://youtu.be/1JMFo_vPaKE?list=RD1JMFo_vPaKE", "192.168.1.35:8889")
}
