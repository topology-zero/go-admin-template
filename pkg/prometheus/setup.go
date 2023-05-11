package prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Setup(e *gin.Engine) {
	prometheus.MustRegister(ApiHit)
	e.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
