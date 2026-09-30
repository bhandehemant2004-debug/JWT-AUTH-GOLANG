package services

import (
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/config"
	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/repositories"
	"go.uber.org/zap"
)

type RoleServiceInterface interface {
	CreateRole()
	GetAllRole()
	GetRoleById()
	UpdateRoleById()
	DeleteRoleById()
	GetRoleByName()
}

type RoleService struct {
	RoleRepository repositories.RoleRepositoryInterface
	logger         *zap.Logger
	serverConfig   *config.ServerConfig
}

// CreateRole implements [RoleServiceInterface].
func (roleService *RoleService) CreateRole() {
	roleService.logger.Info("roleService -> CreateRole")
	roleService.RoleRepository.CreateRole()
}

// DeleteRoleById implements [RoleServiceInterface].
func (roleService *RoleService) DeleteRoleById() {
	roleService.logger.Info("roleService -> DeleteRoleById")
	roleService.RoleRepository.DeleteRoleById()
}

// GetAllRole implements [RoleServiceInterface].
func (roleService *RoleService) GetAllRole() {
	roleService.logger.Info("roleService -> GetAllRole()")
	roleService.RoleRepository.GetAllRole()
}

// GetRoleById implements [RoleServiceInterface].
func (roleService *RoleService) GetRoleById() {
	roleService.logger.Info("roleService -> GetRoleById()")
	roleService.RoleRepository.GetRoleById()
}

// GetRoleByName implements [RoleServiceInterface].
func (roleService *RoleService) GetRoleByName() {
	roleService.logger.Info("roleService -> GetRoleByName()")
	roleService.RoleRepository.GetRoleByName()
}

// UpdateRoleById implements [RoleServiceInterface].
func (roleService *RoleService) UpdateRoleById() {
	roleService.logger.Info("roleService -> UpdateRoleById()")
	roleService.RoleRepository.UpdateRoleById()
}

func NewRoleService(roleRepository repositories.RoleRepositoryInterface, logger *zap.Logger, serverConfig *config.ServerConfig) RoleServiceInterface {
	roleService := &RoleService{
		RoleRepository: roleRepository,
		logger:         logger,
		serverConfig:   serverConfig,
	}

	return roleService
}
