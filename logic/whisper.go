package logic

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func DownloadVideos(list []string, proxy string) {
	for i := range list {
		DownloadVideo(list[i], proxy)
	}
}
func DownloadVideo(uri, proxy string) (fp string) {
	path := "/videos"
	if !isExist(path) {
		path = "videos"
		os.MkdirAll(path, os.ModePerm)
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	file, err := os.OpenFile("下载失败.log", os.O_RDONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		return "错误日志创建失败"
	}
	defer func() {
		file.Sync()
		file.Close()
	}()
	name_cmd := exec.Command("yt-dlp", "--proxy", proxy, "-f", "bestvideo[height<=?1080]+bestaudio/best[height<=?1080]/mp4", "--no-playlist", "--get-filename", uri)
	name := getVideoName(name_cmd)
	log.Printf("当前下载的文件标题:%s", name)
	download_cmd := exec.Command("yt-dlp", "--proxy", proxy, "-f", "bestvideo[ext=mp4][height<=?1080]+bestaudio[ext=m4a]/best[height<=?1080]/mp4", "--no-playlist", "--paths", path, uri)
	output, err := download_cmd.CombinedOutput()
	if err != nil {
		log.Printf("执行命令失败: %v, 错误: %v\n", download_cmd.String(), err)
		file.WriteString(fmt.Sprintf("执行命令失败: %v, 错误: %v\n", download_cmd.String(), err))
		return ""
	}
	log.Printf("当前下载成功的文件标题:%s\t返回%s\n", name, string(output))
	return name
}
func getVideoName(c *exec.Cmd) (name string) {
	output, _ := c.CombinedOutput()
	return string(output)
}
func isExist(dirPath string) bool {
	//dirPath := "/path/to/your/directory" // 替换为你想要检查的目录路径
	_, err := os.Stat(dirPath)
	if err == nil {
		fmt.Printf("目录 %s 存在\n", dirPath)
		return true
	} else if os.IsNotExist(err) {
		fmt.Printf("目录 %s 不存在\n", dirPath)
		return false
	} else {
		fmt.Printf("访问目录 %s 时出错: %v\n", dirPath, err)
		return false
	}
}
