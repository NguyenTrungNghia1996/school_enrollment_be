package repository

import (
	"go_be_enrollment/internal/modules/examinerassignment/dto"
	"go_be_enrollment/internal/modules/examinerassignment/entity"

	"gorm.io/gorm"
)

type ExaminerAssignmentRepository interface {
	GetList(filter *dto.ExaminerAssignmentFilter) ([]entity.ExaminerAssignment, int64, error)
	FindByID(id uint) (*entity.ExaminerAssignment, error)
	FindByUniqueKey(examinerID, roomID uint, role string) (*entity.ExaminerAssignment, error)
	Create(e *entity.ExaminerAssignment) error
	Update(e *entity.ExaminerAssignment) error
	Delete(id uint) error
	GetListByRoomID(roomID uint) ([]entity.ExaminerAssignment, error)
}

type examinerAssignmentRepository struct {
	db *gorm.DB
}

func NewExaminerAssignmentRepository(db *gorm.DB) ExaminerAssignmentRepository {
	return &examinerAssignmentRepository{db: db}
}

func (r *examinerAssignmentRepository) GetList(filter *dto.ExaminerAssignmentFilter) ([]entity.ExaminerAssignment, int64, error) {
	query := r.db.Model(&entity.ExaminerAssignment{}).Joins("Examiner").Joins("ExamRoom")

	if filter.ExaminerID != nil && *filter.ExaminerID > 0 {
		query = query.Where("examiner_assignments.examiner_id = ?", *filter.ExaminerID)
	}

	if filter.ExamRoomID != nil && *filter.ExamRoomID > 0 {
		query = query.Where("examiner_assignments.exam_room_id = ?", *filter.ExamRoomID)
	}

	if filter.Role != "" {
		query = query.Where("examiner_assignments.role = ?", filter.Role)
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

	var list []entity.ExaminerAssignment
	if err := query.Preload("Examiner").Preload("ExamRoom").Order("examiner_assignments.id desc").Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *examinerAssignmentRepository) FindByID(id uint) (*entity.ExaminerAssignment, error) {
	var a entity.ExaminerAssignment
	if err := r.db.Preload("Examiner").Preload("ExamRoom").First(&a, id).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *examinerAssignmentRepository) FindByUniqueKey(examinerID, roomID uint, role string) (*entity.ExaminerAssignment, error) {
	var a entity.ExaminerAssignment
	if err := r.db.Where("examiner_id = ? AND exam_room_id = ? AND role = ?", examinerID, roomID, role).First(&a).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}

func (r *examinerAssignmentRepository) Create(e *entity.ExaminerAssignment) error {
	return r.db.Create(e).Error
}

func (r *examinerAssignmentRepository) Update(e *entity.ExaminerAssignment) error {
	return r.db.Save(e).Error
}

func (r *examinerAssignmentRepository) Delete(id uint) error {
	return r.db.Delete(&entity.ExaminerAssignment{}, id).Error
}

func (r *examinerAssignmentRepository) GetListByRoomID(roomID uint) ([]entity.ExaminerAssignment, error) {
	var list []entity.ExaminerAssignment
	if err := r.db.Preload("Examiner").Preload("ExamRoom").Where("exam_room_id = ?", roomID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
