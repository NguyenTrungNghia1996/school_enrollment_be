package service

import (
	"errors"
	"math"
	"time"

	"go_be_enrollment/internal/modules/subject/dto"
	"go_be_enrollment/internal/modules/subject/entity"
	"go_be_enrollment/internal/modules/subject/repository"
)

type SubjectService interface {
	GetList(filter *dto.SubjectFilter) (*dto.PaginatedSubjectRes, error)
	GetDetail(id uint) (*dto.SubjectRes, error)
	Create(req *dto.SubjectCreateReq) (*dto.SubjectRes, error)
	Update(id uint, req *dto.SubjectUpdateReq) (*dto.SubjectRes, error)
	UpdateStatus(id uint, req *dto.SubjectStatusReq) error
	GetActiveList() ([]dto.SubjectRes, error)
}

type subjectService struct {
	repo repository.SubjectRepository
}

func NewSubjectService(repo repository.SubjectRepository) SubjectService {
	return &subjectService{repo: repo}
}

func (s *subjectService) GetList(filter *dto.SubjectFilter) (*dto.PaginatedSubjectRes, error) {
	subjects, total, err := s.repo.GetList(filter)
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

	var resData []dto.SubjectRes
	for _, sub := range subjects {
		resData = append(resData, dto.SubjectRes{
			ID:           sub.ID,
			Code:         sub.Code,
			Name:         sub.Name,
			DisplayOrder: sub.DisplayOrder,
			IsActive:     sub.IsActive,
			CreatedAt:    sub.CreatedAt.Format(time.RFC3339),
			UpdatedAt:    sub.UpdatedAt.Format(time.RFC3339),
		})
	}
	
	if resData == nil {
		resData = []dto.SubjectRes{}
	}

	return &dto.PaginatedSubjectRes{
		Data:       resData,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}

func (s *subjectService) GetDetail(id uint) (*dto.SubjectRes, error) {
	sub, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("môn học không tồn tại")
	}

	return &dto.SubjectRes{
		ID:           sub.ID,
		Code:         sub.Code,
		Name:         sub.Name,
		DisplayOrder: sub.DisplayOrder,
		IsActive:     sub.IsActive,
		CreatedAt:    sub.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    sub.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *subjectService) Create(req *dto.SubjectCreateReq) (*dto.SubjectRes, error) {
	if s.repo.CheckCodeExists(req.Code, 0) {
		return nil, errors.New("mã môn học đã tồn tại")
	}

	isActive := true
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	sub := &entity.Subject{
		Code:         req.Code,
		Name:         req.Name,
		DisplayOrder: req.DisplayOrder,
		IsActive:     isActive,
	}

	if err := s.repo.Create(sub); err != nil {
		return nil, err
	}

	return &dto.SubjectRes{
		ID:           sub.ID,
		Code:         sub.Code,
		Name:         sub.Name,
		DisplayOrder: sub.DisplayOrder,
		IsActive:     sub.IsActive,
		CreatedAt:    sub.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    sub.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *subjectService) Update(id uint, req *dto.SubjectUpdateReq) (*dto.SubjectRes, error) {
	sub, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("môn học không tồn tại")
	}

	if s.repo.CheckCodeExists(req.Code, id) {
		return nil, errors.New("mã môn học đã tồn tại")
	}

	sub.Code = req.Code
	sub.Name = req.Name
	sub.DisplayOrder = req.DisplayOrder
	if req.IsActive != nil {
		sub.IsActive = *req.IsActive
	}

	if err := s.repo.Update(sub); err != nil {
		return nil, err
	}

	return &dto.SubjectRes{
		ID:           sub.ID,
		Code:         sub.Code,
		Name:         sub.Name,
		DisplayOrder: sub.DisplayOrder,
		IsActive:     sub.IsActive,
		CreatedAt:    sub.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    sub.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *subjectService) UpdateStatus(id uint, req *dto.SubjectStatusReq) error {
	sub, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("môn học không tồn tại")
	}

	sub.IsActive = req.IsActive

	return s.repo.Update(sub)
}

func (s *subjectService) GetActiveList() ([]dto.SubjectRes, error) {
	subjects, err := s.repo.GetActiveList()
	if err != nil {
		return nil, err
	}

	var resData []dto.SubjectRes
	for _, sub := range subjects {
		resData = append(resData, dto.SubjectRes{
			ID:           sub.ID,
			Code:         sub.Code,
			Name:         sub.Name,
			DisplayOrder: sub.DisplayOrder,
			IsActive:     sub.IsActive,
			CreatedAt:    sub.CreatedAt.Format(time.RFC3339),
			UpdatedAt:    sub.UpdatedAt.Format(time.RFC3339),
		})
	}
	
	if resData == nil {
		resData = []dto.SubjectRes{}
	}

	return resData, nil
}
