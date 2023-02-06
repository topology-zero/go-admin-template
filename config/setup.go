package config

import (
	"image/color"
	"log"

	"github.com/mojocn/base64Captcha"
	"github.com/spf13/viper"
)

var (
	ServerConf = Server{}
	JwtConf    = Jwt{}
	LogConf    = Log{}
	MysqlConf  = Mysql{}
	RedisConf  = Redis{}
	Captcha    *base64Captcha.Captcha
)

func init() {
	Captcha = base64Captcha.NewCaptcha(base64Captcha.NewDriverString(
		50,
		150,
		100,
		3,
		4,
		"234567890abcdefghjkmnpqrstuvwxyz",
		&color.RGBA{R: 240, G: 240, B: 246, A: 246},
		nil,
		[]string{"wqy-microhei.ttc"},
	), base64Captcha.DefaultMemStore)
}

func Setup(fileName string) {
	viper.SetConfigFile(fileName)

	if err := viper.ReadInConfig(); err != nil {
		log.Panic("读取配置文件错误", err)
	}

	if err := viper.UnmarshalKey("Server", &ServerConf); err != nil {
		log.Panic("APP配置文件格式错误", err)
	}

	if err := viper.UnmarshalKey("Jwt", &JwtConf); err != nil {
		log.Panic("Jwt配置文件格式错误", err)
	}

	if err := viper.UnmarshalKey("Mysql", &MysqlConf); err != nil {
		log.Panic("Mysql配置文件格式错误", err)
	}

	if err := viper.UnmarshalKey("Redis", &RedisConf); err != nil {
		log.Panic("Redis配置文件格式错误", err)
	}

	if err := viper.UnmarshalKey("Log", &LogConf); err != nil {
		log.Panic("Log配置文件格式错误", err)
	}
}
