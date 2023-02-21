package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-admin-template/pkg/util"
)

func RequestId(c *gin.Context) {
	if c.Request.Method == http.MethodOptions {
		c.Next()
		return
	}
	requestId := c.GetHeader(util.TrafficKey)
	if requestId == "" {
		requestId = c.GetHeader(strings.ToLower(util.TrafficKey))
	}
	if requestId == "" {
		requestId = uuid.New().String()
	}

	c.Header(util.TrafficKey, requestId)
	c.Set(util.TrafficKey, requestId)
	c.Next()
}
