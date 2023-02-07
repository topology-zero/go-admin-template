package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"admin_template/config"
	"admin_template/middleware"
	"admin_template/model"
	"admin_template/pkg/logger"
	"admin_template/pkg/swagger"
	"admin_template/query"
	"admin_template/routes"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//go:generate goctl api plugin -p gengin -api admin.api -dir .
//go:generate goctl api plugin -p "goctl-swagger swagger -filename asset/swagger/swagger.json" -api admin.api -dir .
func main() {
	flag.Parse()

	configFile := fmt.Sprintf("etc/admin-%s.yaml", config.Env)

	e := gin.New()
	e.Use(
		middleware.RequestId,
		middleware.RequestLog,
		gin.Recovery(),
		middleware.CorsMiddleware,
	)

	config.Setup(configFile)
	logger.Setup()
	//redis.Setup()
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
