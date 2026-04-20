package dto

type ExaminerFilter struct {
	Keyword string `query:"keyword"`
	Page    int    `query:"page"`
	Limit   int    `query:"limit"`
}

type ExaminerCreateReq struct {
	FullName         string `json:"full_name" validate:"required"`
	OrganizationName string `json:"organization_name"`
	PhoneNumber      string `json:"phone_number"`
}

type ExaminerUpdateReq struct {
	FullName         string `json:"full_name" validate:"required"`
	OrganizationName string `json:"organization_name"`
	PhoneNumber      string `json:"phone_number"`
}

type ExaminerRes struct {
	ID               uint   `json:"id"`
	FullName         string `json:"full_name"`
	OrganizationName string `json:"organization_name"`
	PhoneNumber      string `json:"phone_number"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type PaginatedExaminerRes struct {
	Data       []ExaminerRes `json:"data"`
	Total      int64         `json:"total"`
	Page       int           `json:"page"`
	Limit      int           `json:"limit"`
	TotalPages int           `json:"total_pages"`
}
