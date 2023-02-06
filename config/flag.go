package config

import "flag"

var Env string

func init() {
	flag.StringVar(&Env, "env", "local", "env model, optional local/dev/prod")
}
