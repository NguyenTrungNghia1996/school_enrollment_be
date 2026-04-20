package entity

import "time"

type Examiner struct {
	ID               uint      `gorm:"primaryKey"`
	FullName         string    `gorm:"size:255;not null"`
	OrganizationName string    `gorm:"size:255"`
	PhoneNumber      string    `gorm:"size:50"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
