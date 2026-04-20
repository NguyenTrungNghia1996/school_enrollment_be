package entity

import (
	"time"
	"go_be_enrollment/internal/modules/application/entity"
	subject_entity "go_be_enrollment/internal/modules/subject/entity"
)

type ApplicationExamScore struct {
	ID            uint                        `gorm:"primaryKey"`
	ApplicationID uint                        `gorm:"not null;uniqueIndex:idx_app_subject_score"`
	SubjectID     uint                        `gorm:"not null;uniqueIndex:idx_app_subject_score"`
	RawScore      *float64                    `gorm:"type:decimal(5,2)"`
	BonusScore    float64                     `gorm:"type:decimal(5,2);default:0.00"`
	FinalScore    *float64                    `gorm:"type:decimal(5,2)"`
	IsAbsent      bool                        `gorm:"default:false"`
	Notes         string                      `gorm:"type:text"`
	CreatedAt     time.Time
	UpdatedAt     time.Time

	Application   *entity.Application         `gorm:"foreignKey:ApplicationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Subject       *subject_entity.Subject     `gorm:"foreignKey:SubjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
