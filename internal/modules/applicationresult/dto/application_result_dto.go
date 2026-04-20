package dto

type ApplicationResultUpdateReq struct {
	PriorityScore   float64 `json:"priority_score"`
	AdditionalScore float64 `json:"additional_score"`
	ResultStatus    string  `json:"result_status" validate:"required,oneof=Pending Passed Failed Waitlisted"`
	Notes           string  `json:"notes"`
}

type ApplicationResultRes struct {
	ID              uint    `json:"id"`
	ApplicationID   uint    `json:"application_id"`
	TotalScore      float64 `json:"total_score"`
	PriorityScore   float64 `json:"priority_score"`
	AdditionalScore float64 `json:"additional_score"`
	FinalTotalScore float64 `json:"final_total_score"`
	Ranking         *int    `json:"ranking"`
	ResultStatus    string  `json:"result_status"`
	Notes           string  `json:"notes"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}
