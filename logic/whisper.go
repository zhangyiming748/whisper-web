package logic

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func Start(language, pattern, model string) {
	path := "/videos"
	if !isExist(path) {
		path = "videos"
		os.MkdirAll(path, os.ModePerm)
	}
	files := GetFiles(path)
	for _, file := range files {
		if strings.HasSuffix(file, pattern) {
			GetSubtitle(file, language, model)
		}
	}
}

/*
生成字幕后返回字幕的绝对路径
*/

func GetSubtitle(fp, language, level string) string {
	//ase := filepath.Base(fp)
	path := "/videos"
	if !isExist(path) {
		path = "videos"
		os.MkdirAll(path, os.ModePerm)
	}
	err := os.Setenv("PYTHONIOENCODING", "utf-8")
	if err == nil {
		log.Println("utf-8环境设置成功")
	}
	var cmd = exec.Command("whisper", fp, "--model", level, "--model_dir", path, "--output_format", "srt", "--prepend_punctuations", ",.?", "--language", language, "--output_dir", path, "--verbose", "True")

	startTime := time.Now()

	err = execCommand(cmd)
	if err != nil {
		log.Printf("当前字幕生成错误\t命令原文:%v\t错误原文:%v\n", cmd.String(), err.Error())
	}
	fp = strings.Replace(fp, filepath.Ext(fp), ".srt", 1)

	//replace.RemoveTrailingNewlines(fp)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	totalMinutes := duration.Seconds() / 60
	log.Printf("文件%v\n总共用时: %.2f 分钟\n", fp, totalMinutes)
	return fp
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

/*
执行命令过程中可以循环打印消息
*/
func execCommand(c *exec.Cmd) (e error) {
	log.Printf("开始执行命令:%v\n", c.String())
	stdout, err := c.StdoutPipe()
	c.Stderr = c.Stdout
	if err != nil {
		log.Printf("连接Stdout产生错误:%v\n", err)
		return err
	}
	if err = c.Start(); err != nil {
		log.Printf("启动cmd命令产生错误:%v\n", err)
		return err
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		t := string(tmp)
		t = strings.Replace(t, "\u0000", "", -1)
		fmt.Printf("\r%v", t)
		if err != nil {
			break
		}
	}
	if err = c.Wait(); err != nil {
		log.Printf("命令执行中产生错误:%v\n", err)
		return err
	}
	return nil
}
func GetFiles(currentDir string) (filePaths []string) {
	err := filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 检查是否是文件
		if !info.IsDir() {
			filePaths = append(filePaths, path) // 将文件的绝对路径添加到切片
		}
		return nil
	})

	if err != nil {
		fmt.Println("遍历目录失败:", err)
		return
	}

	// 打印所有文件的绝对路径
	for _, filePath := range filePaths {
		fmt.Println(filePath)
	}
	return filePaths
}
