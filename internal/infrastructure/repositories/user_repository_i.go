// Package repositories defines interfaces for interacting with the data layer.
package repositories

import (
	"crossword/internal/database/models"
)

// UserRepositoryI defines the contract for user repository operations.
type UserRepositoryI interface {
	// Create inserts a new user record in the database.
	// Returns the created user along with an error, if any.
	Create(user *models.User) (*models.User, error)

	// Get retrieves a user from the database by its unique ID.
	// Returns the user found or an error if the user does not exist or retrieval fails.
	Get(id uint) (*models.User, error)

	// Update modifies an existing user record in the database.
	// Returns the updated user object or an error in case of failure.
	Update(user *models.User) (*models.User, error)

	// Delete removes a user from the database by ID.
	// Returns an error if the deletion fails.
	Delete(id uint) error
}
