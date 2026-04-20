package service

import (
	"errors"
	"math"

	"go_be_enrollment/internal/modules/wardunit/dto"
	"go_be_enrollment/internal/modules/wardunit/entity"
	"go_be_enrollment/internal/modules/wardunit/repository"
)

type WardUnitService interface {
	GetList(filter *dto.WardUnitFilter) (*dto.PaginatedWardUnitRes, error)
	GetDetail(id uint) (*dto.WardUnitRes, error)
	Create(req *dto.WardUnitCreateReq) (*dto.WardUnitRes, error)
	Update(id uint, req *dto.WardUnitUpdateReq) (*dto.WardUnitRes, error)
	UpdateStatus(id uint, req *dto.WardUnitStatusReq) error
	GetActiveListByProvince(provinceID uint) ([]dto.WardUnitRes, error)
}

type wardUnitService struct {
	repo repository.WardUnitRepository
}

func NewWardUnitService(repo repository.WardUnitRepository) WardUnitService {
	return &wardUnitService{repo: repo}
}

func mapToDto(w *entity.WardUnit) dto.WardUnitRes {
	return dto.WardUnitRes{
		ID:         w.ID,
		ProvinceID: w.ProvinceID,
		Code:       w.Code,
		Name:       w.Name,
		UnitType:   w.UnitType,
		IsActive:   w.IsActive,
		CreatedAt:  w.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  w.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func validateUnitType(t string) bool {
	return t == "Ward" || t == "Commune" || t == "SpecialZone"
}

func (s *wardUnitService) GetList(filter *dto.WardUnitFilter) (*dto.PaginatedWardUnitRes, error) {
	list, total, err := s.repo.GetList(filter)
	if err != nil {
		return nil, errors.New("lỗi lấy danh sách ward unit")
	}

	res := make([]dto.WardUnitRes, 0)
	for _, w := range list {
		res = append(res, mapToDto(&w))
	}

	limit := filter.Limit
	if limit <= 0 {
		limit = 10
	} else if limit > 100 {
		limit = 100
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &dto.PaginatedWardUnitRes{
		Data:       res,
		Total:      total,
		Page:       filter.Page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}

func (s *wardUnitService) GetDetail(id uint) (*dto.WardUnitRes, error) {
	item, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("không tìm thấy ward unit")
	}

	res := mapToDto(item)
	return &res, nil
}

func (s *wardUnitService) Create(req *dto.WardUnitCreateReq) (*dto.WardUnitRes, error) {
	if !validateUnitType(req.UnitType) {
		return nil, errors.New("unit_type không hợp lệ")
	}

	if !s.repo.CheckProvinceExists(req.ProvinceID) {
		return nil, errors.New("province_id không tồn tại")
	}

	if s.repo.CheckNameExistsInProvince(req.ProvinceID, req.Name, 0) {
		return nil, errors.New("tên này đã tồn tại trong tỉnh/thành")
	}

	w := &entity.WardUnit{
		ProvinceID: req.ProvinceID,
		Code:       req.Code,
		Name:       req.Name,
		UnitType:   req.UnitType,
		IsActive:   req.IsActive,
	}

	if err := s.repo.Create(w); err != nil {
		return nil, errors.New("lỗi tạo ward unit")
	}

	res := mapToDto(w)
	return &res, nil
}

func (s *wardUnitService) Update(id uint, req *dto.WardUnitUpdateReq) (*dto.WardUnitRes, error) {
	if !validateUnitType(req.UnitType) {
		return nil, errors.New("unit_type không hợp lệ")
	}

	w, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("không tìm thấy ward unit")
	}

	if req.ProvinceID != w.ProvinceID && !s.repo.CheckProvinceExists(req.ProvinceID) {
		return nil, errors.New("province_id không tồn tại")
	}

	if s.repo.CheckNameExistsInProvince(req.ProvinceID, req.Name, id) {
		return nil, errors.New("tên này đã tồn tại trong tỉnh/thành")
	}

	w.ProvinceID = req.ProvinceID
	w.Code = req.Code
	w.Name = req.Name
	w.UnitType = req.UnitType
	w.IsActive = req.IsActive

	if err := s.repo.Update(w); err != nil {
		return nil, errors.New("lỗi cập nhật ward unit")
	}

	res := mapToDto(w)
	return &res, nil
}

func (s *wardUnitService) UpdateStatus(id uint, req *dto.WardUnitStatusReq) error {
	w, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("không tìm thấy ward unit")
	}

	w.IsActive = req.IsActive

	if err := s.repo.Update(w); err != nil {
		return errors.New("lỗi cập nhật trạng thái ward unit")
	}

	return nil
}

func (s *wardUnitService) GetActiveListByProvince(provinceID uint) ([]dto.WardUnitRes, error) {
	if !s.repo.CheckProvinceExists(provinceID) {
		return nil, errors.New("province_id không tồn tại")
	}

	list, err := s.repo.GetActiveListByProvince(provinceID)
	if err != nil {
		return nil, errors.New("lỗi lấy danh sách ward unit")
	}

	res := make([]dto.WardUnitRes, 0)
	for _, w := range list {
		res = append(res, mapToDto(&w))
	}
	return res, nil
}
