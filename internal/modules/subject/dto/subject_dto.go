package dto

type SubjectFilter struct {
	Keyword  string `query:"keyword"`
	IsActive *bool  `query:"is_active"`
	Page     int    `query:"page"`
	Limit    int    `query:"limit"`
}

type SubjectCreateReq struct {
	Code         string `json:"code" validate:"required"`
	Name         string `json:"name" validate:"required"`
	DisplayOrder int    `json:"display_order"`
	IsActive     *bool  `json:"is_active"`
}

type SubjectUpdateReq struct {
	Code         string `json:"code" validate:"required"`
	Name         string `json:"name" validate:"required"`
	DisplayOrder int    `json:"display_order"`
	IsActive     *bool  `json:"is_active"`
}

type SubjectStatusReq struct {
	IsActive bool `json:"is_active" validate:"required"`
}

type SubjectRes struct {
	ID           uint   `json:"id"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	DisplayOrder int    `json:"display_order"`
	IsActive     bool   `json:"is_active"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type PaginatedSubjectRes struct {
	Data       []SubjectRes `json:"data"`
	Total      int64        `json:"total"`
	Page       int          `json:"page"`
	Limit      int          `json:"limit"`
	TotalPages int          `json:"total_pages"`
}
