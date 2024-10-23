package services

import (
	"crossword/internal/database/models"
	"crossword/internal/infrastructure/repositories"
	service "crossword/internal/services/interfaces"
)

type userService struct {
	repo repositories.UserRepositoryI
}

func NewUserService(repo repositories.UserRepositoryI) service.UserServiceI {
	return &userService{repo: repo}
}

// CreateUser implements services.UserServiceI.
func (u *userService) Create(user *models.User) error {
	return u.repo.Create(user)
}

// DeleteUser implements services.UserServiceI.
func (u *userService) Delete(id uint) error {
	return u.repo.Delete(id)
}

// GetUserByID implements services.UserServiceI.
func (u *userService) Get(id uint) (*models.User, error) {
	panic("unimplemented")
}

// UpdateUser implements services.UserServiceI.
func (u *userService) Update(user *models.User) error {
	panic("unimplemented")
}
