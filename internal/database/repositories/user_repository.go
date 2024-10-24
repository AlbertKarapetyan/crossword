// Package repositories provides concrete implementations of data access methods
// defined in the interfaces. This implementation interacts with the database using GORM.
package repositories

import (
	model "crossword/internal/database/models"            // Importing user model
	repo "crossword/internal/infrastructure/repositories" // Importing repository interface

	"gorm.io/gorm" // Importing GORM, an ORM library for Golang
)

// userRepository is a struct that holds the GORM DB connection,
// used to perform database operations on user records.
type userRepository struct {
	db *gorm.DB // Database connection
}

// NewUserRepository initializes and returns an instance of userRepository
// which implements the UserRepositoryI interface.
func NewUserRepository(db *gorm.DB) repo.UserRepositoryI {
	return &userRepository{db: db}
}

// Create inserts a new user record into the database.
// If the operation is successful, it returns the created user object.
// If an error occurs, it returns nil and the error.
func (r *userRepository) Create(user *model.User) (*model.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err // Return nil and error if the creation fails
	}
	return user, nil // Return the created user if successful
}

// Delete removes a user record from the database by its ID.
// It returns an error if the deletion operation fails.
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error // Perform deletion by ID
}

// Get retrieves a user from the database by its unique ID.
// It returns the user object if found or an error if retrieval fails.
func (r *userRepository) Get(id uint) (*model.User, error) {
	var user model.User                // Create a variable to hold the user data
	err := r.db.First(&user, id).Error // Query the database for the user by ID
	return &user, err                  // Return the user and any error encountered
}

// Update modifies an existing user record in the database.
// It returns the updated user object or an error if the update fails.
func (r *userRepository) Update(user *model.User) (*model.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err // Return nil and error if the update fails
	}
	return user, nil // Return the updated user if successful
}
