package repositories

import (
	model "crossword/internal/database/models"
	repo "crossword/internal/infrastructure/repositories"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repo.UserRepositoryI {
	return &userRepository{db: db}
}

// Create implements UserRepositoryI.
func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// Delete implements UserRepositoryI.
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

// Get implements UserRepositoryI.
func (r *userRepository) Get(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return &user, err
}

// Update implements UserRepositoryI.
func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}
