package repository

import (
	"go_be_enrollment/internal/modules/adminauth/entity"
	"gorm.io/gorm"
)

type AdminUserRepository interface {
	FindByUsername(username string) (*entity.AdminUser, error)
	FindByID(id uint) (*entity.AdminUser, error)
}

type adminUserRepository struct {
	db *gorm.DB
}

func NewAdminUserRepository(db *gorm.DB) AdminUserRepository {
	return &adminUserRepository{db: db}
}

func (r *adminUserRepository) FindByUsername(username string) (*entity.AdminUser, error) {
	var admin entity.AdminUser
	err := r.db.Where("username = ?", username).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminUserRepository) FindByID(id uint) (*entity.AdminUser, error) {
	var admin entity.AdminUser
	err := r.db.First(&admin, id).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
