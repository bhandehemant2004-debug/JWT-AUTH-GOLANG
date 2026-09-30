package controller

import (
	"net/http"

	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/config"
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/services"
	"go.uber.org/zap"
)

type UserRoleControllerInterface interface {
	CreateUserRole(resWriter http.ResponseWriter, req *http.Request)
	GetAllUserRole(resWriter http.ResponseWriter, req *http.Request)
	GetUserRoleById(resWriter http.ResponseWriter, req *http.Request)
	UpdateUserRoleById(resWriter http.ResponseWriter, req *http.Request)
	DeleteUserRoleById(resWriter http.ResponseWriter, req *http.Request)

	GetRolesOfUser(resWriter http.ResponseWriter, req *http.Request)
	AssaignRoleToUser(resWriter http.ResponseWriter, req *http.Request)
	RemoveRoleFromUser(resWriter http.ResponseWriter, req *http.Request)
	CheckUserHasSingleRole(resWriter http.ResponseWriter, req *http.Request)
	CheckUserHasAnyRole(resWriter http.ResponseWriter, req *http.Request)
	CheckUserHasAllRole(resWriter http.ResponseWriter, req *http.Request)
}

type UserRoleController struct {
	userroleservice services.UserRoleServiceInterface
	logger          *zap.Logger
	serverConfig    *config.ServerConfig
}

func NewUserRoleController(userroleservice services.UserRoleServiceInterface,logger *zap.Logger,serverConfig *config.ServerConfig,
) UserRoleControllerInterface {
	return &UserRoleController{
		userroleservice: userroleservice,
		logger:          logger,
		serverConfig:    serverConfig,
	}
}

func (controller *UserRoleController) CreateUserRole(
	resWriter http.ResponseWriter,
	req *http.Request,
) {
	controller.userroleservice.CreateUserRole()
}

func (controller *UserRoleController) GetAllUserRole(
	resWriter http.ResponseWriter,
	req *http.Request,
) {
	controller.userroleservice.GetAllUserRole()
}

func (controller *UserRoleController) GetUserRoleById(
	resWriter http.ResponseWriter,
	req *http.Request,
) {
	controller.userroleservice.GetUserRoleById()
}

func (controller *UserRoleController) UpdateUserRoleById(
	resWriter http.ResponseWriter,
	req *http.Request,
) {
	controller.userroleservice.UpdateUserRoleById()
}

func (controller *UserRoleController) DeleteUserRoleById(
	resWriter http.ResponseWriter,
	req *http.Request,
) {
	controller.userroleservice.DeleteUserRoleById()
}

func (controller *UserRoleController) GetRolesOfUser(
	resWriter http.ResponseWriter,
	req *http.Request,
) {
	controller.userroleservice.GetRolesOfUser()
}

func (controller *UserRoleController) AssaignRoleToUser(
	resWriter http.ResponseWriter,
	req *http.Request,
) {
	controller.userroleservice.AssaignRoleToUser()
}

func (controller *UserRoleController) RemoveRoleFromUser(
	resWriter http.ResponseWriter,
	req *http.Request,
) {
	controller.userroleservice.RemoveRoleFromUser()
}

func (controller *UserRoleController) CheckUserHasSingleRole(
	resWriter http.ResponseWriter,
	req *http.Request,
) {
	controller.userroleservice.CheckUserHasSingleRole()
}

func (controller *UserRoleController) CheckUserHasAnyRole(
	resWriter http.ResponseWriter,
	req *http.Request,
) {
	controller.userroleservice.CheckUserHasAnyRole()
}

func (controller *UserRoleController) CheckUserHasAllRole(
	resWriter http.ResponseWriter,
	req *http.Request,
) {
	controller.userroleservice.CheckUserHasAllRole()
}
