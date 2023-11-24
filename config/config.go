package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"time"
)

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

type Config struct {
	Name        string
	Version     string
	Port        string
	Level       string
	MaxPoolSize int
	URL         string
	SignKey     string
	TokenTTL    time.Duration
	Salt        string
}

func InitConfig(configPath, configName string) error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	return viper.ReadInConfig()
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg := &Config{
		Name:        viper.GetString("app.name"),
		Version:     viper.GetString("app.version"),
		Port:        viper.GetString("http.port"),
		Level:       viper.GetString("log.level"),
		MaxPoolSize: viper.GetInt("postgres.max_pool_size"),
		TokenTTL:    viper.GetDuration("jwt.token_ttl"),
		URL:         getEnv("PG_URL", ""),
		SignKey:     getEnv("JWT_SIGN_KEY", ""),
		Salt:        getEnv("HASHER_SALT", ""),
	}

	return cfg, nil
}
