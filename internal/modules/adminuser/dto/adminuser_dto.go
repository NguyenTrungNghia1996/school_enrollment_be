package dto

type AdminUserFilter struct {
	Keyword  string `query:"keyword"`
	IsActive *bool  `query:"is_active"`
	Page     int    `query:"page"`
	Limit    int    `query:"limit"`
}

type AdminUserCreateReq struct {
	Username     string  `json:"username" validate:"required"`
	Password     string  `json:"password" validate:"required,min=6"`
	FullName     string  `json:"full_name" validate:"required"`
	Email        *string `json:"email"`
	PhoneNumber  *string `json:"phone_number"`
	IsSuperAdmin bool    `json:"is_super_admin"`
}

type AdminUserUpdateReq struct {
	FullName     string  `json:"full_name" validate:"required"`
	Email        *string `json:"email"`
	PhoneNumber  *string `json:"phone_number"`
	IsSuperAdmin bool    `json:"is_super_admin"`
}

type AdminUserStatusReq struct {
	IsActive bool `json:"is_active"`
}

type AdminUserResetPassReq struct {
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

type AdminUserRes struct {
	ID           uint    `json:"id"`
	Username     string  `json:"username"`
	FullName     string  `json:"full_name"`
	Email        *string `json:"email"`
	PhoneNumber  *string `json:"phone_number"`
	IsSuperAdmin bool    `json:"is_super_admin"`
	IsActive     bool    `json:"is_active"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type PaginatedAdminUserRes struct {
	Data       []AdminUserRes `json:"data"`
	Total      int64          `json:"total"`
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
	TotalPages int            `json:"total_pages"`
}

type AdminUserRoleGroupReq struct {
	RoleGroupIDs []uint `json:"role_group_ids" validate:"required"`
}

type AssignedRoleGroupRes struct {
	ID   uint   `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
