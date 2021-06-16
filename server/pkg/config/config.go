package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

var path = "./config/config.toml"
var serverConfig *MultiConfig

//SetConfigPath set config path
func SetConfigPath(p string) {
	if p != "" {
		path = p
	}
}

//DatabaseConfig db
type DatabaseConfig struct {
	Addr            string
	DB              string
	UserName        string
	Driver          string
	Password        string
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifeTime int
}

type RedisConfig struct {
	Addr     string
	Password string
}

type RabbitMQConfig struct {
	Mqurl string
}

type MultiConfig struct {
	Database DatabaseConfig `toml:"database"`
	Redis    RedisConfig    `toml:"redis"`
	RabbitMQ RabbitMQConfig `toml:"rabbitmq"`
}

func GetMultiConfig() *MultiConfig {
	if serverConfig != nil {
		return serverConfig
	}
	var mc MultiConfig
	if _, err := toml.DecodeFile(path, &mc); err != nil {
		log.Printf("Config Get failed. Path: %s", path)
		log.Fatal(err)
	}
	serverConfig = &mc
	log.Println(serverConfig)
	return serverConfig
}

//GetDBConfig get config
func GetDBConfig() *DatabaseConfig {
	mc := GetMultiConfig()
	return &(mc.Database)
}

//GetDBConfig get config
func GetRedisConfig() *RedisConfig {
	mc := GetMultiConfig()
	return &(mc.Redis)
}

//GetDBConfig get config
func GetRabbitMQConfig() *RabbitMQConfig {
	mc := GetMultiConfig()
	return &(mc.RabbitMQ)
}
