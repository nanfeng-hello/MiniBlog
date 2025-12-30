package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nanfeng/mini-blog/internal/config"
	"github.com/nanfeng/mini-blog/internal/handler/admin"
	"github.com/nanfeng/mini-blog/internal/repository"
	"github.com/nanfeng/mini-blog/internal/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		})

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("数据库连接异常, err: " + err.Error())
	}
	user_repo := repository.NewUserRepository(db)
	user_service := service.NewUserService(user_repo)
	user_handler := admin.NewUserHandler(user_service)

	r := gin.Default()
	v1 := r.Group("/api/v1/")

	gin.SetMode(config.Cfg.Server.Mode)

	user_handler.Register(v1)

	r.Run()
}
