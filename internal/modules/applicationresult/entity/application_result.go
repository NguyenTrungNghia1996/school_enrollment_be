package entity

import (
	"time"
	"go_be_enrollment/internal/modules/application/entity"
)

type ApplicationResult struct {
	ID              uint                `gorm:"primaryKey"`
	ApplicationID   uint                `gorm:"not null;uniqueIndex"`
	TotalScore      float64             `gorm:"type:decimal(5,2);default:0.00"`
	PriorityScore   float64             `gorm:"type:decimal(5,2);default:0.00"`
	AdditionalScore float64             `gorm:"type:decimal(5,2);default:0.00"`
	FinalTotalScore float64             `gorm:"type:decimal(5,2);default:0.00"`
	Ranking         *int                
	ResultStatus    string              `gorm:"type:enum('Pending','Passed','Failed','Waitlisted');default:'Pending'"`
	Notes           string              `gorm:"type:text"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	Application     *entity.Application `gorm:"foreignKey:ApplicationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
