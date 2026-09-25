package main

import (
	"log"

	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/cmd/app"
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/config"
)



func main() {
	//create the server config instance

	serverConfig ,err := config.LoadServerConfig();

	if err!= nil{
		log.Fatal("something went wrong while creating our server config instance"+err.Error())
	}

	//create the db config instance 
	dbConfig , err := config.LoadDbConfig()

	if err!= nil{
		log.Fatal("something went wrong while creating db server config instance"+err.Error())
	}


	//create the app instance 

	serverapp := &app.App{
		ServerConfig: serverConfig,
		DbConfig: dbConfig,
	}


	//run the app instance

	serverapp.Run()

}



