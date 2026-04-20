package service

import (
	"go_be_enrollment/internal/modules/dashboard/dto"
	"go_be_enrollment/internal/modules/dashboard/repository"
)

type DashboardService interface {
	GetSummary(filter *dto.DashboardSummaryFilter) (*dto.DashboardSummaryRes, error)
}

type dashboardService struct {
	repo repository.DashboardRepository
}

func NewDashboardService(repo repository.DashboardRepository) DashboardService {
	return &dashboardService{repo: repo}
}

func (s *dashboardService) GetSummary(filter *dto.DashboardSummaryFilter) (*dto.DashboardSummaryRes, error) {
	return s.repo.GetSummary(filter)
}
