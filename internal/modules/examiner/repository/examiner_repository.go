package repository

import (
	"go_be_enrollment/internal/modules/examiner/dto"
	"go_be_enrollment/internal/modules/examiner/entity"

	"gorm.io/gorm"
)

type ExaminerRepository interface {
	GetList(filter *dto.ExaminerFilter) ([]entity.Examiner, int64, error)
	FindByID(id uint) (*entity.Examiner, error)
	Create(e *entity.Examiner) error
	Update(e *entity.Examiner) error
	Delete(id uint) error
	HasRelatedAssignments(id uint) (bool, error)
}

type examinerRepository struct {
	db *gorm.DB
}

func NewExaminerRepository(db *gorm.DB) ExaminerRepository {
	return &examinerRepository{db: db}
}

func (r *examinerRepository) GetList(filter *dto.ExaminerFilter) ([]entity.Examiner, int64, error) {
	query := r.db.Model(&entity.Examiner{})

	if filter.Keyword != "" {
		query = query.Where("full_name LIKE ? OR organization_name LIKE ? OR phone_number LIKE ?", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")
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

	var examiners []entity.Examiner
	if err := query.Order("id desc").Offset(offset).Limit(limit).Find(&examiners).Error; err != nil {
		return nil, 0, err
	}

	return examiners, total, nil
}

func (r *examinerRepository) FindByID(id uint) (*entity.Examiner, error) {
	var examiner entity.Examiner
	if err := r.db.First(&examiner, id).Error; err != nil {
		return nil, err
	}
	return &examiner, nil
}

func (r *examinerRepository) Create(e *entity.Examiner) error {
	return r.db.Create(e).Error
}

func (r *examinerRepository) Update(e *entity.Examiner) error {
	return r.db.Save(e).Error
}

func (r *examinerRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Examiner{}, id).Error
}

func (r *examinerRepository) HasRelatedAssignments(id uint) (bool, error) {
	if r.db.Migrator().HasTable("examiner_assignments") {
		var count int64
		if err := r.db.Table("examiner_assignments").Where("examiner_id = ?", id).Count(&count).Error; err != nil {
			return false, err
		}
		if count > 0 {
			return true, nil
		}
	}
	return false, nil
}
