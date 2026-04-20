package repository

import (
	"go_be_enrollment/internal/modules/dashboard/dto"
	"gorm.io/gorm"
	"strconv"
)

type DashboardRepository interface {
	GetSummary(filter *dto.DashboardSummaryFilter) (*dto.DashboardSummaryRes, error)
}

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepository{db: db}
}

func (r *dashboardRepository) GetSummary(filter *dto.DashboardSummaryFilter) (*dto.DashboardSummaryRes, error) {
	var res dto.DashboardSummaryRes
	
	// Initialize default empty slices so they don't return null in JSON
	res.ApplicationsByStatus = []dto.KeyValueInt{}
	res.ApplicationsByAdmissionPeriod = []dto.KeyValueInt{}
	res.PaidVsUnpaid = []dto.KeyValueInt{}
	res.AssignedVsUnassignedExamRoom = []dto.KeyValueInt{}
	res.ScoredVsUnscored = []dto.KeyValueInt{}
	res.PassedFailedWaitlisted = []dto.KeyValueInt{}

	appQuery := r.db.Table("applications")
	if filter.AdmissionPeriodID != nil {
		appQuery = appQuery.Where("admission_period_id = ?", *filter.AdmissionPeriodID)
	}

	// 1. Total Applications
	appQuery.Count(&res.TotalApplications)

	// 2. By Status
	var byStatus []struct {
		Status string
		Count  int64
	}
	appQuery.Select("application_status as status, count(id) as count").Group("application_status").Scan(&byStatus)
	for _, item := range byStatus {
		res.ApplicationsByStatus = append(res.ApplicationsByStatus, dto.KeyValueInt{Key: item.Status, Value: item.Count})
	}

	// 3. By Admission Period
	var byPeriod []struct {
		AdmissionPeriodID uint
		Count             int64
	}
	appQuery.Select("admission_period_id, count(id) as count").Group("admission_period_id").Scan(&byPeriod)
	for _, item := range byPeriod {
		res.ApplicationsByAdmissionPeriod = append(res.ApplicationsByAdmissionPeriod, dto.KeyValueInt{Key: strconv.Itoa(int(item.AdmissionPeriodID)), Value: item.Count})
	}

	// 4. Paid vs Unpaid
	var byPaid []struct {
		IsPaid bool
		Count  int64
	}
	appQuery.Select("is_paid, count(id) as count").Group("is_paid").Scan(&byPaid)
	for _, item := range byPaid {
		key := "Unpaid"
		if item.IsPaid {
			key = "Paid"
		}
		res.PaidVsUnpaid = append(res.PaidVsUnpaid, dto.KeyValueInt{Key: key, Value: item.Count})
	}

	// 5. Assigned vs Unassigned
	var assignedCount int64
	aQuery := r.db.Table("exam_room_assignments")
	if filter.AdmissionPeriodID != nil {
		aQuery = aQuery.Joins("JOIN applications ON applications.id = exam_room_assignments.application_id").Where("applications.admission_period_id = ?", *filter.AdmissionPeriodID)
	}
	aQuery.Count(&assignedCount)
	unassignedCount := res.TotalApplications - assignedCount
	res.AssignedVsUnassignedExamRoom = []dto.KeyValueInt{
		{Key: "Assigned", Value: assignedCount},
		{Key: "Unassigned", Value: unassignedCount},
	}

	// 6. Scored vs Unscored
	var scoredCount int64
	sQuery := r.db.Table("application_results")
	if filter.AdmissionPeriodID != nil {
		sQuery = sQuery.Joins("JOIN applications ON applications.id = application_results.application_id").Where("applications.admission_period_id = ?", *filter.AdmissionPeriodID)
	}
	sQuery.Count(&scoredCount)
	unscoredCount := res.TotalApplications - scoredCount
	res.ScoredVsUnscored = []dto.KeyValueInt{
		{Key: "Scored", Value: scoredCount},
		{Key: "Unscored", Value: unscoredCount},
	}

	// 7. Passed vs Failed vs Waitlisted (and Pending)
	var byResult []struct {
		ResultStatus string
		Count        int64
	}
	rQuery := r.db.Table("application_results")
	if filter.AdmissionPeriodID != nil {
		rQuery = rQuery.Joins("JOIN applications ON applications.id = application_results.application_id").Where("applications.admission_period_id = ?", *filter.AdmissionPeriodID)
	}
	rQuery.Select("result_status, count(application_results.id) as count").Group("result_status").Scan(&byResult)
	for _, item := range byResult {
		res.PassedFailedWaitlisted = append(res.PassedFailedWaitlisted, dto.KeyValueInt{Key: item.ResultStatus, Value: item.Count})
	}

	return &res, nil
}
