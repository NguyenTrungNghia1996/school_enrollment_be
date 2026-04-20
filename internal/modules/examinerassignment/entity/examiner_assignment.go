package entity

import (
	"time"
	examiner_entity "go_be_enrollment/internal/modules/examiner/entity"
	exam_room_entity "go_be_enrollment/internal/modules/examroom/entity"
)

type ExaminerAssignment struct {
	ID         uint                               `gorm:"primaryKey"`
	ExaminerID uint                               `gorm:"not null;uniqueIndex:idx_examiner_room_role"`
	ExamRoomID uint                               `gorm:"not null;uniqueIndex:idx_examiner_room_role"`
	Role       string                             `gorm:"type:enum('Primary','Secondary','Backup');not null;uniqueIndex:idx_examiner_room_role"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Examiner   *examiner_entity.Examiner          `gorm:"foreignKey:ExaminerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ExamRoom   *exam_room_entity.ExamRoom         `gorm:"foreignKey:ExamRoomID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
