package models

import (
	"gorm.io/gorm"
)

// Category represents a category in the database.
type Category struct {
	ID             int16          `gorm:"primaryKey;column:id;autoIncrement"`
	Name           string         `gorm:"column:name;size:50;not null"`
	Difficulty     int16          `gorm:"column:difficulty;not null"`
	Group          string         `gorm:"column:group;size:50;not null"`
	Description    string         `gorm:"column:description;size:250;not null"`
	AvailableLevel *int16         `gorm:"column:available_level"`
	Lang           string         `gorm:"column:lang;size:2;not null;index:ix_categories_lang"`
	IsActive       bool           `gorm:"column:is_active;not null"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

// TableName sets the table name explicitly.
func (Category) TableName() string {
	return "categories"
}
