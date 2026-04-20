package service

import (
	"go_be_enrollment/internal/modules/adminauth/entity"
	"go_be_enrollment/internal/modules/adminauth/repository"
)

type PermissionService interface {
	GetMergedPermissions(adminID uint) (map[string]int64, error)
	CheckPermission(adminID uint, isSuperAdmin bool, permissionKey string, permissionBit int) (bool, error)
	GetAllowedMenus(adminID uint, isSuperAdmin bool) ([]entity.Menu, error)
}

type permissionService struct {
	repo repository.PermissionRepository
}

func NewPermissionService(repo repository.PermissionRepository) PermissionService {
	return &permissionService{repo: repo}
}

func (s *permissionService) GetMergedPermissions(adminID uint) (map[string]int64, error) {
	permissions, err := s.repo.GetPermissionsByAdminID(adminID)
	if err != nil {
		return nil, err
	}

	merged := make(map[string]int64)
	for _, p := range permissions {
		merged[p.PermissionKey] |= p.PermissionValue
	}

	return merged, nil
}

func (s *permissionService) CheckPermission(adminID uint, isSuperAdmin bool, permissionKey string, permissionBit int) (bool, error) {
	if isSuperAdmin {
		return true, nil
	}

	merged, err := s.GetMergedPermissions(adminID)
	if err != nil {
		return false, err
	}

	val, ok := merged[permissionKey]
	if !ok {
		return false, nil
	}

	return CheckBitmask(val, permissionBit), nil
}

func (s *permissionService) GetAllowedMenus(adminID uint, isSuperAdmin bool) ([]entity.Menu, error) {
	allMenus, err := s.repo.GetAllMenus()
	if err != nil {
		return nil, err
	}

	if isSuperAdmin {
		return allMenus, nil
	}

	merged, err := s.GetMergedPermissions(adminID)
	if err != nil {
		return nil, err
	}

	var allowed []entity.Menu
	for _, m := range allMenus {
		if m.MenuKey == "" {
			// Menus without key are considered public/always allowed
			allowed = append(allowed, m)
			continue
		}

		val, ok := merged[m.MenuKey]
		if ok && CheckBitmask(val, m.PermissionBit) {
			allowed = append(allowed, m)
		}
	}

	return allowed, nil
}

// CheckBitmask is a helper to check pure bitmask matching
func CheckBitmask(mergedValue int64, targetBit int) bool {
	return (mergedValue & (1 << targetBit)) != 0
}
