package entity

import "time"

type AdminUser struct {
	ID           uint    `gorm:"primaryKey"`
	Username     string  `gorm:"uniqueIndex;not null"`
	PasswordHash string  `gorm:"not null"`
	FullName     string  `gorm:"not null"`
	Email        *string `gorm:"uniqueIndex"`
	PhoneNumber  *string
	IsSuperAdmin bool    `gorm:"default:false"`
	IsActive     bool    `gorm:"default:true"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (AdminUser) TableName() string {
	return "admin_users"
}
