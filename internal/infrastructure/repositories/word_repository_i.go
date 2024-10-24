// Package repositories defines interfaces for interacting with the data layer.
package repositories

import (
	"crossword/internal/database/models"
)

// WordRepositoryI defines the contract for word repository operations.
type WordRepositoryI interface {
	// Create inserts a new word record in the database.
	// Returns the created word along with an error, if any.
	Create(word *models.Word) (*models.Word, error)

	// Get retrieves a word from the database by its unique ID.
	// Returns the word found or an error if the word does not exist or retrieval fails.
	Get(id uint) (*models.Word, error)

	// Update modifies an existing word record in the database.
	// Returns the updated word object or an error in case of failure.
	Update(word *models.Word) (*models.Word, error)

	// Delete removes a word from the database by ID.
	// Returns an error if the deletion fails.
	Delete(id uint) error
}
