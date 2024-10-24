// Package services contains the business logic implementations.
// It interacts with repository layers to perform operations on word data.
package services

import (
	"crossword/internal/database/models"             // Importing word model
	"crossword/internal/infrastructure/repositories" // Importing repository interfaces
	service "crossword/internal/services/interfaces" // Importing service interface definitions
)

// wordService is a struct that holds a reference to the WordRepositoryI interface,
// which is used to interact with the word data in the database.
type wordService struct {
	repo repositories.WordRepositoryI // Repository for word operations
}

// NewWordService initializes and returns an instance of wordService,
// implementing the WordServiceI interface.
func NewWordService(repo repositories.WordRepositoryI) service.WordServiceI {
	return &wordService{repo: repo} // Dependency injection of repository
}

// Create implements services.WordServiceI.
func (w *wordService) Create(word *models.Word) (*models.Word, error) {
	return w.repo.Create(word)
}

// Delete implements services.WordServiceI.
func (w *wordService) Delete(id uint) error {
	return w.repo.Delete(id)
}

// Get implements services.WordServiceI.
func (w *wordService) Get(id uint) (*models.Word, error) {
	return w.repo.Get(id)
}

// Update implements services.WordServiceI.
func (w *wordService) Update(word *models.Word) (*models.Word, error) {
	return w.repo.Update(word)
}
