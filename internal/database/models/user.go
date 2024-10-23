package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the database.
type User struct {
	ID             uint           `gorm:"primaryKey;autoIncrement;column:id"`
	ExternalID     int64          `gorm:"not null;index:ix_user_external_id;column:external_id"`
	Level          int16          `gorm:"not null;column:level"`
	UserName       string         `gorm:"size:150;column:user_name"`
	Email          string         `gorm:"size:150;column:email"`
	RefID          uint           `gorm:"column:ref_id"`
	IsPremium      bool           `gorm:"not null;column:is_premium"`
	RegionID       int            `gorm:"column:region_id"`
	CreatedAt      time.Time      `gorm:"autoCreateTime;column:created_at;index"`
	LastSignedDate *time.Time     `gorm:"column:last_signed_date"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;index"`
	IsDeleted      bool           `gorm:"default:false;column:is_deleted"`
}

// TableName sets the table name explicitly.
func (User) TableName() string {
	return "users"
}
