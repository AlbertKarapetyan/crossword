// Package repositories provides concrete implementations of data access methods
// defined in the interfaces. This implementation interacts with the database using GORM.
package repositories

// Importing necessary packages
import (
	// Importing models from the internal database models package
	model "crossword/internal/database/models"
	// Importing repository interfaces from the internal infrastructure repositories package
	repo "crossword/internal/infrastructure/repositories"

	// Importing GORM, an ORM library for Golang
	"gorm.io/gorm"
)

// categoryRepository struct implements the repo.CategoryRepositoryI interface.
// It holds a pointer to the GORM database connection.
type categoryRepository struct {
	db *gorm.DB // Database connection
}

// NewCategoryRepository is a constructor function that returns
// a new instance of the categoryRepository, implementing the CategoryRepositoryI interface.
func NewCategoryRepository(db *gorm.DB) repo.CategoryRepositoryI {
	return &categoryRepository{db: db}
}

// Create inserts a new Category record into the database.
// It returns the created Category or an error if the operation fails.
func (c *categoryRepository) Create(category *model.Category) (*model.Category, error) {
	// Using GORM's Create method to insert the category into the database
	if err := c.db.Create(category).Error; err != nil {
		return nil, err // Return nil and the error if creation fails
	}
	return category, nil // Return the created category if successful
}

// Delete removes a Category record from the database by its ID.
// It returns an error if the deletion fails.
func (c *categoryRepository) Delete(id uint) error {
	// Using GORM's Delete method to remove the category by ID
	return c.db.Delete(&model.Category{}, id).Error
}

// Get retrieves a Category record by its ID from the database.
// It returns the found Category or an error if retrieval fails.
func (c *categoryRepository) Get(id uint) (*model.Category, error) {
	var category model.Category // Variable to store the retrieved category
	// Using GORM's First method to find the category by ID
	err := c.db.First(&category, id).Error
	return &category, err // Return the category and any error encountered
}

// Update modifies an existing Category record in the database.
// It returns the updated Category or an error if the operation fails.
func (c *categoryRepository) Update(category *model.Category) (*model.Category, error) {
	// Using GORM's Save method to update the category
	if err := c.db.Save(category).Error; err != nil {
		return nil, err // Return nil and the error if the update fails
	}
	return category, nil // Return the updated category if successful
}
