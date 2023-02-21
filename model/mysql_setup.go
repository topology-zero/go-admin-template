package model

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go-admin-template/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glog "gorm.io/gorm/logger"
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

	logMode := glog.Silent
	if config.Env == "local" || config.Env == "dev" {
		logMode = glog.Info
	}
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: glog.New(logrus.StandardLogger(), glog.Config{
			SlowThreshold:             200 * time.Millisecond,
			Colorful:                  false,
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

// WarpError 将错误打印，且不包含敏感信息返回
func WarpError(err error) error {
	if err == nil {
		return nil
	}
	logrus.Errorf("数据库异常：%+v", errors.WithStack(err))
	return errors.New("系统错误")
}
