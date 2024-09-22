package svc

import (
	"bookservice/internal/config"
	"bookservice/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Book{})
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
