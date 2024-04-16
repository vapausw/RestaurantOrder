package setting

import (
	"gopkg.in/ini.v1"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Release        bool `ini:"release"`
	Port           int  `ini:"port"`
	*MySQLConfig   `ini:"mysql"`
	*MyEmailConfig `ini:"email"`
	*RedisConfig   `ini:"redis"`
	*LogConfig     `ini:"logger"`
}

type MyEmailConfig struct {
	Email    string `ini:"email"`
	Password string `ini:"password"`
}

type MySQLConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	DB       string `ini:"db"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
}

type RedisConfig struct {
	Addr     string `ini:"addr"`
	Password string `ini:"password"`
	DB       int    `ini:"db"`
	PoolSize int    `ini:"pool_size"`
}

type LogConfig struct {
	Level      string `ini:"level"`
	Filename   string `ini:"filename"`
	MaxSize    int    `ini:"maxsize"`
	MaxAge     int    `ini:"max_age"`
	MaxBackups int    `ini:"max_backups"`
}

func Init(file string) error {
	return ini.MapTo(Conf, file)
}
