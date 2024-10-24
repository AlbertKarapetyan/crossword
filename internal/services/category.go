// Package services contains the business logic implementations.
// It interacts with repository layers to perform operations on category data.
package services

import (
	"crossword/internal/database/models"             // Importing model
	"crossword/internal/infrastructure/repositories" // Importing repository interfaces
	service "crossword/internal/services/interfaces" // Importing service interface definitions
)

// categoryService is a struct that holds a reference to the CategoryRepositoryI interface,
// which is used to interact with the category data in the database.
type categoryService struct {
	repo repositories.CategoryRepositoryI // Repository for category operations
}

// NewCategoryService initializes and returns an instance of categoryService,
// implementing the CategoryServiceI interface.
func NewCategoryService(repo repositories.CategoryRepositoryI) service.CategoryServiceI {
	return &categoryService{repo: repo} // Dependency injection of repository
}

// Create implements services.CategoryServiceI.
func (c *categoryService) Create(category *models.Category) (*models.Category, error) {
	return c.repo.Create(category)
}

// Delete implements services.CategoryServiceI.
func (c *categoryService) Delete(id uint) error {
	return c.repo.Delete(id)
}

// Get implements services.CategoryServiceI.
func (c *categoryService) Get(id uint) (*models.Category, error) {
	return c.repo.Get(id)
}

// Update implements services.CategoryServiceI.
func (c *categoryService) Update(category *models.Category) (*models.Category, error) {
	return c.repo.Update(category)
}
