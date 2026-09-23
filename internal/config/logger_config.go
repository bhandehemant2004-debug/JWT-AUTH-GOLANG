package config

import (
	"fmt"

	"go.uber.org/zap"
)

func GetLogger(AppEnv string) *zap.Logger{	

	logger := zap.Must(zap.NewProduction())

	if AppEnv == "developement"{
		logger = zap.Must(zap.NewDevelopment())
	}
	defer func(logger *zap.Logger){
		error := logger.Sync()

		if error != nil{
			fmt.Println("Something error while sync the zap logger ")
		}	
	}(logger)

	return logger
}