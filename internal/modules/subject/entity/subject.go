package entity

import "time"

type Subject struct {
	ID           uint      `gorm:"primaryKey"`
	Code         string    `gorm:"size:50;uniqueIndex;not null"`
	Name         string    `gorm:"size:255;not null"`
	DisplayOrder int       `gorm:"default:0"`
	IsActive     bool      `gorm:"default:true"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
