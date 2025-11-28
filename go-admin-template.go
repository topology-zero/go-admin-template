package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"net/http"
	"time"

	"go-admin-template/config"
	"go-admin-template/middleware"
	"go-admin-template/model"
	"go-admin-template/pkg/logger"
	"go-admin-template/pkg/prometheus"
	"go-admin-template/pkg/redis"
	"go-admin-template/pkg/swagger"
	"go-admin-template/query"
	"go-admin-template/routes"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//go:generate gengin go-admin-template.api
//go:generate gen-swagger --local_api= go-admin-template.api

func main() {
	flag.Parse()

	configFile := fmt.Sprintf("etc/go-admin-template-%s.yaml", config.Env)

	e := gin.New()
	e.Use(
		middleware.RequestId,
		middleware.RequestLog,
		gin.Recovery(),
		middleware.CorsMiddleware,
		middleware.ApiHitRecord,
	)

	pprof.Register(e)
	config.Setup(configFile)
	logger.Setup()
	prometheus.Setup(e)
	redis.Setup()
	model.Setup()
	query.SetDefault(model.DB())
	routes.Setup(e)
	swagger.Setup(e)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.ServerConf.Host, config.ServerConf.Port),
		Handler: e,
	}
	go func() {
		logrus.Info("listen to ", server.Addr)
		server.ListenAndServe()
	}()
	wait := config.RegisterCloseFn(func() {
		defer logrus.Warning("closed api server")

		ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancelFunc()
		server.Shutdown(ctx)
	})
	wait()
}
