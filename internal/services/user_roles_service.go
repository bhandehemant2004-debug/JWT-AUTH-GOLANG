
package services

import (
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/config"
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/repositories"
	"go.uber.org/zap"
)

type UserRoleServiceInterface interface {
	CreateUserRole()
	GetAllUserRole()
	GetUserRoleById()
	UpdateUserRoleById()
	DeleteUserRoleById()

	GetRolesOfUser()
	AssaignRoleToUser()
	RemoveRoleFromUser()
	CheckUserHasSingleRole()
	CheckUserHasAnyRole()
	CheckUserHasAllRole()
}

type UserRoleService struct {
	UserRoleRepository repositories.UserRoleRepositoryInterface
	logger             *zap.Logger
	serverConfig       *config.ServerConfig
}

// CreateUserRole implements [UserRoleServiceInterface].
func (userRoleService *UserRoleService) CreateUserRole() {
	userRoleService.logger.Info("userRoleService -> CreateUserRole")
	userRoleService.UserRoleRepository.CreateUserRole()
}

// DeleteUserRoleById implements [UserRoleServiceInterface].
func (userRoleService *UserRoleService) DeleteUserRoleById() {
	userRoleService.logger.Info("userRoleService -> DeleteUserRoleById")
	userRoleService.UserRoleRepository.DeleteUserRoleById()
}

// GetAllUserRole implements [UserRoleServiceInterface].
func (userRoleService *UserRoleService) GetAllUserRole() {
	userRoleService.logger.Info("userRoleService -> GetAllUserRole()")
	userRoleService.UserRoleRepository.GetAllUserRole()
}

// GetUserRoleById implements [UserRoleServiceInterface].
func (userRoleService *UserRoleService) GetUserRoleById() {
	userRoleService.logger.Info("userRoleService -> GetUserRoleById()")
	userRoleService.UserRoleRepository.GetUserRoleById()
}

// UpdateUserRoleById implements [UserRoleServiceInterface].
func (userRoleService *UserRoleService) UpdateUserRoleById() {
	userRoleService.logger.Info("userRoleService -> UpdateUserRoleById()")
	userRoleService.UserRoleRepository.UpdateUserRoleById()
}

// GetRolesOfUser implements [UserRoleServiceInterface].
func (userRoleService *UserRoleService) GetRolesOfUser() {
	userRoleService.logger.Info("userRoleService -> GetRolesOfUser()")
	userRoleService.UserRoleRepository.GetRolesOfUser()
}

// AssaignRoleToUser implements [UserRoleServiceInterface].
func (userRoleService *UserRoleService) AssaignRoleToUser() {
	userRoleService.logger.Info("userRoleService -> AssaignRoleToUser()")
	userRoleService.UserRoleRepository.AssaignRoleToUser()
}

// RemoveRoleFromUser implements [UserRoleServiceInterface].
func (userRoleService *UserRoleService) RemoveRoleFromUser() {
	userRoleService.logger.Info("userRoleService -> RemoveRoleFromUser()")
	userRoleService.UserRoleRepository.RemoveRoleFromUser()
}

// CheckUserHasSingleRole implements [UserRoleServiceInterface].
func (userRoleService *UserRoleService) CheckUserHasSingleRole() {
	userRoleService.logger.Info("userRoleService -> CheckUserHasSingleRole()")
	userRoleService.UserRoleRepository.CheckUserHasSingleRole()
}

// CheckUserHasAnyRole implements [UserRoleServiceInterface].
func (userRoleService *UserRoleService) CheckUserHasAnyRole() {
	userRoleService.logger.Info("userRoleService -> CheckUserHasAnyRole()")
	userRoleService.UserRoleRepository.CheckUserHasAnyRole()
}

// CheckUserHasAllRole implements [UserRoleServiceInterface].
func (userRoleService *UserRoleService) CheckUserHasAllRole() {
	userRoleService.logger.Info("userRoleService -> CheckUserHasAllRole()")
	userRoleService.UserRoleRepository.CheckUserHasAllRole()
}

func NewUserRoleService(userRoleRepository repositories.UserRoleRepositoryInterface,logger *zap.Logger,serverConfig *config.ServerConfig,) UserRoleServiceInterface{
	userRoleService := &UserRoleService{
		UserRoleRepository: userRoleRepository,
		logger:             logger,
		serverConfig:       serverConfig,
	}

	return userRoleService
}
