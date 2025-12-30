package main

import (
	"fmt"

	"github.com/nanfeng/mini-blog/internal/config"
	"github.com/nanfeng/mini-blog/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// 1.加载配置
	config.Init()
	// 2.实例
	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Cfg.DataSource.Mysql.User,
		config.Cfg.DataSource.Mysql.Password,
		config.Cfg.DataSource.Mysql.Host,
		config.Cfg.DataSource.Mysql.Port,
		config.Cfg.DataSource.Mysql.DBName,
	)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("数据库连接异常, err: " + err.Error())
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.Comment{})
}
