package redis

import (
	"context"
	"fmt"

	"go-admin-template/config"

	redigo "github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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
		logrus.Fatalf("%+v", errors.WithStack(err))
	}
}
