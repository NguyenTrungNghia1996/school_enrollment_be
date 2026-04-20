package dto

type ExamScoreItemReq struct {
	SubjectID  uint     `json:"subject_id" validate:"required"`
	RawScore   *float64 `json:"raw_score"`
	BonusScore float64  `json:"bonus_score"`
	FinalScore *float64 `json:"final_score"`
	IsAbsent   bool     `json:"is_absent"`
	Notes      string   `json:"notes"`
}

type ReplaceExamScoresReq struct {
	Scores []ExamScoreItemReq `json:"scores" validate:"dive"`
}

type ExamScoreRes struct {
	ID            uint     `json:"id"`
	ApplicationID uint     `json:"application_id"`
	SubjectID     uint     `json:"subject_id"`
	SubjectCode   string   `json:"subject_code"`
	SubjectName   string   `json:"subject_name"`
	RawScore      *float64 `json:"raw_score"`
	BonusScore    float64  `json:"bonus_score"`
	FinalScore    *float64 `json:"final_score"`
	IsAbsent      bool     `json:"is_absent"`
	Notes         string   `json:"notes"`
}
