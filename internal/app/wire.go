//go:build wireinject
// +build wireinject

package app

import (
	"crossword/internal/database"
	repo "crossword/internal/database/repositories"
	"crossword/internal/services"
	serviceI "crossword/internal/services/interfaces"

	"github.com/google/wire"
)

// InitializeUserService sets up the dependency injection for UserServiceI.
func InitUserService() (serviceI.UserServiceI, error) {
	wire.Build(
		database.ProvideDB, // Provide *gorm.DB
		repo.NewUserRepository,
		services.NewUserService,
	)
	return nil, nil
}

// InitializeCategoryService sets up the dependency injection for CategoryServiceI.
func InitCategoryService() (serviceI.CategoryServiceI, error) {
	wire.Build(
		database.ProvideDB, // Provide *gorm.DB
		repo.NewCategoryRepository,
		services.NewCategoryService,
	)
	return nil, nil
}
