package repositories

import (
	"database/sql"

	"github.com/bhandehemant2004-debug/JWT-AUTH-GOLANG/internal/config"
	"go.uber.org/zap"
)

type UserRepositoryInterface interface {
	CreateUser()
	GetAllUser()
	GetUserById()
	UpdateUserById()
	DeleteUserById()
	GetUserByEmail()
	GetUserByUsernameAndEmail()
}

type UserRepository struct {
	db           *sql.DB
	logger       *zap.Logger
	serverConfig *config.ServerConfig
}

// CreateUser implements [UserRepositoryInterface].
func (userRepository *UserRepository) CreateUser() {
	userRepository.logger.Info("userRepository->Createuser")
}

// DeleteUserById implements [UserRepositoryInterface].
func (userRepository *UserRepository) DeleteUserById() {
	userRepository.logger.Info("userRepository->deleteuserbyid")
}

// GetAllUser implements [UserRepositoryInterface].
func (userRepository *UserRepository) GetAllUser() {
	userRepository.logger.Info("userRepository->getalluser")
}

// GetUserByEmail implements [UserRepositoryInterface].
func (userRepository *UserRepository) GetUserByEmail() {
	userRepository.logger.Info("userRepository->getuserbyemail")
}

// GetUserById implements [UserRepositoryInterface].
func (userRepository *UserRepository) GetUserById() {
	userRepository.logger.Info("userRepository->getuserbyid")
}

// GetUserByUsernameAndEmail implements [UserRepositoryInterface].
func (userRepository *UserRepository) GetUserByUsernameAndEmail() {
	userRepository.logger.Info("userRepository->getuserbyusernameandemail")
}

// UpdateUserById implements [UserRepositoryInterface].
func (userRepository *UserRepository) UpdateUserById() {
	userRepository.logger.Info("userRepository->updatebyid")
}

func NewUserRepository(db *sql.DB , logger *zap.Logger,ServerConfig *config.ServerConfig)UserRepositoryInterface{
	UserRepository:= &UserRepository{
		db: db,
		logger: logger,
		serverConfig: ServerConfig,
	}
	return UserRepository
}