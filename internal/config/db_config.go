package config

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type DbConfig struct {
	DbUsername string `validate:"required"`
	Dbpassword string `validate:"required"`
	DbNet      string `validate:"required"`
	DbAddress  string `validate:"required"`
	DbName     string `validate:"required"`
}

func LoadDbConfig() (*DbConfig, error) {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("something went wrong while loading the env vars:" + err.Error())
		return nil, err
	}

	cfg := &DbConfig{
		DbUsername: LoadSingleEnvVar("DB_USERNAME", "root"),
		Dbpassword: LoadSingleEnvVar("DB_PASSWORD", "root"),
		DbNet:      LoadSingleEnvVar("DB_NET", "tcp"),
		DbAddress:  LoadSingleEnvVar("DB_ADDRESS", "127.0.0.1:3306"),
		DbName:     LoadSingleEnvVar("DB_NAME", "auth_db"),
	}
	return cfg, nil
}

func SetupDB(dbconfig *DbConfig, logger *zap.Logger) (*sql.DB, error) {

	cfg := mysql.NewConfig()
	cfg.User = dbconfig.DbUsername
	cfg.Passwd = dbconfig.Dbpassword
	cfg.Net = dbconfig.DbNet
	cfg.Addr = dbconfig.DbAddress
	cfg.DBName = dbconfig.DbName

	//Open a new sql conection

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		logger.Fatal("something went wrong while opening db connection ",
			zap.String("error", err.Error()))
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		logger.Fatal("something went wrong while pinging db connection ",
			zap.String("error", err.Error()))
		return nil, err
	}

	logger.Info("Successfully connected to the db ",
		zap.String("db_name", dbconfig.DbName))

	return db, nil

}
