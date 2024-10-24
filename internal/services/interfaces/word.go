package services

import (
	"crossword/internal/database/models"
)

type WordServiceI interface {
	Create(word *models.Word) (*models.Word, error)
	Get(id uint) (*models.Word, error)
	Update(word *models.Word) (*models.Word, error)
	Delete(id uint) error
}
