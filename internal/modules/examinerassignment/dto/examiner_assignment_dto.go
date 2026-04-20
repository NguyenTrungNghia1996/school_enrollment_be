package dto

type ExaminerAssignmentFilter struct {
	ExaminerID *uint  `query:"examiner_id"`
	ExamRoomID *uint  `query:"exam_room_id"`
	Role       string `query:"role"`
	Page       int    `query:"page"`
	Limit      int    `query:"limit"`
}

type ExaminerAssignmentCreateReq struct {
	ExaminerID uint   `json:"examiner_id" validate:"required"`
	ExamRoomID uint   `json:"exam_room_id" validate:"required"`
	Role       string `json:"role" validate:"required,oneof=Primary Secondary Backup"`
}

type ExaminerAssignmentUpdateReq struct {
	ExaminerID uint   `json:"examiner_id" validate:"required"`
	ExamRoomID uint   `json:"exam_room_id" validate:"required"`
	Role       string `json:"role" validate:"required,oneof=Primary Secondary Backup"`
}

type ExaminerAssignmentRes struct {
	ID               uint   `json:"id"`
	ExaminerID       uint   `json:"examiner_id"`
	ExaminerFullName string `json:"examiner_full_name"`
	ExamRoomID       uint   `json:"exam_room_id"`
	RoomName         string `json:"room_name"`
	Role             string `json:"role"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type PaginatedExaminerAssignmentRes struct {
	Data       []ExaminerAssignmentRes `json:"data"`
	Total      int64                   `json:"total"`
	Page       int                     `json:"page"`
	Limit      int                     `json:"limit"`
	TotalPages int                     `json:"total_pages"`
}
