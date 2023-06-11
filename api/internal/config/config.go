package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql struct {
		Dsn string `json:"Dsn" yaml:"Dsn"`
	} `json:"MySql" yaml:"MySql"`
}
