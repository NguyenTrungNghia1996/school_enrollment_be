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
	GetAssignedRoleGroups(adminID uint) ([]entity.RoleGroup, error)
	ReplaceRoleGroups(adminID uint, roleGroupIDs []uint) error
	CheckRoleGroupsExist(roleGroupIDs []uint) (bool, error)
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

func (r *adminUserRepository) GetAssignedRoleGroups(adminID uint) ([]entity.RoleGroup, error) {
	var roles []entity.RoleGroup
	err := r.db.Joins("JOIN admin_user_role_groups ON admin_user_role_groups.role_group_id = role_groups.id").
		Where("admin_user_role_groups.admin_user_id = ?", adminID).
		Find(&roles).Error
	return roles, err
}

func (r *adminUserRepository) ReplaceRoleGroups(adminID uint, roleGroupIDs []uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Xóa các nhóm quyền cũ
		if err := tx.Where("admin_user_id = ?", adminID).Delete(&entity.AdminUserRoleGroup{}).Error; err != nil {
			return err
		}

		// Thêm mới
		if len(roleGroupIDs) > 0 {
			var insertData []entity.AdminUserRoleGroup
			for _, rid := range roleGroupIDs {
				insertData = append(insertData, entity.AdminUserRoleGroup{
					AdminUserID: adminID,
					RoleGroupID: rid,
				})
			}
			if err := tx.Create(&insertData).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *adminUserRepository) CheckRoleGroupsExist(roleGroupIDs []uint) (bool, error) {
	if len(roleGroupIDs) == 0 {
		return true, nil
	}
	var count int64
	err := r.db.Model(&entity.RoleGroup{}).Where("id IN ?", roleGroupIDs).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == int64(len(roleGroupIDs)), nil
}
