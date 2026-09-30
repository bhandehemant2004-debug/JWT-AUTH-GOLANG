package controller

import (
	"net/http"

	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/config"
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/services"
	"go.uber.org/zap"
)


type RoleContorllerInterface interface{
	CreateRole(resWriter http.ResponseWriter,req *http.Request)
	GetAllRole(resWriter http.ResponseWriter,req *http.Request)
	GetRoleById(resWriter http.ResponseWriter,req *http.Request)
	UpdateRoleById(resWriter http.ResponseWriter,req *http.Request)
	DeleteRoleById(resWriter http.ResponseWriter,req *http.Request)
	GetRoleByName(resWriter http.ResponseWriter,req *http.Request)
}


type RoleController struct{
	roleController services.RoleServiceInterface
	logger *zap.Logger
	serverConfig *config.ServerConfig
}

func NewRoleController(roleController services.RoleServiceInterface,logger *zap.Logger,serverConfig *config.ServerConfig,) RoleContorllerInterface {
	return &RoleController{
		roleController: roleController,
		logger:         logger,
		serverConfig:   serverConfig,
	}
}

func (roleController *RoleController) CreateRole(resWriter http.ResponseWriter,req *http.Request,) {
	roleController.roleController.CreateRole()
}

func (roleController *RoleController) GetAllRole(resWriter http.ResponseWriter,req *http.Request,) {
	roleController.roleController.GetAllRole()
}

func (roleController *RoleController) GetRoleById(resWriter http.ResponseWriter,req *http.Request,) {
	roleController.roleController.GetRoleById()
}

func (roleController *RoleController) UpdateRoleById(resWriter http.ResponseWriter,req *http.Request,) {
	roleController.roleController.UpdateRoleById()
}

func (roleController *RoleController) DeleteRoleById(resWriter http.ResponseWriter,req *http.Request,) {
	roleController.roleController.DeleteRoleById()
}

func (roleController *RoleController) GetRoleByName(resWriter http.ResponseWriter,req *http.Request,) {
	roleController.roleController.GetRoleByName()
}