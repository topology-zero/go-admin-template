package model

import (
	_ "embed"
	"fmt"
	"time"

	"go-admin-template/config"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db       *gorm.DB
	Enforcer *casbin.Enforcer
)

func Setup() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MysqlConf.User,
		config.MysqlConf.Pwd,
		config.MysqlConf.Host,
		config.MysqlConf.Port,
		config.MysqlConf.Db,
	)
	var err error

	logMode := logger.Silent
	if config.Env == "local" || config.Env == "dev" {
		logMode = logger.Info
	}
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: NewLogger(logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logMode,
		}),
	})
	if err != nil {
		logrus.Panic(err)
	}

	casbinSetup()
}

func casbinSetup() {
	adapter, err := gormadapter.NewAdapterByDBUseTableName(db, "", "admin_casbin_rule")
	if err != nil {
		logrus.Panicf("%+v", errors.WithStack(err))
	}
	Enforcer, err = casbin.NewEnforcer("etc/rbac_model.conf", adapter)
	if err != nil {
		logrus.Panicf("%+v", errors.WithStack(err))
	}

	if err = Enforcer.LoadPolicy(); err != nil {
		logrus.Panicf("%+v", errors.WithStack(err))
	}
}

// DB 获取数据库实例
func DB() *gorm.DB {
	return db
}