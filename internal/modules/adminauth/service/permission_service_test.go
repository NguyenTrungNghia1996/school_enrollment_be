package service

import (
	"testing"
	"go_be_enrollment/internal/modules/adminauth/entity"
)

type mockPermissionRepo struct {
	menus []entity.Menu
	perms []entity.RoleGroupPermission
}

func (m *mockPermissionRepo) GetRoleGroupsByAdminID(adminID uint) ([]entity.RoleGroup, error) {
	return nil, nil
}

func (m *mockPermissionRepo) GetPermissionsByAdminID(adminID uint) ([]entity.RoleGroupPermission, error) {
	return m.perms, nil
}

func (m *mockPermissionRepo) GetAllMenus() ([]entity.Menu, error) {
	return m.menus, nil
}

func TestCheckBitmask(t *testing.T) {
	// 5 = 101 in binary -> bit 0 and bit 2 are set
	val := int64(5)
	
	if !CheckBitmask(val, 0) {
		t.Errorf("expected bit 0 to be true")
	}
	if CheckBitmask(val, 1) {
		t.Errorf("expected bit 1 to be false")
	}
	if !CheckBitmask(val, 2) {
		t.Errorf("expected bit 2 to be true")
	}
	if CheckBitmask(val, 3) {
		t.Errorf("expected bit 3 to be false")
	}
}

func TestGetMergedPermissions(t *testing.T) {
	repo := &mockPermissionRepo{
		perms: []entity.RoleGroupPermission{
			{PermissionKey: "users", PermissionValue: 1},    // bit 0
			{PermissionKey: "users", PermissionValue: 4},    // bit 2
			{PermissionKey: "settings", PermissionValue: 2}, // bit 1
		},
	}
	svc := NewPermissionService(repo)

	merged, err := svc.GetMergedPermissions(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if merged["users"] != 5 {
		t.Errorf("expected 5 for users, got %d", merged["users"])
	}
	if merged["settings"] != 2 {
		t.Errorf("expected 2 for settings, got %d", merged["settings"])
	}
}

func TestCheckPermission(t *testing.T) {
	repo := &mockPermissionRepo{
		perms: []entity.RoleGroupPermission{
			{PermissionKey: "docs", PermissionValue: 2}, // bit 1 is set
		},
	}
	svc := NewPermissionService(repo)

	// Super Admin bypass
	res, _ := svc.CheckPermission(1, true, "no_access", 5)
	if !res {
		t.Errorf("super admin should bypass")
	}

	// Normal check
	res, _ = svc.CheckPermission(1, false, "docs", 1)
	if !res {
		t.Errorf("expected to have docs bit 1")
	}

	res, _ = svc.CheckPermission(1, false, "docs", 0)
	if res {
		t.Errorf("expected to NOT have docs bit 0")
	}
}
