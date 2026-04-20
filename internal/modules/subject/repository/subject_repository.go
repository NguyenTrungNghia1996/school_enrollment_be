package repository

import (
	"go_be_enrollment/internal/modules/subject/dto"
	"go_be_enrollment/internal/modules/subject/entity"

	"gorm.io/gorm"
)

type SubjectRepository interface {
	GetList(filter *dto.SubjectFilter) ([]entity.Subject, int64, error)
	FindByID(id uint) (*entity.Subject, error)
	Create(s *entity.Subject) error
	Update(s *entity.Subject) error
	CheckCodeExists(code string, excludeID uint) bool
	GetActiveList() ([]entity.Subject, error)
}

type subjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) SubjectRepository {
	return &subjectRepository{db: db}
}

func (r *subjectRepository) GetList(filter *dto.SubjectFilter) ([]entity.Subject, int64, error) {
	query := r.db.Model(&entity.Subject{})

	if filter.Keyword != "" {
		query = query.Where("code LIKE ? OR name LIKE ?", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")
	}

	if filter.IsActive != nil {
		query = query.Where("is_active = ?", *filter.IsActive)
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

	var subjects []entity.Subject
	if err := query.Order("display_order asc, id desc").Offset(offset).Limit(limit).Find(&subjects).Error; err != nil {
		return nil, 0, err
	}

	return subjects, total, nil
}

func (r *subjectRepository) FindByID(id uint) (*entity.Subject, error) {
	var subject entity.Subject
	if err := r.db.First(&subject, id).Error; err != nil {
		return nil, err
	}
	return &subject, nil
}

func (r *subjectRepository) Create(s *entity.Subject) error {
	return r.db.Create(s).Error
}

func (r *subjectRepository) Update(s *entity.Subject) error {
	return r.db.Save(s).Error
}

func (r *subjectRepository) CheckCodeExists(code string, excludeID uint) bool {
	var count int64
	r.db.Model(&entity.Subject{}).Where("code = ? AND id != ?", code, excludeID).Count(&count)
	return count > 0
}

func (r *subjectRepository) GetActiveList() ([]entity.Subject, error) {
	var subjects []entity.Subject
	if err := r.db.Where("is_active = ?", true).Order("display_order asc, id desc").Find(&subjects).Error; err != nil {
		return nil, err
	}
	return subjects, nil
}
