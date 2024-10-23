package services

import (
	model "crossword/internal/database/models"
)

type UserServiceI interface {
	Create(user *model.User) error
	Get(id uint) (*model.User, error)
	Update(user *model.User) error
	Delete(id uint) error
}
