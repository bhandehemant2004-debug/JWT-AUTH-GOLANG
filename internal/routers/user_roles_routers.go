
package routers

import (
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/config"
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/controller"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type UserRoleRouter struct{
	UserRoleController controller.UserRoleControllerInterface
	logger *zap.Logger
	serverConfig *config.ServerConfig
}

func (UserRoleRouter *UserRoleRouter) Register(router *chi.Mux){
	router.Route("/api/v1/users-roles",func(r chi.Router) {
		r.Get("/user/{user_id}",UserRoleRouter.UserRoleController.GetRolesOfUser)
		r.Post("/assign",UserRoleRouter.UserRoleController.AssaignRoleToUser)
		r.Post("/remove",UserRoleRouter.UserRoleController.RemoveRoleFromUser)
		r.Get("/check-single",UserRoleRouter.UserRoleController.CheckUserHasSingleRole)
		r.Get("/check-all",UserRoleRouter.UserRoleController.CheckUserHasAllRole)
		r.Get("/check-any",UserRoleRouter.UserRoleController.CheckUserHasAnyRole)
	})
}
func NewUserRoleRouter(userRolecontroller controller.UserRoleControllerInterface, logger *zap.Logger , serverConfig *config.ServerConfig)RouterInterface{
	userRoleRouter := &UserRoleRouter{
		UserRoleController: userRolecontroller,
		logger: logger,
		serverConfig: serverConfig,
	}
	return userRoleRouter
}