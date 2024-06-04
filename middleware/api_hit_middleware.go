package middleware

import (
	"time"

	"go-admin-template/pkg/prometheus"

	"github.com/gin-gonic/gin"
)

func ApiHitRecord(c *gin.Context) {
	// 开始时间
	startTime := time.Now()

	c.Next()

	if c.Writer.Status() != 200 {
		return
	}

	if c.FullPath() == "/metrics" {
		return
	}

	// 只记录合法操作
	prometheus.ApiHit.With(map[string]string{
		"api":    c.FullPath(),
		"method": c.Request.Method,
	}).Observe(float64(time.Since(startTime).Milliseconds()))
}
