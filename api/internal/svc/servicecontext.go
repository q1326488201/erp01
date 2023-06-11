package svc

import (
	"app/api/internal/config"
	"app/dao/query"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	DB, err := gorm.Open(mysql.Open(c.Mysql.Dsn))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	query.SetDefault(DB)
	return &ServiceContext{
		Config: c,
	}
}
