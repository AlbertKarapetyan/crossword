// Package services contains the business logic implementations.
// It interacts with repository layers to perform operations on user data.
package services

import (
	"crossword/internal/database/models"             // Importing model
	"crossword/internal/infrastructure/repositories" // Importing repository interfaces
	service "crossword/internal/services/interfaces" // Importing service interface definitions
)

// userService is a struct that holds a reference to the UserRepositoryI interface,
// which is used to interact with the user data in the database.
type userService struct {
	repo repositories.UserRepositoryI // Repository for user operations
}

// NewUserService initializes and returns an instance of userService,
// implementing the UserServiceI interface.
func NewUserService(repo repositories.UserRepositoryI) service.UserServiceI {
	return &userService{repo: repo} // Dependency injection of repository
}

// Create implements services.UserServiceI.
func (u *userService) Create(user *models.User) (*models.User, error) {
	return u.repo.Create(user)
}

// Delete implements services.UserServiceI.
func (u *userService) Delete(id uint) error {
	return u.repo.Delete(id)
}

// Get implements services.UserServiceI.
func (u *userService) Get(id uint) (*models.User, error) {
	return u.repo.Get(id)
}

// Update implements services.UserServiceI.
func (u *userService) Update(user *models.User) (*models.User, error) {
	return u.repo.Update(user)
}
