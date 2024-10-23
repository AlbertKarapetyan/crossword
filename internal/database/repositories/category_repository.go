package repositories

import (
	"crossword/internal/database/models"
	repo "crossword/internal/infrastructure/repositories"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) repo.CategoryRepositoryI {
	return &categoryRepository{db: db}
}

// Create implements repositories.CategoryRepositoryI.
func (c *categoryRepository) Create(category *models.Category) error {
	panic("unimplemented")
}

// Delete implements repositories.CategoryRepositoryI.
func (c *categoryRepository) Delete(id uint) error {
	panic("unimplemented")
}

// Get implements repositories.CategoryRepositoryI.
func (c *categoryRepository) Get(id uint) (*models.Category, error) {
	panic("unimplemented")
}

// Update implements repositories.CategoryRepositoryI.
func (c *categoryRepository) Update(category *models.Category) error {
	panic("unimplemented")
}
