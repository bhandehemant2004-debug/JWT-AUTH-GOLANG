package routers

import (
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/config"
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/controller"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type UserRouter struct{
	UserController controller.UserControllerInterface
	logger *zap.Logger
	serverConfig *config.ServerConfig
}

func (UserRouter *UserRouter) Register(router *chi.Mux){
	router.Route("/api/v1/users",func(r chi.Router) {
		r.Get("/",UserRouter.UserController.GetAllUser)
		r.Get("/{id}",UserRouter.UserController.GetUserById)
		r.Post("/",UserRouter.UserController.CreateUser)
		r.Put("/{id}",UserRouter.UserController.UpdateUserById)
		r.Delete("/{id}",UserRouter.UserController.DeleteUserById)
	})
}
func NewUserRouter(usercontroller controller.UserControllerInterface, logger *zap.Logger , serverConfig *config.ServerConfig)RouterInterface{
	userRouter := &UserRouter{
		UserController: usercontroller,
		logger: logger,
		serverConfig: serverConfig,
	}
	return userRouter
}