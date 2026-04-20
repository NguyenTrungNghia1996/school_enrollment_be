package dto

type ExamRoomAssignmentFilter struct {
	AdmissionPeriodID *uint  `query:"admission_period_id"`
	ExamRoomID        *uint  `query:"exam_room_id"`
	Keyword           string `query:"keyword"`
	Page              int    `query:"page"`
	Limit             int    `query:"limit"`
}

type ExamRoomAssignmentCreateReq struct {
	ApplicationID uint   `json:"application_id" validate:"required"`
	ExamRoomID    uint   `json:"exam_room_id" validate:"required"`
	SeatNumber    string `json:"seat_number" validate:"required"`
}

type ExamRoomAssignmentUpdateReq struct {
	ExamRoomID uint   `json:"exam_room_id" validate:"required"`
	SeatNumber string `json:"seat_number" validate:"required"`
}

type ExamRoomAssignmentRes struct {
	ID                uint   `json:"id"`
	ApplicationID     uint   `json:"application_id"`
	CandidateFullName string `json:"candidate_full_name"`
	NationalID        string `json:"national_id"`
	ExamRoomID        uint   `json:"exam_room_id"`
	RoomName          string `json:"room_name"`
	SeatNumber        string `json:"seat_number"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type PaginatedExamRoomAssignmentRes struct {
	Data       []ExamRoomAssignmentRes `json:"data"`
	Total      int64                   `json:"total"`
	Page       int                     `json:"page"`
	Limit      int                     `json:"limit"`
	TotalPages int                     `json:"total_pages"`
}
