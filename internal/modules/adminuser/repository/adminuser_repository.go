package repository

import (
	"strings"

	"go_be_enrollment/internal/modules/adminauth/entity"
	"go_be_enrollment/internal/modules/adminuser/dto"

	"gorm.io/gorm"
)

type AdminUserRepository interface {
	FindAll(filter *dto.AdminUserFilter) ([]entity.AdminUser, int64, error)
	FindByID(id uint) (*entity.AdminUser, error)
	Create(admin *entity.AdminUser) error
	Update(admin *entity.AdminUser) error
	CheckUsernameExists(username string, excludeID uint) bool
	CheckEmailExists(email string, excludeID uint) bool
}

type adminUserRepository struct {
	db *gorm.DB
}

func NewAdminUserRepository(db *gorm.DB) AdminUserRepository {
	return &adminUserRepository{db: db}
}

func (r *adminUserRepository) FindAll(filter *dto.AdminUserFilter) ([]entity.AdminUser, int64, error) {
	var admins []entity.AdminUser
	var total int64

	query := r.db.Model(&entity.AdminUser{})

	if filter.Keyword != "" {
		kw := "%" + strings.ToLower(filter.Keyword) + "%"
		query = query.Where("LOWER(username) LIKE ? OR LOWER(full_name) LIKE ? OR LOWER(email) LIKE ?", kw, kw, kw)
	}

	if filter.IsActive != nil {
		query = query.Where("is_active = ?", *filter.IsActive)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	page := filter.Page
	if page < 1 {
		page = 1
	}
	limit := filter.Limit
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	err = query.Order("id desc").Offset(offset).Limit(limit).Find(&admins).Error
	return admins, total, err
}

func (r *adminUserRepository) FindByID(id uint) (*entity.AdminUser, error) {
	var admin entity.AdminUser
	err := r.db.First(&admin, id).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminUserRepository) Create(admin *entity.AdminUser) error {
	return r.db.Create(admin).Error
}

func (r *adminUserRepository) Update(admin *entity.AdminUser) error {
	return r.db.Save(admin).Error
}

func (r *adminUserRepository) CheckUsernameExists(username string, excludeID uint) bool {
	var count int64
	r.db.Model(&entity.AdminUser{}).Where("username = ? AND id != ?", username, excludeID).Count(&count)
	return count > 0
}

func (r *adminUserRepository) CheckEmailExists(email string, excludeID uint) bool {
	var count int64
	r.db.Model(&entity.AdminUser{}).Where("email = ? AND email != '' AND id != ?", email, excludeID).Count(&count)
	return count > 0
}
