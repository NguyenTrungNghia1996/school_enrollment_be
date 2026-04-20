package repository

import (
	"go_be_enrollment/internal/modules/examroomassignment/dto"
	"go_be_enrollment/internal/modules/examroomassignment/entity"

	"gorm.io/gorm"
)

type ExamRoomAssignmentRepository interface {
	GetList(filter *dto.ExamRoomAssignmentFilter) ([]entity.ExamRoomAssignment, int64, error)
	FindByID(id uint) (*entity.ExamRoomAssignment, error)
	FindByApplicationID(appID uint) (*entity.ExamRoomAssignment, error)
	FindByRoomAndSeat(roomID uint, seat string) (*entity.ExamRoomAssignment, error)
	CountByRoomID(roomID uint) (int64, error)
	Create(e *entity.ExamRoomAssignment) error
	Update(e *entity.ExamRoomAssignment) error
	Delete(id uint) error
	GetListByRoomID(roomID uint) ([]entity.ExamRoomAssignment, error)
}

type examRoomAssignmentRepository struct {
	db *gorm.DB
}

func NewExamRoomAssignmentRepository(db *gorm.DB) ExamRoomAssignmentRepository {
	return &examRoomAssignmentRepository{db: db}
}

func (r *examRoomAssignmentRepository) GetList(filter *dto.ExamRoomAssignmentFilter) ([]entity.ExamRoomAssignment, int64, error) {
	query := r.db.Model(&entity.ExamRoomAssignment{}).Joins("Application").Joins("ExamRoom")

	if filter.Keyword != "" {
		key := "%" + filter.Keyword + "%"
		query = query.Where("Application.candidate_full_name LIKE ? OR Application.national_id LIKE ? OR exam_room_assignments.seat_number LIKE ?", key, key, key)
	}

	if filter.AdmissionPeriodID != nil && *filter.AdmissionPeriodID > 0 {
		query = query.Where("Application.admission_period_id = ?", *filter.AdmissionPeriodID)
	}

	if filter.ExamRoomID != nil && *filter.ExamRoomID > 0 {
		query = query.Where("exam_room_assignments.exam_room_id = ?", *filter.ExamRoomID)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	page := filter.Page
	if page <= 0 {
		page = 1
	}
	limit := filter.Limit
	switch {
	case limit > 100:
		limit = 100
	case limit <= 0:
		limit = 10
	}
	offset := (page - 1) * limit

	var assignments []entity.ExamRoomAssignment
	if err := query.Preload("Application").Preload("ExamRoom").Order("exam_room_assignments.id desc").Offset(offset).Limit(limit).Find(&assignments).Error; err != nil {
		return nil, 0, err
	}

	return assignments, total, nil
}

func (r *examRoomAssignmentRepository) FindByID(id uint) (*entity.ExamRoomAssignment, error) {
	var a entity.ExamRoomAssignment
	if err := r.db.Preload("Application").Preload("ExamRoom").First(&a, id).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *examRoomAssignmentRepository) FindByApplicationID(appID uint) (*entity.ExamRoomAssignment, error) {
	var a entity.ExamRoomAssignment
	if err := r.db.Where("application_id = ?", appID).First(&a).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}

func (r *examRoomAssignmentRepository) FindByRoomAndSeat(roomID uint, seat string) (*entity.ExamRoomAssignment, error) {
	var a entity.ExamRoomAssignment
	if err := r.db.Where("exam_room_id = ? AND seat_number = ?", roomID, seat).First(&a).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}

func (r *examRoomAssignmentRepository) CountByRoomID(roomID uint) (int64, error) {
	var count int64
	if err := r.db.Model(&entity.ExamRoomAssignment{}).Where("exam_room_id = ?", roomID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *examRoomAssignmentRepository) Create(e *entity.ExamRoomAssignment) error {
	return r.db.Create(e).Error
}

func (r *examRoomAssignmentRepository) Update(e *entity.ExamRoomAssignment) error {
	return r.db.Save(e).Error
}

func (r *examRoomAssignmentRepository) Delete(id uint) error {
	return r.db.Delete(&entity.ExamRoomAssignment{}, id).Error
}

func (r *examRoomAssignmentRepository) GetListByRoomID(roomID uint) ([]entity.ExamRoomAssignment, error) {
	var assignments []entity.ExamRoomAssignment
	if err := r.db.Preload("Application").Preload("ExamRoom").Where("exam_room_id = ?", roomID).Find(&assignments).Error; err != nil {
		return nil, err
	}
	return assignments, nil
}
