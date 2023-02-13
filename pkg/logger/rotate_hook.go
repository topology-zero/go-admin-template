package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"admin_template/config"
	"github.com/sirupsen/logrus"
)

type RotateHook struct{}

var openFile *os.File

func (h *RotateHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *RotateHook) Fire(entry *logrus.Entry) (err error) {
	fileName := time.Now().Format("2006-01-02")
	fullPath := fmt.Sprintf("%s/%s.log", config.LogConf.Dir, fileName)

	// 无需多次获取文件句柄
	if openFile != nil && openFile.Name() == fullPath {
		return
	}

	if err = os.MkdirAll(config.LogConf.Dir, os.ModePerm); err != nil {
		log.Panic("创建文件夹错误", err)
		return
	}

	openFile, err = os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Panic("写入日志文件错误", err)
		return
	}

	//设置输出
	logrus.SetOutput(io.MultiWriter(openFile, os.Stdout))
	return
}
