package service

import (
	"errors"

	"go_be_enrollment/internal/modules/applicationexamscore/dto"
	"go_be_enrollment/internal/modules/applicationexamscore/entity"
	"go_be_enrollment/internal/modules/applicationexamscore/repository"
	app_repo "go_be_enrollment/internal/modules/application/repository"
	subject_repo "go_be_enrollment/internal/modules/subject/repository"
)

type ApplicationExamScoreService interface {
	GetByApplicationID(appID uint) ([]dto.ExamScoreRes, error)
	ReplaceScores(appID uint, req *dto.ReplaceExamScoresReq) ([]dto.ExamScoreRes, error)
}

type applicationExamScoreService struct {
	repo        repository.ApplicationExamScoreRepository
	appRepo     app_repo.ApplicationRepository
	subjectRepo subject_repo.SubjectRepository
}

func NewApplicationExamScoreService(repo repository.ApplicationExamScoreRepository, appRepo app_repo.ApplicationRepository, subjectRepo subject_repo.SubjectRepository) ApplicationExamScoreService {
	return &applicationExamScoreService{repo: repo, appRepo: appRepo, subjectRepo: subjectRepo}
}

func (s *applicationExamScoreService) GetByApplicationID(appID uint) ([]dto.ExamScoreRes, error) {
	_, err := s.appRepo.GetAdminDetail(appID)
	if err != nil {
		return nil, errors.New("hồ sơ không tồn tại")
	}

	scores, err := s.repo.GetByApplicationID(appID)
	if err != nil {
		return nil, err
	}

	var res []dto.ExamScoreRes
	for _, sc := range scores {
		code := ""
		name := ""
		if sc.Subject != nil {
			code = sc.Subject.Code
			name = sc.Subject.Name
		}
		res = append(res, dto.ExamScoreRes{
			ID:            sc.ID,
			ApplicationID: sc.ApplicationID,
			SubjectID:     sc.SubjectID,
			SubjectCode:   code,
			SubjectName:   name,
			RawScore:      sc.RawScore,
			BonusScore:    sc.BonusScore,
			FinalScore:    sc.FinalScore,
			IsAbsent:      sc.IsAbsent,
			Notes:         sc.Notes,
		})
	}
	if res == nil {
		res = []dto.ExamScoreRes{}
	}
	return res, nil
}

func (s *applicationExamScoreService) ReplaceScores(appID uint, req *dto.ReplaceExamScoresReq) ([]dto.ExamScoreRes, error) {
	_, err := s.appRepo.GetAdminDetail(appID)
	if err != nil {
		return nil, errors.New("hồ sơ không tồn tại")
	}

	var entities []entity.ApplicationExamScore
	subjectMap := make(map[uint]bool)

	for _, item := range req.Scores {
		if subjectMap[item.SubjectID] {
			return nil, errors.New("danh sách điểm có chứa môn học trùng lặp")
		}
		subjectMap[item.SubjectID] = true

		_, err := s.subjectRepo.FindByID(item.SubjectID)
		if err != nil {
			return nil, errors.New("môn học không tồn tại")
		}

		if item.RawScore != nil && *item.RawScore < 0 {
			return nil, errors.New("điểm thi không được âm")
		}
		if item.BonusScore < 0 {
			return nil, errors.New("điểm ưu tiên không được âm")
		}
		if item.FinalScore != nil && *item.FinalScore < 0 {
			return nil, errors.New("điểm tổng kết không được âm")
		}

		// Xử lý vắng thi (is_absent)
		var tRaw *float64 = item.RawScore
		var tFinal *float64 = item.FinalScore
		
		if item.IsAbsent {
			zero := 0.0
			tRaw = &zero
			tFinal = &zero
		}

		entities = append(entities, entity.ApplicationExamScore{
			ApplicationID: appID,
			SubjectID:     item.SubjectID,
			RawScore:      tRaw,
			BonusScore:    item.BonusScore,
			FinalScore:    tFinal,
			IsAbsent:      item.IsAbsent,
			Notes:         item.Notes,
		})
	}

	if err := s.repo.ReplaceScores(appID, entities); err != nil {
		return nil, err
	}

	return s.GetByApplicationID(appID)
}
