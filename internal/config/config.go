package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Configuration struct {
	Server     Server     `mapstructure:"server"`
	DataSource DataSource `mapstructure:"datasource"`
}

type Server struct {
	Name string `mapstructure:"mini-blog" json:"name"`
	Port int    `mapstructure:"port" json:"port"`
}

type DataSource struct {
	Mysql Mysql `mapstructure:"mysql" json:"mysql"`
}

type Mysql struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	DBName   string `mapstructure:"db_name" json:"db_name"`
}

var Cfg = &Configuration{}

func Init() {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./")
	v.AddConfigPath("./config")

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		panic("读取配置文件失败")
	}

	//
	if err := v.Unmarshal(&Cfg); err != nil {
		panic("配置映射失败")
	}

	// 启动看门狗，进行热加载
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed: ", in.Name)
		if err := v.Unmarshal(&Cfg); err != nil {
			panic("配置文件热加载失败")
		}
	})
}
