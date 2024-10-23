package repositories

import (
	model "crossword/internal/database/models"
)

type CategoryRepositoryI interface {
	Create(category *model.Category) error
	Get(id uint) (*model.Category, error)
	Update(category *model.Category) error
	Delete(id uint) error
}
