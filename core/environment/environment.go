package environment

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

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

	_ = godotenv.Load(".env")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Bind env for known fields (both styles: APPPORT and APP_PORT)
	t := reflect.TypeOf(cfg)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		upper := strings.ToUpper(field.Name)
		viper.BindEnv(upper)
		viper.BindEnv(strings.ReplaceAll(upper, "PORT", "_PORT"))
		viper.BindEnv(strings.ReplaceAll(upper, "HOST", "_HOST"))
		viper.BindEnv(strings.ReplaceAll(upper, "ENV", "_ENV"))
		viper.BindEnv(strings.ReplaceAll(upper, "NAME", "_NAME"))
		viper.BindEnv(strings.ReplaceAll(upper, "PROTOCOL", "_PROTOCOL"))
	}

	// Read values explicitly to avoid type issues and support underscore variants
	cfg.AppName = firstNonEmpty(viper.GetString("APPNAME"), viper.GetString("APP_NAME"))
	cfg.AppEnv = firstNonEmpty(viper.GetString("APPENV"), viper.GetString("APP_ENV"))
	cfg.AppHost = firstNonEmpty(viper.GetString("APPHOST"), viper.GetString("APP_HOST"))
	cfg.AppProtocol = firstNonEmpty(viper.GetString("APPPROTOCOL"), viper.GetString("APP_PROTOCOL"))

	// Port: prefer int getters; fall back to parsing string
	port := viper.GetInt("APPPORT")
	if port == 0 {
		port = viper.GetInt("APP_PORT")
	}
	if port == 0 {
		ps := firstNonEmpty(viper.GetString("APPPORT"), viper.GetString("APP_PORT"))
		if ps != "" {
			if p, perr := strconv.Atoi(ps); perr == nil {
				port = p
			}
		}
	}
	cfg.AppPort = port

	cfg.DatabaseName = firstNonEmpty(viper.GetString("DATABASENAME"), viper.GetString("DATABASE_NAME"))
	cfg.DatabaseUsr = firstNonEmpty(viper.GetString("DATABASEUSR"), viper.GetString("DATABASE_USR"))
	cfg.DatabasePw = firstNonEmpty(viper.GetString("DATABASEPW"), viper.GetString("DATABASE_PW"))
	cfg.DatabasePort = firstNonEmpty(viper.GetString("DATABASEPORT"), viper.GetString("DATABASE_PORT"))
	cfg.DatabaseHost = firstNonEmpty(viper.GetString("DATABASEHOST"), viper.GetString("DATABASE_HOST"))

	// Minimal validation: require essential fields
	if cfg.AppPort <= 0 {
		return nil, fmt.Errorf("invalid AppPort (from .env): %d", cfg.AppPort)
	}
	if cfg.AppProtocol == "" {
		return nil, fmt.Errorf("invalid AppProtocol (from .env): empty")
	}

	return &cfg, nil
}

func firstNonEmpty(a, b string) string {
	if a != "" {
		return a
	}
	return b
}
