package entity

import (
	"time"
	"go_be_enrollment/internal/modules/application/entity"
	exam_room_entity "go_be_enrollment/internal/modules/examroom/entity"
)

type ExamRoomAssignment struct {
	ID            uint                               `gorm:"primaryKey"`
	ApplicationID uint                               `gorm:"not null;uniqueIndex"`
	ExamRoomID    uint                               `gorm:"not null;uniqueIndex:idx_room_seat"`
	SeatNumber    string                             `gorm:"size:50;not null;uniqueIndex:idx_room_seat"`
	CreatedAt     time.Time
	UpdatedAt     time.Time

	// Relations
	Application   *entity.Application                `gorm:"foreignKey:ApplicationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ExamRoom      *exam_room_entity.ExamRoom         `gorm:"foreignKey:ExamRoomID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
