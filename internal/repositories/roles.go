package repositories

import (
	"database/sql"

	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/config"
	"go.uber.org/zap"
)

type RoleRepositoryInterface interface {
	CreateRole()
	GetAllRole()
	GetRoleById()
	UpdateRoleById()
	DeleteRoleById()
	GetRoleByName()
}

type RoleRepository struct {
	db           *sql.DB
	logger       *zap.Logger
	serverConfig *config.ServerConfig
}

// CreateRole implements [RoleRepositoryInterface].
func (roleRepository *RoleRepository) CreateRole() {
	roleRepository.logger.Info("roleRepository->CreateRole")
}

// DeleteRoleById implements [RoleRepositoryInterface].
func (roleRepository *RoleRepository) DeleteRoleById() {
	roleRepository.logger.Info("roleRepository->DeleteRoleById")
}

// GetAllRole implements [RoleRepositoryInterface].
func (roleRepository *RoleRepository) GetAllRole() {
	roleRepository.logger.Info("roleRepository->GetAllRole")
}

// GetRoleById implements [RoleRepositoryInterface].
func (roleRepository *RoleRepository) GetRoleById() {
	roleRepository.logger.Info("roleRepository->GetRoleById")
}

// GetRoleByName implements [RoleRepositoryInterface].
func (roleRepository *RoleRepository) GetRoleByName() {
	roleRepository.logger.Info("roleRepository->GetRoleByName")
}

// UpdateRoleById implements [RoleRepositoryInterface].
func (roleRepository *RoleRepository) UpdateRoleById() {
	roleRepository.logger.Info("roleRepository->UpdateRoleById")
}

func NewRoleRepository(db *sql.DB,logger *zap.Logger,serverConfig *config.ServerConfig,) *RoleRepository {
	roleRepository := &RoleRepository{
		db:           db,
		logger:       logger,
		serverConfig: serverConfig,
	}

	return roleRepository
}
