package svc

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-admin-template/pkg/util"
)

type ServiceContext struct {
	GinContext *gin.Context
	Log        *logrus.Entry
}

func NewServiceContext(c *gin.Context) *ServiceContext {
	traceId, _ := c.Get(util.TrafficKey)
	return &ServiceContext{
		GinContext: c,
		Log:        logrus.WithField(util.TraceId, traceId.(string)),
	}
}
