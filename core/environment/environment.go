package environment

import (
	"fmt"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	AppName     string
	AppEnv      string
	AppHost     string
	AppPort     int
	AppProtocol string

	DatabaseName string
	DatabaseUsr  string
	DatabasePw   string
	DatabasePort string
	DatabaseHost string
}

func ProvideConfig() (*Config, error) {
	var cfg Config

	// Load .env into process env
	_ = godotenv.Load(".env")

	// Read only from environment; single canonical keys
	viper.AutomaticEnv()

	cfg.AppName = viper.GetString("APP_NAME")
	cfg.AppEnv = viper.GetString("APP_ENV")
	cfg.AppHost = viper.GetString("APP_HOST")
	cfg.AppProtocol = viper.GetString("APP_PROTOCOL")

	// Port: prefer int getter, fallback to string parse
	port := viper.GetInt("APP_PORT")
	if port == 0 {
		ps := viper.GetString("APP_PORT")
		if ps != "" {
			if p, perr := strconv.Atoi(ps); perr == nil {
				port = p
			}
		}
	}
	cfg.AppPort = port

	cfg.DatabaseName = viper.GetString("DATABASE_NAME")
	cfg.DatabaseUsr = viper.GetString("DATABASE_USR")
	cfg.DatabasePw = viper.GetString("DATABASE_PW")
	cfg.DatabasePort = viper.GetString("DATABASE_PORT")
	cfg.DatabaseHost = viper.GetString("DATABASE_HOST")

	// Minimal validation
	if cfg.AppPort <= 0 {
		return nil, fmt.Errorf("invalid AppPort: %d", cfg.AppPort)
	}
	if cfg.AppProtocol == "" {
		return nil, fmt.Errorf("invalid AppProtocol: empty")
	}

	return &cfg, nil
}
