package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

// Config Declare a global configuration variable.
var Config MainConfig

// ServerConfig 代表服务器配置
type ServerConfig struct {
	Address      string `toml:"address"`
	Port         int    `toml:"port"`
	ReadTimeout  int    `toml:"read_timeout"`
	WriteTimeout int    `toml:"write_timeout"`
}

// PostgresDBConfig 代表PostgreSQL数据库配置
type PostgresDBConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
	SSLMode  string `toml:"sslmode"`
	MaxOpen  int `toml:"max_open"`
	MaxIdle  int `toml:"max_idle"`
}

// MySQLConfig 代表MySQL数据库配置
type MySQLConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
	Charset  string `toml:"charset,omitempty"` // 可选，默认值可省略
	SSLMode  string `toml:"sslmode,omitempty"` // 可选，默认或根据需要设置
	MaxOpen  int `toml:"max_open"`
	MaxIdle  int `toml:"max_idle"`
}

// RedisCacheConfig 代表Redis缓存配置
type RedisCacheConfig struct {
	Address  string `toml:"address"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
	PoolSize int    `toml:"pool_size"`
}

// KafkaConfig 代表Kafka消息队列配置
type KafkaConfig struct {
	Brokers []string `toml:"brokers"`
	GroupID string   `toml:"group_id"`
	Topics  []string `toml:"topics"`
}

// LoggingConfig 代表日志配置
type LoggingConfig struct {
	Level      string `toml:"level"`
	Format     string `toml:"format"`
	Filename   string `toml:"filename"`
	MaxSize    int    `toml:"max_size"` // 单位是MB
	MaxBackups int    `toml:"max_backups"`
	MaxAge     int    `toml:"max_age"` // 单位是天
}

// MonitoringConfig 代表监控配置
type MonitoringConfig struct {
	PrometheusEndpoint string `toml:"prometheus_endpoint"`
	LogFormat          string `toml:"log_format"`
}

// AuthConfig 代表认证配置
type AuthConfig struct {
	JWTSecret     string        `toml:"jwt_secret"`
	JWTEXpiration time.Duration `toml:"jwt_expiration"` // 直接使用time.Duration类型，需要字符串转义
}

// MainConfig 作为整个配置的根结构体，包含所有子配置
type MainConfig struct {
	Server     ServerConfig     `toml:"server"`
	Postgres   PostgresDBConfig `toml:"postgres"`
	MySQL      MySQLConfig      `toml:"mysql"`
	Redis      RedisCacheConfig `toml:"redis"`
	Kafka      KafkaConfig      `toml:"kafka"`
	Logging    LoggingConfig    `toml:"logging"`
	Monitoring MonitoringConfig `toml:"monitoring"`
	Auth       AuthConfig       `toml:"auth"`
}

// loadConfig loads the configuration from the TOML file into the Config variable.
func loadConfig(configFilePath string) error {
	_, err := toml.DecodeFile(configFilePath, &Config)
	if err != nil {
		return fmt.Errorf("error decoding config file: %w", err)
	}
	return nil
}

// InitConfig  initializes the Config variable by loading the configuration from the default or provided file.
func InitConfig() {
	defaultConfigPath := "config.toml"
	if configFilePath := os.Getenv("CONFIG_PATH"); configFilePath != "" {
		defaultConfigPath = configFilePath
	}

	if err := loadConfig(defaultConfigPath); err != nil {
		log.Fatalf("Failed to load config: %v", err)
		os.Exit(1)
	}
	log.Println("Configuration loaded successfully.")
}
