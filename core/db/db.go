package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"github.com/xeventa/base-service/core/environment"
)

type MySQL struct {
	DB *sql.DB
}

func NewMySQL(config *environment.Config) *MySQL {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DatabaseUsr, config.DatabasePw, config.DatabaseHost, config.DatabasePort, config.DatabaseName)

	DB, err := sql.Open("mysql", conn)
	if err != nil {
		log.Panic()
	}

	// optional: test ping
	if err := DB.Ping(); err != nil {
		log.Panic()
	}

	return &MySQL{DB: DB}
}
