package services

import (
	"crossword/internal/database/models"
	"crossword/internal/infrastructure/repositories"
	service "crossword/internal/services/interfaces"
)

type categoryService struct {
	repo repositories.CategoryRepositoryI
}

func NewCategoryService(repo repositories.CategoryRepositoryI) service.CategoryServiceI {
	return &categoryService{repo: repo}
}

// Create implements services.CategoryServiceI.
func (c *categoryService) Create(category *models.Category) error {
	return c.repo.Create(category)
}

// Delete implements services.CategoryServiceI.
func (c *categoryService) Delete(id uint) error {
	return c.repo.Delete(id)
}

// Get implements services.CategoryServiceI.
func (c *categoryService) Get(id uint) (*models.Category, error) {
	panic("unimplemented")
}

// Update implements services.CategoryServiceI.
func (c *categoryService) Update(category *models.Category) error {
	panic("unimplemented")
}
