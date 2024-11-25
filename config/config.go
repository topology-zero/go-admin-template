package config

type Cfg struct {
	Server Server
	Jwt    Jwt
	Log    Log
	Mysql  Mysql
	Redis  Redis
}

type Server struct {
	Name string
	Host string
	Port int
}

type Jwt struct {
	Secret string
	Expire int
}

type Log struct {
	Dir   string
	Level string
}

type Mysql struct {
	Host string
	Port int
	User string
	Pwd  string
	Db   string
}

type Redis struct {
	Host string
	Port int
	User string
	Pwd  string
	Db   int
}
