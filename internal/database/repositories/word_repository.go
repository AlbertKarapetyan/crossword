// Package repositories provides concrete implementations of data access methods
// defined in the interfaces. This implementation interacts with the database using GORM.
package repositories

import (
	model "crossword/internal/database/models"            // Importing user model
	repo "crossword/internal/infrastructure/repositories" // Importing repository interface

	"gorm.io/gorm" // Importing GORM, an ORM library for Golang
)

// wordRepository is a struct that holds the GORM DB connection,
// used to perform database operations on word records.
type wordRepository struct {
	db *gorm.DB // Database connection
}

// NewWordRepository initializes and returns an instance of wordRepository
// which implements the WordRepositoryI interface.
func NewWordRepository(db *gorm.DB) repo.WordRepositoryI {
	return &wordRepository{db: db}
}

// Create implements repositories.WordRepositoryI.
func (w *wordRepository) Create(word *model.Word) (*model.Word, error) {
	if err := w.db.Create(word).Error; err != nil {
		return nil, err // Return nil and error if the creation fails
	}
	return word, nil // Return the created word if successful
}

// Delete implements repositories.WordRepositoryI.
func (w *wordRepository) Delete(id uint) error {
	return w.db.Delete(&model.Word{}, id).Error // Perform deletion by ID
}

// Get implements repositories.WordRepositoryI.
func (w *wordRepository) Get(id uint) (*model.Word, error) {
	var word model.Word                // Create a variable to hold the word data
	err := w.db.First(&word, id).Error // Query the database for the word by ID
	return &word, err                  // Return the word and any error encountered
}

// Update implements repositories.WordRepositoryI.
func (w *wordRepository) Update(word *model.Word) (*model.Word, error) {
	if err := w.db.Save(word).Error; err != nil {
		return nil, err // Return nil and error if the update fails
	}
	return word, nil // Return the updated word if successful
}
