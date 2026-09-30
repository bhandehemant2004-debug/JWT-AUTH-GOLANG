package repositories

import (
	"database/sql"

	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/config"
	"go.uber.org/zap"
)

type UserRoleRepositoryInterface interface {
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

type UserRoleRepository struct {
	db           *sql.DB
	logger       *zap.Logger
	serverConfig *config.ServerConfig
}

// CreateUserRole implements [UserRoleRepositoryInterface].
func (userRoleRepository *UserRoleRepository) CreateUserRole() {
	userRoleRepository.logger.Info("userRoleRepository->CreateUserRole")
}

// DeleteUserRoleById implements [UserRoleRepositoryInterface].
func (userRoleRepository *UserRoleRepository) DeleteUserRoleById() {
	userRoleRepository.logger.Info("userRoleRepository->DeleteUserRoleById")
}

// GetAllUserRole implements [UserRoleRepositoryInterface].
func (userRoleRepository *UserRoleRepository) GetAllUserRole() {
	userRoleRepository.logger.Info("userRoleRepository->GetAllUserRole")
}

// GetUserRoleById implements [UserRoleRepositoryInterface].
func (userRoleRepository *UserRoleRepository) GetUserRoleById() {
	userRoleRepository.logger.Info("userRoleRepository->GetUserRoleById")
}

// UpdateUserRoleById implements [UserRoleRepositoryInterface].
func (userRoleRepository *UserRoleRepository) UpdateUserRoleById() {
	userRoleRepository.logger.Info("userRoleRepository->UpdateUserRoleById")
}

// GetRolesOfUser implements [UserRoleRepositoryInterface].
func (userRoleRepository *UserRoleRepository) GetRolesOfUser() {
	userRoleRepository.logger.Info("userRoleRepository->GetRolesOfUser")
}

// AssaignRoleToUser implements [UserRoleRepositoryInterface].
func (userRoleRepository *UserRoleRepository) AssaignRoleToUser() {
	userRoleRepository.logger.Info("userRoleRepository->AssaignRoleToUser")
}

// RemoveRoleFromUser implements [UserRoleRepositoryInterface].
func (userRoleRepository *UserRoleRepository) RemoveRoleFromUser() {
	userRoleRepository.logger.Info("userRoleRepository->RemoveRoleFromUser")
}

// CheckUserHasSingleRole implements [UserRoleRepositoryInterface].
func (userRoleRepository *UserRoleRepository) CheckUserHasSingleRole() {
	userRoleRepository.logger.Info("userRoleRepository->CheckUserHasSingleRole")
}

// CheckUserHasAnyRole implements [UserRoleRepositoryInterface].
func (userRoleRepository *UserRoleRepository) CheckUserHasAnyRole() {
	userRoleRepository.logger.Info("userRoleRepository->CheckUserHasAnyRole")
}

// CheckUserHasAllRole implements [UserRoleRepositoryInterface].
func (userRoleRepository *UserRoleRepository) CheckUserHasAllRole() {
	userRoleRepository.logger.Info("userRoleRepository->CheckUserHasAllRole")
}

func NewUserRoleRepository(db *sql.DB,logger *zap.Logger,serverConfig *config.ServerConfig,) UserRoleRepositoryInterface {
	userRoleRepository := &UserRoleRepository{
		db:           db,
		logger:       logger,
		serverConfig: serverConfig,
	}

	return userRoleRepository
}
