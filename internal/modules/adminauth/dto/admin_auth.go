package dto

type AdminLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AdminLoginResponse struct {
	Token string `json:"token"`
}

type AdminMeResponse struct {
	ID           uint    `json:"id"`
	Username     string  `json:"username"`
	FullName     string  `json:"full_name"`
	Email        *string `json:"email"`
	PhoneNumber  *string `json:"phone_number"`
	IsSuperAdmin bool    `json:"is_super_admin"`
	IsActive     bool    `json:"is_active"`
}
