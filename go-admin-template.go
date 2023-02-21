package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"go-admin-template/config"
	"go-admin-template/middleware"
	"go-admin-template/model"
	"go-admin-template/pkg/logger"
	"go-admin-template/pkg/swagger"
	"go-admin-template/query"
	"go-admin-template/routes"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//go:generate goctl api plugin -p gengin -api go-admin-template.api -dir .
//go:generate goctl api plugin -p "goctl-swagger swagger -filename asset/swagger/swagger.json" -api go-admin-template.api -dir .
func main() {
	flag.Parse()

	configFile := fmt.Sprintf("etc/go-admin-template-%s.yaml", config.Env)

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
