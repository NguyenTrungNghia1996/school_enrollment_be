package dto

type WardUnitFilter struct {
	ProvinceID *uint   `query:"province_id"`
	UnitType   *string `query:"unit_type"`
	Keyword    string  `query:"keyword"`
	IsActive   *bool   `query:"is_active"`
	Page       int     `query:"page"`
	Limit      int     `query:"limit"`
}

type WardUnitCreateReq struct {
	ProvinceID uint   `json:"province_id" validate:"required"`
	Code       string `json:"code" validate:"required"`
	Name       string `json:"name" validate:"required"`
	UnitType   string `json:"unit_type" validate:"required,oneof=Ward Commune SpecialZone"`
	IsActive   bool   `json:"is_active"`
}

type WardUnitUpdateReq struct {
	ProvinceID uint   `json:"province_id" validate:"required"`
	Code       string `json:"code" validate:"required"`
	Name       string `json:"name" validate:"required"`
	UnitType   string `json:"unit_type" validate:"required,oneof=Ward Commune SpecialZone"`
	IsActive   bool   `json:"is_active"`
}

type WardUnitStatusReq struct {
	IsActive bool `json:"is_active"`
}

type WardUnitRes struct {
	ID         uint   `json:"id"`
	ProvinceID uint   `json:"province_id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	UnitType   string `json:"unit_type"`
	IsActive   bool   `json:"is_active"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type PaginatedWardUnitRes struct {
	Data       []WardUnitRes `json:"data"`
	Total      int64         `json:"total"`
	Page       int           `json:"page"`
	Limit      int           `json:"limit"`
	TotalPages int           `json:"total_pages"`
}
