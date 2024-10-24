package services

import (
	"crossword/internal/database/models"
)

type CategoryServiceI interface {
	Create(category *models.Category) (*models.Category, error)
	Get(id uint) (*models.Category, error)
	Update(category *models.Category) (*models.Category, error)
	Delete(id uint) error
}
