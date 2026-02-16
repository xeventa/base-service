package public

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xeventa/base-service/core/environment"
)

type Service struct {
	config *environment.Config
}

func NewService(config *environment.Config) *Service {
	return &Service{config: config}
}

func (s *Service) HealthCheck() interface{} {

	data := make(map[string]interface{})

	data["status"] = "OK"
	data["config"] = s.config
	return data
}

func (s *Service) DBPing() interface{} {
	resp := make(map[string]interface{})

	// Build DSN following core/db/db.go convention
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		s.config.DatabaseUsr,
		s.config.DatabasePw,
		s.config.DatabaseHost,
		s.config.DatabasePort,
		s.config.DatabaseName,
	)

	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		resp["status"] = "ERROR"
		resp["error"] = err.Error()
		return resp
	}
	defer DB.Close()

	if err := DB.Ping(); err != nil {
		resp["status"] = "ERROR"
		resp["error"] = err.Error()
		return resp
	}

	resp["status"] = "OK"
	return resp
}
