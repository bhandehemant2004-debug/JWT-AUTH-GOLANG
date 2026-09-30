package routers

import (
	"database/sql"

	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/config"
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/controller"
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/repositories"
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/services"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type RouterInterface interface {
	Register(router *chi.Mux)
}

func RegisterRouters(logger *zap.Logger,db *sql.DB , serverConfig *config.ServerConfig )*chi.Mux{
	//crate all router instance 

	router := chi.NewRouter()

	// create all repositories 

	userrepository := repositories.NewUserRepository(db,logger,serverConfig)

	rolerepository := repositories.NewRoleRepository(db,logger,serverConfig)

	userrolerepository := repositories.NewUserRoleRepository(db,logger,serverConfig)

	//create all services 
	userservice := services.NewUserService(userrepository,logger,serverConfig)

	roleservice := services.NewRoleService(rolerepository,logger,serverConfig)

	userroleservice := services.NewUserRoleService(userrolerepository,logger,serverConfig)



	//create all controller 


	usercontroller := controller.NewUserController(userservice,logger,serverConfig)
	rolecontroller := controller.NewRoleController(roleservice,logger,serverConfig)
	userrolecontroller := controller.NewUserRoleController(userroleservice,logger,serverConfig)

	//create all routers 

	userrouter  := NewUserRouter(usercontroller,logger,serverConfig)

	rolerouter  := NewRoleRouter(rolecontroller,logger,serverConfig)

	userrolerouter := NewUserRoleRouter(userrolecontroller,logger,serverConfig)

	//register them 


	

	userrouter.Register(router)
	rolerouter.Register(router)
	userrolerouter.Register(router)

	return router

	
}