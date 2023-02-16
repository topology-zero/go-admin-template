package logger

import (
	"fmt"
	"log"
	"runtime"

	"admin_template/config"

	"github.com/sirupsen/logrus"
)

func Setup() {
	level, err := logrus.ParseLevel(config.LogConf.Level)
	if err != nil {
		log.Panic("日志level格式设置错误", err)
	}
	logrus.SetLevel(level)

	//设置日志格式
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceQuote:      false,
		DisableQuote:    true,
		TimestampFormat: "2006-01-02 15:04:05.000",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return "", fmt.Sprintf("%s:%d", frame.File, frame.Line)
		},
	})

	// 取消线程安全
	logrus.StandardLogger().SetNoLock()

	// 打印堆栈信息
	logrus.SetReportCaller(true)

	logrus.AddHook(&RotateHook{})
}
