package dto

type DashboardSummaryFilter struct {
	AdmissionPeriodID *uint `query:"admission_period_id"`
}

type KeyValueInt struct {
	Key   string `json:"key"`
	Value int64  `json:"value"`
}

type DashboardSummaryRes struct {
	TotalApplications             int64         `json:"total_applications"`
	ApplicationsByStatus          []KeyValueInt `json:"applications_by_status"`
	ApplicationsByAdmissionPeriod []KeyValueInt `json:"applications_by_admission_period"`
	PaidVsUnpaid                  []KeyValueInt `json:"paid_vs_unpaid"`
	AssignedVsUnassignedExamRoom  []KeyValueInt `json:"assigned_vs_unassigned_exam_room"`
	ScoredVsUnscored              []KeyValueInt `json:"scored_vs_unscored"`
	PassedFailedWaitlisted        []KeyValueInt `json:"passed_vs_failed_vs_waitlisted"`
}
