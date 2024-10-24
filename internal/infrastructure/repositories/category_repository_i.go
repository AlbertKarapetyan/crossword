// Package repositories defines interfaces for interacting with the data layer.
package repositories

import (
	"crossword/internal/database/models"
)

// CategoryRepositoryI defines the contract for category repository operations.
type CategoryRepositoryI interface {
	// Create inserts a new category record in the database.
	// Returns the created category along with an error, if any.
	Create(category *models.Category) (*models.Category, error)

	// Get retrieves a category from the database by its unique ID.
	// Returns the category found or an error if the category does not exist or retrieval fails.
	Get(id uint) (*models.Category, error)

	// Update modifies an existing category record in the database.
	// Returns the updated category object or an error in case of failure.
	Update(category *models.Category) (*models.Category, error)

	// Delete removes a category from the database by ID.
	// Returns an error if the deletion fails.
	Delete(id uint) error
}
