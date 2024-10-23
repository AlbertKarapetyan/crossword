package models

import (
	"gorm.io/gorm"
)

// Word represents a word in the database.
type Word struct {
	ID              int            `gorm:"primaryKey;column:id;autoIncrement"`
	CategoryID      int16          `gorm:"column:category_id;not null;index:ix_words_category_id"`
	Lang            string         `gorm:"column:lang;size:2;not null;index:ix_words_lang"`
	Answer          string         `gorm:"column:answer;size:30;not null"`
	Question        string         `gorm:"column:question;size:250;not null"`
	Len             int16          `gorm:"column:len;not null"`
	WhiteSpaceCount int16          `gorm:"column:white_space_count;not null"`
	Difficulty      int16          `gorm:"column:difficulty;not null"`
	AvailableLevel  *int16         `gorm:"column:available_level"`
	About           string         `gorm:"column:about;size:250;not null"`
	IsActive        bool           `gorm:"column:is_active;not null"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

// TableName sets the explicit table name for the Word model.
func (Word) TableName() string {
	return "words"
}
