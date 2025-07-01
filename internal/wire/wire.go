//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/Adcols/gin-server-api/config"
	repo "github.com/Adcols/gin-server-api/internal/repo"
	service "github.com/Adcols/gin-server-api/internal/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProvideDB 提供数据库实例
func ProvideDB() *gorm.DB {
	return config.DB
}

// ProvideUserRepo 提供用户仓库实例
func ProvideUserRepo(db *gorm.DB) *repo.UserRepo {
	return repo.NewUserRepo(db)
}

// ProvideUserService 提供用户服务实例
func ProvideUserService(userRepo *repo.UserRepo) *service.UserService {
	return &service.UserService{UserRepo: userRepo}
}

// InitializeUserService 初始化用户服务
func InitializeUserService() *service.UserService {
	wire.Build(
		ProvideDB,
		ProvideUserRepo,
		ProvideUserService,
	)
	return nil
}
