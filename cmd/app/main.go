package app

import (
	"log"
	"net/http"

	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/config"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type AppInterface interface{
	Run()
}

type App struct{
	ServerConfig *config.ServerConfig
	DbConfig *config.DbConfig
}

func (app *App)Run(){
	//validate server config
	validate := validator.New()
	
	if err:= validate.Struct(app.ServerConfig); err != nil{
		log.Fatal("Failed to validate the server config"+err.Error())
	}

	//validate db config

	if err:= validate.Struct(app.DbConfig); err != nil{
		log.Fatal("Failed to validate the server config"+err.Error())
	}

	//setup zap logger 
	logger := config.GetLogger(app.ServerConfig.AppEnv)

	// set up db 

	db , err := config.SetupDB(app.DbConfig,logger)

	if err != nil{
		log.Fatal("Something went wrong while setting up the db:"+err.Error() )
	}

	defer db.Close()

	// create a server instance 

	server := &http.Server{
		Addr: 		app.ServerConfig.Port,
		Handler:    nil,
		ReadTimeout: app.ServerConfig.ReadTimeout,
		WriteTimeout: app.ServerConfig.WriteTimeout,
		IdleTimeout: app.ServerConfig.IdleTimeout,
	}


	// start a server 

	logger.Info("Starting the server...",zap.String("port",app.ServerConfig.Port))

	err  = server.ListenAndServe()

	if err != nil{
		log.Fatal("Something went wrong while starting a server",zap.String("error",err.Error()))
	}
}