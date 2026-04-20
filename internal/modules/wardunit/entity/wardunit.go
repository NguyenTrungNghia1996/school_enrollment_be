package entity

import "time"

type WardUnit struct {
	ID         uint      `gorm:"primaryKey"`
	ProvinceID uint      `gorm:"index;not null;uniqueIndex:idx_province_name"`
	Code       string    `gorm:"size:50;not null"`
	Name       string    `gorm:"size:255;not null;uniqueIndex:idx_province_name"`
	UnitType   string    `gorm:"type:ENUM('Ward','Commune','SpecialZone');not null"`
	IsActive   bool      `gorm:"default:true"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
