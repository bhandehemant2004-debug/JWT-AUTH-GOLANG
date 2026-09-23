package config

import (
    "log"
	"time"
    "github.com/joho/godotenv"
)

type ServerConfig struct{
	Port string `validate:"required"`
	ReadTimeout time.Duration `validate:"required"`
	WriteTimeout time.Duration `validate:"required"`
	IdleTimeout time.Duration `validate:"required"`
	AppEnv string `validate:"required"`
}

func LoadServerConfig()(*ServerConfig,error){

	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  cnf:= &ServerConfig{
	Port: LoadSingleEnvVar("PORT",":3001"),
	ReadTimeout: time.Duration(LoadSingleEnvVar("READTIME_OUT",15))*time.Second ,
	WriteTimeout: time.Duration(LoadSingleEnvVar("WRITETIME_OUT",15))*time.Second ,
	IdleTimeout: time.Duration(LoadSingleEnvVar("READTIME_OUT",15))*time.Second ,
	AppEnv: LoadSingleEnvVar("APP_ENV","developement"),
  }

  return cnf,nil

}