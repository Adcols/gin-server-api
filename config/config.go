package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Port    int    `mapstructure:"port"`
	Version string `mapstructure:"version"`
}

type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Charset  string `mapstructure:"charset"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type JWTConfig struct {
	Secret     string `mapstructure:"secret"`
	Expiration int    `mapstructure:"expiration"`
}

var GlobalConfig Config

// LoadConfig 加载配置文件
func LoadConfig(configPath string) (config Config, err error) {
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	GlobalConfig = config
	return
}

// InitConfig 初始化配置
func InitConfig() {
	workDir, _ := os.Getwd()
	configPath := fmt.Sprintf("%s/etc/config.yaml", workDir)

	// 检查配置文件是否存在，不存在则创建默认配置
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Println("配置文件不存在，创建默认配置")
		createDefaultConfig(configPath)
	}

	config, err := LoadConfig(configPath)
	if err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}

	log.Printf("配置加载成功: %s, 运行模式: %s", config.App.Name, config.App.Mode)
}

// 创建默认配置文件
func createDefaultConfig(configPath string) {
	defaultConfig := `app:
  name: gin-server-api
  mode: development
  port: 8080
  version: 1.0.0

database:
  driver: mysql
  host: localhost
  port: 3306
  username: root
  password: password
  database: gin_server
  charset: utf8mb4

redis:
  host: localhost
  port: 6379
  password: ""
  db: 0

jwt:
  secret: "your-secret-key"
  expiration: 86400 # 24小时
`

	// 确保目录存在
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("创建配置目录失败: %v", err)
	}

	err := os.WriteFile(configPath, []byte(defaultConfig), 0644)
	if err != nil {
		log.Fatalf("创建默认配置文件失败: %v", err)
	}
}
