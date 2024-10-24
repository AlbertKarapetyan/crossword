package services

import (
	"crossword/internal/database/models"
)

type UserServiceI interface {
	Create(user *models.User) (*models.User, error)
	Get(id uint) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id uint) error
}
