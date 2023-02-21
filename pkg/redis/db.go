package redis

import (
	"context"
	"fmt"

	redigo "github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"go-admin-template/config"
	"go-admin-template/pkg/logger"
)

var Client *redigo.Client

func Setup() {
	Client = redigo.NewClient(&redigo.Options{
		Addr:     fmt.Sprintf("%s:%d", config.RedisConf.Host, config.RedisConf.Port),
		Username: config.RedisConf.User,
		Password: config.RedisConf.Pwd,
		DB:       config.RedisConf.Db,
		PoolSize: 10,
	})
	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		logger.Fatalf("%+v", errors.WithStack(err))
	}
}
