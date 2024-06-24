package model

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	"go-admin-template/config"
	internalLogger "go-admin-template/pkg/logger"
	"go-admin-template/pkg/util"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

type gormLogger struct {
	config logger.Config
	log    *logrus.Logger
}

func NewLogger(cfg logger.Config) logger.Interface {
	log := logrus.New()
	level, _ := logrus.ParseLevel(config.LogConf.Level)
	log.SetLevel(level)

	//设置日志格式
	log.SetFormatter(&logrus.TextFormatter{
		ForceQuote:      false,
		DisableQuote:    true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	})

	// 取消线程安全
	log.SetNoLock()

	log.AddHook(&internalLogger.RotateHook{})

	return &gormLogger{
		config: cfg,
		log:    log,
	}
}

func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.config.LogLevel = level
	return &newLogger
}

func (l *gormLogger) Info(ctx context.Context, msg string, data ...any) {
	if l.config.LogLevel >= logger.Info {
		return
	}
	l.getLogger(ctx).Info(append([]any{msg}, data...)...)
}

func (l *gormLogger) Warn(ctx context.Context, msg string, data ...any) {
	if l.config.LogLevel >= logger.Warn {
		return
	}
	l.getLogger(ctx).Warn(append([]any{msg}, data...)...)
}

func (l *gormLogger) Error(ctx context.Context, msg string, data ...any) {
	if l.config.LogLevel >= logger.Error {
		return
	}
	l.getLogger(ctx).Error(append([]any{msg}, data...)...)
}

func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.config.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()
	field := logrus.Fields{
		"useTime": fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6),
		"sql":     sql,
	}
	if rows != -1 {
		field["rows"] = strconv.FormatInt(rows, 10)
	}

	switch {
	case err != nil && l.config.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound) || !l.config.IgnoreRecordNotFoundError):
		l.getLogger(ctx).WithFields(field).Error(err.Error())
	case elapsed > l.config.SlowThreshold && l.config.SlowThreshold != 0 && l.config.LogLevel >= logger.Warn:
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.config.SlowThreshold)
		l.getLogger(ctx).WithFields(field).Warn(slowLog)
	case l.config.LogLevel == logger.Info:
		l.getLogger(ctx).WithFields(field).Trace()
	}
}

func (l *gormLogger) getLogger(ctx context.Context) *logrus.Entry {
	requestId := ctx.Value(util.TrafficKey)
	field := logrus.Fields{
		"file": getStack(),
	}
	if requestId != nil {
		field[util.TraceId] = requestId.(string)
	}
	return l.log.WithFields(field)
}

// 获取pkg目录
func getStack() string {
	pcs := make([]uintptr, 25)
	n := runtime.Callers(4, pcs[:])
	stack := pcs[0:n]
	frames := runtime.CallersFrames(stack)
	for {
		frame, more := frames.Next()
		if !more {
			break
		}
		if strings.HasPrefix(frame.Function, util.GetPkgName()) && !strings.HasSuffix(frame.File, ".gen.go") {
			return fmt.Sprintf("%s:%d", frame.File, frame.Line)
		}
	}

	return ""
}
