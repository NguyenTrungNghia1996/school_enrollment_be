package repository

import (
	"go_be_enrollment/internal/modules/wardunit/dto"
	"go_be_enrollment/internal/modules/wardunit/entity"
	province_entity "go_be_enrollment/internal/modules/province/entity"

	"gorm.io/gorm"
)

type WardUnitRepository interface {
	GetList(filter *dto.WardUnitFilter) ([]entity.WardUnit, int64, error)
	FindByID(id uint) (*entity.WardUnit, error)
	Create(w *entity.WardUnit) error
	Update(w *entity.WardUnit) error
	CheckNameExistsInProvince(provinceID uint, name string, excludeID uint) bool
	CheckProvinceExists(id uint) bool
	GetActiveListByProvince(provinceID uint) ([]entity.WardUnit, error)
}

type wardUnitRepository struct {
	db *gorm.DB
}

func NewWardUnitRepository(db *gorm.DB) WardUnitRepository {
	return &wardUnitRepository{db: db}
}

func (r *wardUnitRepository) GetList(filter *dto.WardUnitFilter) ([]entity.WardUnit, int64, error) {
	query := r.db.Model(&entity.WardUnit{})

	if filter.Keyword != "" {
		query = query.Where("code LIKE ? OR name LIKE ?", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")
	}

	if filter.IsActive != nil {
		query = query.Where("is_active = ?", *filter.IsActive)
	}

	if filter.ProvinceID != nil {
		query = query.Where("province_id = ?", *filter.ProvinceID)
	}

	if filter.UnitType != nil && *filter.UnitType != "" {
		query = query.Where("unit_type = ?", *filter.UnitType)
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

	var list []entity.WardUnit
	if err := query.Order("id desc").Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (r *wardUnitRepository) FindByID(id uint) (*entity.WardUnit, error) {
	var item entity.WardUnit
	if err := r.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *wardUnitRepository) Create(w *entity.WardUnit) error {
	return r.db.Create(w).Error
}

func (r *wardUnitRepository) Update(w *entity.WardUnit) error {
	return r.db.Save(w).Error
}

func (r *wardUnitRepository) CheckNameExistsInProvince(provinceID uint, name string, excludeID uint) bool {
	var count int64
	r.db.Model(&entity.WardUnit{}).Where("province_id = ? AND name = ? AND id != ?", provinceID, name, excludeID).Count(&count)
	return count > 0
}

func (r *wardUnitRepository) CheckProvinceExists(id uint) bool {
	var count int64
	r.db.Model(&province_entity.Province{}).Where("id = ?", id).Count(&count)
	return count > 0
}

func (r *wardUnitRepository) GetActiveListByProvince(provinceID uint) ([]entity.WardUnit, error) {
	var list []entity.WardUnit
	if err := r.db.Where("province_id = ? AND is_active = ?", provinceID, true).Order("name ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
