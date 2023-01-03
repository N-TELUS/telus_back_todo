package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigList struct {
	DBDriverName   string
	DBName         string
	DBUserName     string
	DBUserPassword string
	DBHost         string
	DBPort         string
	ServerPort     int
}

// 外部パッケージで読み込めるようにパブリックで宣言
var Config ConfigList

func init() {
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		DBDriverName:   cfg.Section("db").Key("db_driver_name").String(),
		DBName:         cfg.Section("db").Key("db_name").String(),
		DBUserName:     cfg.Section("db").Key("db_user_name").String(),
		DBUserPassword: cfg.Section("db").Key("db_user_password").String(),
		DBHost:         cfg.Section("db").Key("db_host").String(),
		DBPort:         cfg.Section("db").Key("db_port").String(),
		ServerPort:     cfg.Section("api").Key("server_port").MustInt(),
	}
}

func NewDB() *gorm.DB {
	dsn := Config.DBUserName + ":" + Config.DBUserPassword + "@tcp(" + Config.DBHost + ")/" + Config.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// db.AutoMigrate(&domain.TodoList{})

	return db
}
