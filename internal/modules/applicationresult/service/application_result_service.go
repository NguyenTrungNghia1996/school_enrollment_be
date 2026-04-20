package service

import (
	"errors"
	"time"

	"go_be_enrollment/internal/modules/applicationresult/dto"
	"go_be_enrollment/internal/modules/applicationresult/entity"
	"go_be_enrollment/internal/modules/applicationresult/repository"
	app_repo "go_be_enrollment/internal/modules/application/repository"
	score_repo "go_be_enrollment/internal/modules/applicationexamscore/repository"
	period_repo "go_be_enrollment/internal/modules/admissionperiod/repository"
)

type ApplicationResultService interface {
	GetAdminResult(appID uint) (*dto.ApplicationResultRes, error)
	GetUserResult(appID, userID uint) (*dto.ApplicationResultRes, error)
	UpdateManual(appID uint, req *dto.ApplicationResultUpdateReq) (*dto.ApplicationResultRes, error)
	RecalculateSingle(appID uint) (*dto.ApplicationResultRes, error)
	RecalculateRankingByPeriod(periodID uint) error
}

type applicationResultService struct {
	repo       repository.ApplicationResultRepository
	appRepo    app_repo.ApplicationRepository
	scoreRepo  score_repo.ApplicationExamScoreRepository
	periodRepo period_repo.AdmissionPeriodRepository
}

func NewApplicationResultService(repo repository.ApplicationResultRepository, appRepo app_repo.ApplicationRepository, scoreRepo score_repo.ApplicationExamScoreRepository, periodRepo period_repo.AdmissionPeriodRepository) ApplicationResultService {
	return &applicationResultService{repo: repo, appRepo: appRepo, scoreRepo: scoreRepo, periodRepo: periodRepo}
}

func (s *applicationResultService) getOrCreateResult(appID uint) (*entity.ApplicationResult, error) {
	res, err := s.repo.FindByApplicationID(appID)
	if err != nil {
		return nil, err
	}
	if res == nil {
		res = &entity.ApplicationResult{
			ApplicationID: appID,
			ResultStatus:  "Pending",
		}
	}
	return res, nil
}

func (s *applicationResultService) computeFinal(res *entity.ApplicationResult) {
	res.FinalTotalScore = res.TotalScore + res.PriorityScore + res.AdditionalScore
}

func (s *applicationResultService) GetAdminResult(appID uint) (*dto.ApplicationResultRes, error) {
	_, err := s.appRepo.GetAdminDetail(appID)
	if err != nil {
		return nil, errors.New("hồ sơ không tồn tại")
	}

	res, err := s.repo.FindByApplicationID(appID)
	if err != nil {
		return nil, err
	}
	
	if res == nil {
		res = &entity.ApplicationResult{
			ApplicationID:   appID,
			ResultStatus:    "Pending",
			TotalScore:      0,
			PriorityScore:   0,
			AdditionalScore: 0,
			FinalTotalScore: 0,
		}
	}

	return &dto.ApplicationResultRes{
		ID:              res.ID,
		ApplicationID:   res.ApplicationID,
		TotalScore:      res.TotalScore,
		PriorityScore:   res.PriorityScore,
		AdditionalScore: res.AdditionalScore,
		FinalTotalScore: res.FinalTotalScore,
		Ranking:         res.Ranking,
		ResultStatus:    res.ResultStatus,
		Notes:           res.Notes,
		CreatedAt:       res.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       res.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *applicationResultService) GetUserResult(appID, userID uint) (*dto.ApplicationResultRes, error) {
	_, err := s.appRepo.GetUserDetail(appID, userID)
	if err != nil {
		return nil, errors.New("hồ sơ không tồn tại hoặc không thuộc quyền sở hữu của bạn")
	}

	return s.GetAdminResult(appID)
}

func (s *applicationResultService) RecalculateSingle(appID uint) (*dto.ApplicationResultRes, error) {
	_, err := s.appRepo.GetAdminDetail(appID)
	if err != nil {
		return nil, errors.New("hồ sơ không tồn tại")
	}

	res, err := s.getOrCreateResult(appID)
	if err != nil {
		return nil, err
	}

	scores, err := s.scoreRepo.GetByApplicationID(appID)
	if err != nil {
		return nil, err
	}

	total := 0.0
	for _, sc := range scores {
		if sc.FinalScore != nil {
			total += *sc.FinalScore
		}
	}

	res.TotalScore = total
	s.computeFinal(res)

	if err := s.repo.Save(res); err != nil {
		return nil, err
	}

	return s.GetAdminResult(appID)
}

func (s *applicationResultService) UpdateManual(appID uint, req *dto.ApplicationResultUpdateReq) (*dto.ApplicationResultRes, error) {
	_, err := s.appRepo.GetAdminDetail(appID)
	if err != nil {
		return nil, errors.New("hồ sơ không tồn tại")
	}

	if req.PriorityScore < 0 || req.AdditionalScore < 0 {
		return nil, errors.New("điểm không được lưu giá trị âm")
	}

	res, err := s.getOrCreateResult(appID)
	if err != nil {
		return nil, err
	}

	res.PriorityScore = req.PriorityScore
	res.AdditionalScore = req.AdditionalScore
	res.ResultStatus = req.ResultStatus
	res.Notes = req.Notes
	
	s.computeFinal(res)

	if err := s.repo.Save(res); err != nil {
		return nil, err
	}

	return s.GetAdminResult(appID)
}

func (s *applicationResultService) RecalculateRankingByPeriod(periodID uint) error {
	_, err := s.periodRepo.FindByID(periodID)
	if err != nil {
		return errors.New("kỳ tuyển sinh không tồn tại")
	}

	list, err := s.repo.GetByAdmissionPeriod(periodID)
	if err != nil {
		return err
	}

	rank := 1
	for i := range list {
		list[i].Ranking = &rank
		rank++
	}

	if err := s.repo.UpdateBatch(list); err != nil {
		return err
	}
	return nil
}
