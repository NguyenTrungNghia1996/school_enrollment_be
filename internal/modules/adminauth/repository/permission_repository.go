package repository

import (
	"go_be_enrollment/internal/modules/adminauth/entity"
	"gorm.io/gorm"
)

type PermissionRepository interface {
	GetRoleGroupsByAdminID(adminID uint) ([]entity.RoleGroup, error)
	GetPermissionsByAdminID(adminID uint) ([]entity.RoleGroupPermission, error)
	GetAllMenus() ([]entity.Menu, error)
}

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{db: db}
}

func (r *permissionRepository) GetRoleGroupsByAdminID(adminID uint) ([]entity.RoleGroup, error) {
	var roles []entity.RoleGroup
	err := r.db.Joins("JOIN admin_user_role_groups ON admin_user_role_groups.role_group_id = role_groups.id").
		Where("admin_user_role_groups.admin_user_id = ?", adminID).
		Find(&roles).Error
	return roles, err
}

func (r *permissionRepository) GetPermissionsByAdminID(adminID uint) ([]entity.RoleGroupPermission, error) {
	var permissions []entity.RoleGroupPermission
	err := r.db.Joins("JOIN admin_user_role_groups ON admin_user_role_groups.role_group_id = role_group_permissions.role_group_id").
		Where("admin_user_role_groups.admin_user_id = ?", adminID).
		Find(&permissions).Error
	return permissions, err
}

func (r *permissionRepository) GetAllMenus() ([]entity.Menu, error) {
	var menus []entity.Menu
	err := r.db.Order("sort_order asc").Find(&menus).Error
	return menus, err
}
