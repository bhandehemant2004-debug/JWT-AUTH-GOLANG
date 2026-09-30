package services

import (
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/config"
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/repositories"
	"go.uber.org/zap"
)

type UserServiceInterface interface {
	CreateUser()
	GetAllUser()
	GetUserById()
	UpdateUserById()
	DeleteUserById()
	GetUserByEmail()
	GetUserByUsernameAndEmail()
}

type UserService struct {
	UserRepository repositories.UserRepositoryInterface
	logger       *zap.Logger
	serverConfig *config.ServerConfig
}

// CreateUser implements [UserServiceInterface].
func (userService *UserService) CreateUser() {
	userService.logger.Info("userService -> CreatUser")
	userService.UserRepository.CreateUser()
}

// DeleteUserById implements [UserServiceInterface].
func ( userService *UserService) DeleteUserById() {
	userService.logger.Info("userService -> DeleteUserById")
	userService.UserRepository.DeleteUserById()
}

// GetAllUser implements [UserServiceInterface].
func (userService *UserService) GetAllUser() {
	userService.logger.Info("userService -> GetAllUser()")
	userService.UserRepository.GetAllUser()
}

// GetUserByEmail implements [UserServiceInterface].
func (userService *UserService) GetUserByEmail() {
	userService.logger.Info("userService -> GetUserByEmail()")
	userService.UserRepository.GetUserByEmail()
}

// GetUserById implements [UserServiceInterface].
func (userService *UserService) GetUserById() {
	userService.logger.Info("userService -> GetUserById()")
	userService.UserRepository.GetUserById()
}

// GetUserByUsernameAndEmail implements [UserServiceInterface].
func (userService *UserService) GetUserByUsernameAndEmail() {
	userService.logger.Info("userService -> GetUserByUsernameAndEmail()")
	userService.UserRepository.GetUserByUsernameAndEmail()

}

// UpdateUserById implements [UserServiceInterface].
func (userService *UserService) UpdateUserById() {
	userService.logger.Info("userService -> updateUserById()")
	userService.UserRepository.UpdateUserById()
}

func NewUserService(userRepository repositories.UserRepositoryInterface , logger *zap.Logger ,serverConfig *config.ServerConfig)UserServiceInterface{
	userService := &UserService{
		UserRepository: userRepository,
		logger: logger,
		serverConfig: serverConfig,
	}
	return userService
}
