package routers


import (
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/config"
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/controller"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type RoleRouter struct{
	RoleController controller.RoleContorllerInterface
	logger *zap.Logger
	serverConfig *config.ServerConfig
}

func (roleRouter *RoleRouter) Register(router *chi.Mux){
	router.Route("/api/v1/roles",func(r chi.Router) {
		r.Get("/",roleRouter.RoleController.GetAllRole)
		r.Get("/{id}",roleRouter.RoleController.GetRoleById)
		r.Post("/",roleRouter.RoleController.CreateRole)
		r.Put("/{id}",roleRouter.RoleController.UpdateRoleById)
		r.Delete("/{id}",roleRouter.RoleController.DeleteRoleById)
	})
}
func NewRoleRouter(roleController controller.RoleContorllerInterface, logger *zap.Logger , serverConfig *config.ServerConfig)RouterInterface{
	RoleRouter := &RoleRouter{
		RoleController: roleController,
		logger: logger,
		serverConfig: serverConfig,
	}
	return RoleRouter
}