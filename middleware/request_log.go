package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"admin_template/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var writeList = []string{"/login"}

func RequestLog(c *gin.Context) {
	if c.Request.Method == http.MethodOptions {
		c.Next()
		return
	}

	// 必须放在 next 之前,否则无法读取
	// TODO 需要处理 form-data 请求
	// TODO 需要处理 file 传参
	body, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	requestBody := string(body)

	// 开始时间
	startTime := time.Now()

	c.Next()

	for _, url := range writeList {
		if c.Request.RequestURI == url {
			return
		}
	}

	requestId, _ := c.Get(util.TrafficKey)

	data := logrus.Fields{
		"useTime":   time.Since(startTime),
		"uri":       "[" + c.Request.Method + "]" + c.Request.RequestURI,
		"status":    c.Writer.Status(),
		"requestId": requestId,
	}

	if len(requestBody) > 0 {
		data["request"] = requestBody
	}

	logrus.WithFields(data).Info()
}
