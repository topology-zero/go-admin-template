package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

var (
	cfgCnf     = Cfg{}
	ServerConf = Server{}
	JwtConf    = Jwt{}
	LogConf    = Log{}
	MysqlConf  = Mysql{}
	RedisConf  = Redis{}
)

func Setup(fileName string) {
	viper.SetConfigFile(fileName)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Panic("读取配置文件错误", err)
	}

	if err := viper.Unmarshal(&cfgCnf); err != nil {
		log.Panic("配置文件格式错误", err)
	}

	ServerConf = cfgCnf.Server
	JwtConf = cfgCnf.Jwt
	LogConf = cfgCnf.Log
	MysqlConf = cfgCnf.Mysql
	RedisConf = cfgCnf.Redis

}
