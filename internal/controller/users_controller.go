package controller


import (
	"net/http"

	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/config"
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/services"
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/utils"
	"go.uber.org/zap"
)

type UserControllerInterface interface{
	CreateUser(resWriter http.ResponseWriter,req *http.Request)
	GetAllUser(resWriter http.ResponseWriter,req *http.Request)
	GetUserById(resWriter http.ResponseWriter,req *http.Request)
	UpdateUserById(resWriter http.ResponseWriter,req *http.Request)
	DeleteUserById(resWriter http.ResponseWriter,req *http.Request)
	GetUserByEmail(resWriter http.ResponseWriter,req *http.Request)
	GetUserByUsernameAndEmail(resWriter http.ResponseWriter,req *http.Request)
}

type UserController struct{
	userService services.UserServiceInterface
	logger *zap.Logger
	serverConfig *config.ServerConfig
}


func (UserController *UserController)CreateUser( resWriter http.ResponseWriter , req *http.Request){
	UserController.userService.CreateUser()

	utils.WriteJsonResponse(http.StatusCreated,resWriter,map[string]any{
		"success":true,
		"message": "A user is created ",
	})
}

func (UserController *UserController) GetAllUser(resWriter http.ResponseWriter,req *http.Request,) {
	UserController.userService.GetAllUser()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "All users fetched",
	})
}

func (UserController *UserController) GetUserById(resWriter http.ResponseWriter,req *http.Request,) {
	UserController.userService.GetUserById()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "User fetched successfully",
	})
}

func (UserController *UserController) UpdateUserById(resWriter http.ResponseWriter,req *http.Request,) {
	UserController.userService.UpdateUserById()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "User updated successfully",
	})
}

func (UserController *UserController) DeleteUserById(resWriter http.ResponseWriter,req *http.Request,) {
	UserController.userService.DeleteUserById()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "User deleted successfully",
	})
}

func (UserController *UserController) GetUserByEmail(resWriter http.ResponseWriter,req *http.Request,) {
	UserController.userService.GetUserByEmail()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "User fetched successfully",
	})
}

func (UserController *UserController) GetUserByUsernameAndEmail(resWriter http.ResponseWriter,req *http.Request,) {
	UserController.userService.GetUserByUsernameAndEmail()

	utils.WriteJsonResponse(http.StatusOK, resWriter, map[string]any{
		"success": true,
		"message": "User fetched successfully",
	})
}

func NewUserController(UserService services.UserServiceInterface, logger *zap.Logger , serverConfig *config.ServerConfig)UserControllerInterface{
	usercontroller :=  &UserController{
		userService: UserService,
		logger: logger,
		serverConfig: serverConfig,
	}
	return  usercontroller
}