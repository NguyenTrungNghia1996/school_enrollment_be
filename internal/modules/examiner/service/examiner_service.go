package service

import (
	"errors"
	"math"
	"time"

	"go_be_enrollment/internal/modules/examiner/dto"
	"go_be_enrollment/internal/modules/examiner/entity"
	"go_be_enrollment/internal/modules/examiner/repository"
)

type ExaminerService interface {
	GetList(filter *dto.ExaminerFilter) (*dto.PaginatedExaminerRes, error)
	GetDetail(id uint) (*dto.ExaminerRes, error)
	Create(req *dto.ExaminerCreateReq) (*dto.ExaminerRes, error)
	Update(id uint, req *dto.ExaminerUpdateReq) (*dto.ExaminerRes, error)
	Delete(id uint) error
}

type examinerService struct {
	repo repository.ExaminerRepository
}

func NewExaminerService(repo repository.ExaminerRepository) ExaminerService {
	return &examinerService{repo: repo}
}

func (s *examinerService) GetList(filter *dto.ExaminerFilter) (*dto.PaginatedExaminerRes, error) {
	examiners, total, err := s.repo.GetList(filter)
	if err != nil {
		return nil, err
	}

	limit := filter.Limit
	if limit <= 0 {
		limit = 10
	} else if limit > 100 {
		limit = 100
	}

	page := filter.Page
	if page <= 0 {
		page = 1
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	var resData []dto.ExaminerRes
	for _, e := range examiners {
		resData = append(resData, dto.ExaminerRes{
			ID:               e.ID,
			FullName:         e.FullName,
			OrganizationName: e.OrganizationName,
			PhoneNumber:      e.PhoneNumber,
			CreatedAt:        e.CreatedAt.Format(time.RFC3339),
			UpdatedAt:        e.UpdatedAt.Format(time.RFC3339),
		})
	}
	
	if resData == nil {
		resData = []dto.ExaminerRes{}
	}

	return &dto.PaginatedExaminerRes{
		Data:       resData,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}

func (s *examinerService) GetDetail(id uint) (*dto.ExaminerRes, error) {
	examiner, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("cán bộ coi thi không tồn tại")
	}

	return &dto.ExaminerRes{
		ID:               examiner.ID,
		FullName:         examiner.FullName,
		OrganizationName: examiner.OrganizationName,
		PhoneNumber:      examiner.PhoneNumber,
		CreatedAt:        examiner.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        examiner.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *examinerService) Create(req *dto.ExaminerCreateReq) (*dto.ExaminerRes, error) {
	examiner := &entity.Examiner{
		FullName:         req.FullName,
		OrganizationName: req.OrganizationName,
		PhoneNumber:      req.PhoneNumber,
	}

	if err := s.repo.Create(examiner); err != nil {
		return nil, err
	}

	return &dto.ExaminerRes{
		ID:               examiner.ID,
		FullName:         examiner.FullName,
		OrganizationName: examiner.OrganizationName,
		PhoneNumber:      examiner.PhoneNumber,
		CreatedAt:        examiner.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        examiner.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *examinerService) Update(id uint, req *dto.ExaminerUpdateReq) (*dto.ExaminerRes, error) {
	examiner, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("cán bộ coi thi không tồn tại")
	}

	examiner.FullName = req.FullName
	examiner.OrganizationName = req.OrganizationName
	examiner.PhoneNumber = req.PhoneNumber

	if err := s.repo.Update(examiner); err != nil {
		return nil, err
	}

	return &dto.ExaminerRes{
		ID:               examiner.ID,
		FullName:         examiner.FullName,
		OrganizationName: examiner.OrganizationName,
		PhoneNumber:      examiner.PhoneNumber,
		CreatedAt:        examiner.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        examiner.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *examinerService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("cán bộ coi thi không tồn tại")
	}

	hasRelated, err := s.repo.HasRelatedAssignments(id)
	if err != nil {
		return err
	}
	if hasRelated {
		return errors.New("không thể xóa vì cán bộ đã được phân công coi thi")
	}

	return s.repo.Delete(id)
}
