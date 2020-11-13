package database

import (
	"fmt"
	_const "github.com/imminoglobulin/file-service/pkg/const"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"os"
	"strconv"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     _const.DefaultDbHost,
		Port:     _const.DefaultDbPort,
		User:     _const.DefaultDbUser,
		Password: _const.DefaultDbPassword,
		DBName:   _const.DefaultDbName,
	}

	dbHost := os.Getenv("DATABASE_HOST")
	if dbHost != "" {
		dbConfig.Host = dbHost
	}

	dbPort := os.Getenv("DATABASE_PORT")
	if dbPort != "" {
		parseInt, err := strconv.ParseInt(dbPort, 10, 64)
		if err != nil {
			log.Err(err).Msg("DATABASE_PORT must be int value.")
			panic(err)
		}
		dbConfig.Port = int(parseInt)
	}

	dbUser := os.Getenv("DATABASE_USER")
	if dbUser != "" {
		dbConfig.User = dbUser
	}

	dbPassword := os.Getenv("DATABASE_PASSWORD")
	if dbPassword != "" {
		dbConfig.Password = dbPassword
	}

	dbSchema := os.Getenv("DATABASE_SCHEMA")
	if dbSchema != "" {
		dbConfig.DBName = dbSchema
	}

	return &dbConfig
}
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func Migration() {
	err := DB.AutoMigrate(&ApplicationUser{})

	if err != nil {
		errors.Wrap(err, "Db migration error !")
	}

	CreateDefaultUser()
}
