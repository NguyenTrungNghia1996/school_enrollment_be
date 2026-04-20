package service

import (
	"errors"
	"math"
	"time"

	"go_be_enrollment/internal/modules/application/dto"
	"go_be_enrollment/internal/modules/application/entity"
	"go_be_enrollment/internal/modules/application/repository"
)

type ApplicationService interface {
	// Admin
	GetAdminList(filter *dto.ApplicationAdminFilter) (*dto.PaginatedApplicationRes, error)
	GetAdminDetail(id uint) (*dto.ApplicationRes, error)
	
	// User
	GetUserList(userID uint, page, limit int) (*dto.PaginatedApplicationRes, error)
	GetUserDetail(id, userID uint) (*dto.ApplicationRes, error)
	Create(userID uint, req *dto.ApplicationReq) (*dto.ApplicationRes, error)
	Update(id, userID uint, req *dto.ApplicationReq) (*dto.ApplicationRes, error)
	Submit(id, userID uint) error
	Approve(id uint) error
	Reject(id uint, req *dto.ApplicationRejectReq) error
}

type applicationService struct {
	repo repository.ApplicationRepository
}

func NewApplicationService(repo repository.ApplicationRepository) ApplicationService {
	return &applicationService{repo: repo}
}

func formatDate(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02")
}

func formatDatePtr(t *time.Time) *string {
	if t == nil {
		return nil
	}
	s := t.Format("2006-01-02 15:04:05")
	return &s
}

func mapToDto(app *entity.Application) dto.ApplicationRes {
	return dto.ApplicationRes{
		ID:                 app.ID,
		UserAccountID:      app.UserAccountID,
		AdmissionPeriodID:  app.AdmissionPeriodID,
		CandidateFullName:  app.CandidateFullName,
		DateOfBirth:        formatDate(app.DateOfBirth),
		Gender:             app.Gender,
		PlaceOfBirth:       app.PlaceOfBirth,
		Ethnicity:          app.Ethnicity,
		NationalID:         app.NationalID,
		ProvinceID:         app.ProvinceID,
		WardUnitID:         app.WardUnitID,
		AddressLine:        app.AddressLine,
		ContactFullName:    app.ContactFullName,
		ContactPhoneNumber: app.ContactPhoneNumber,
		ApplicationStatus:  app.ApplicationStatus,
		RejectReason:       app.RejectReason,
		IsPaid:             app.IsPaid,
		SubmittedAt:        formatDatePtr(app.SubmittedAt),
		CandidateNumber:    app.CandidateNumber,
		CreatedAt:          app.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:          app.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (s *applicationService) GetAdminList(filter *dto.ApplicationAdminFilter) (*dto.PaginatedApplicationRes, error) {
	list, total, err := s.repo.GetAdminList(filter)
	if err != nil {
		return nil, errors.New("lỗi lấy danh sách hồ sơ")
	}

	res := make([]dto.ApplicationRes, 0)
	for _, app := range list {
		res = append(res, mapToDto(&app))
	}

	limit := filter.Limit
	if limit <= 0 {
		limit = 10
	} else if limit > 100 {
		limit = 100
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &dto.PaginatedApplicationRes{
		Data:       res,
		Total:      total,
		Page:       filter.Page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}

func (s *applicationService) GetAdminDetail(id uint) (*dto.ApplicationRes, error) {
	app, err := s.repo.GetAdminDetail(id)
	if err != nil {
		return nil, errors.New("không tìm thấy hồ sơ")
	}
	res := mapToDto(app)
	return &res, nil
}

func (s *applicationService) GetUserList(userID uint, page, limit int) (*dto.PaginatedApplicationRes, error) {
	list, total, err := s.repo.GetUserList(userID, page, limit)
	if err != nil {
		return nil, errors.New("lỗi lấy danh sách hồ sơ cá nhân")
	}

	res := make([]dto.ApplicationRes, 0)
	for _, app := range list {
		res = append(res, mapToDto(&app))
	}

	if limit <= 0 {
		limit = 10
	} else if limit > 100 {
		limit = 100
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &dto.PaginatedApplicationRes{
		Data:       res,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}

func (s *applicationService) GetUserDetail(id, userID uint) (*dto.ApplicationRes, error) {
	app, err := s.repo.GetUserDetail(id, userID)
	if err != nil {
		return nil, errors.New("không tìm thấy hồ sơ của bạn")
	}
	res := mapToDto(app)
	return &res, nil
}

func (s *applicationService) validateForeignKeys(req *dto.ApplicationReq) error {
	if !s.repo.CheckAdmissionPeriodExists(req.AdmissionPeriodID) {
		return errors.New("kỳ tuyển sinh không tồn tại")
	}
	if req.ProvinceID != nil {
		if !s.repo.CheckProvinceExists(*req.ProvinceID) {
			return errors.New("tỉnh/thành phố không tồn tại")
		}
		if req.WardUnitID != nil {
			if !s.repo.CheckWardUnitExists(*req.WardUnitID, *req.ProvinceID) {
				return errors.New("phường/xã không tồn tại hoặc không thuộc tỉnh này")
			}
		}
	} else if req.WardUnitID != nil {
		return errors.New("không thể chọn phường/xã khi chưa chọn tỉnh/thành phố")
	}
	return nil
}

func (s *applicationService) Create(userID uint, req *dto.ApplicationReq) (*dto.ApplicationRes, error) {
	if s.repo.CheckNationalIDExists(req.NationalID, 0) {
		return nil, errors.New("CMND/CCCD đã được sử dụng")
	}

	if err := s.validateForeignKeys(req); err != nil {
		return nil, err
	}

	dob, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		return nil, errors.New("ngày sinh không hợp lệ")
	}

	app := &entity.Application{
		UserAccountID:      userID,
		AdmissionPeriodID:  req.AdmissionPeriodID,
		CandidateFullName:  req.CandidateFullName,
		DateOfBirth:        &dob,
		Gender:             req.Gender,
		PlaceOfBirth:       req.PlaceOfBirth,
		Ethnicity:          req.Ethnicity,
		NationalID:         req.NationalID,
		ProvinceID:         req.ProvinceID,
		WardUnitID:         req.WardUnitID,
		AddressLine:        req.AddressLine,
		ContactFullName:    req.ContactFullName,
		ContactPhoneNumber: req.ContactPhoneNumber,
		ApplicationStatus:  "Draft",
		IsPaid:             false,
	}

	if err := s.repo.Create(app); err != nil {
		return nil, errors.New("lỗi khởi tạo hồ sơ")
	}

	res := mapToDto(app)
	return &res, nil
}

func (s *applicationService) Update(id, userID uint, req *dto.ApplicationReq) (*dto.ApplicationRes, error) {
	app, err := s.repo.GetUserDetail(id, userID)
	if err != nil {
		return nil, errors.New("không tìm thấy hồ sơ của bạn")
	}

	if app.ApplicationStatus != "Draft" {
		return nil, errors.New("hồ sơ đã nộp không thể chỉnh sửa")
	}

	if s.repo.CheckNationalIDExists(req.NationalID, app.ID) {
		return nil, errors.New("CMND/CCCD đã được sử dụng bởi người khác")
	}

	if err := s.validateForeignKeys(req); err != nil {
		return nil, err
	}

	dob, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		return nil, errors.New("ngày sinh không hợp lệ")
	}

	app.AdmissionPeriodID = req.AdmissionPeriodID
	app.CandidateFullName = req.CandidateFullName
	app.DateOfBirth = &dob
	app.Gender = req.Gender
	app.PlaceOfBirth = req.PlaceOfBirth
	app.Ethnicity = req.Ethnicity
	app.NationalID = req.NationalID
	app.ProvinceID = req.ProvinceID
	app.WardUnitID = req.WardUnitID
	app.AddressLine = req.AddressLine
	app.ContactFullName = req.ContactFullName
	app.ContactPhoneNumber = req.ContactPhoneNumber

	if err := s.repo.Update(app); err != nil {
		return nil, errors.New("lỗi cập nhật hồ sơ")
	}

	res := mapToDto(app)
	return &res, nil
}

func (s *applicationService) Submit(id, userID uint) error {
	app, err := s.repo.GetUserDetail(id, userID)
	if err != nil {
		return errors.New("không tìm thấy hồ sơ của bạn")
	}

	if app.ApplicationStatus != "Draft" {
		return errors.New("hồ sơ này không ở trạng thái nháp")
	}

	if app.CandidateFullName == "" || app.DateOfBirth == nil || app.AdmissionPeriodID == 0 {
		return errors.New("hồ sơ thiếu thông tin bắt buộc, không thể gửi")
	}
	
	// TODO: Kiểm tra tài liệu bắt buộc (ví dụ query bảng thư mục hồ sơ) tại đây

	now := time.Now()
	app.ApplicationStatus = "Submitted"
	app.SubmittedAt = &now

	// TODO: Audit log submit application

	if err := s.repo.Update(app); err != nil {
		return errors.New("lỗi khi gửi hồ sơ")
	}

	return nil
}

func (s *applicationService) Approve(id uint) error {
	app, err := s.repo.GetAdminDetail(id)
	if err != nil {
		return errors.New("không tìm thấy hồ sơ")
	}

	if app.ApplicationStatus != "Submitted" {
		return errors.New("chỉ có thể duyệt hồ sơ khi đang ở trạng thái đã gửi (Submitted)")
	}

	app.ApplicationStatus = "Approved"

	// TODO: Audit log approve application

	if err := s.repo.Update(app); err != nil {
		return errors.New("lỗi duyệt hồ sơ")
	}

	return nil
}

func (s *applicationService) Reject(id uint, req *dto.ApplicationRejectReq) error {
	app, err := s.repo.GetAdminDetail(id)
	if err != nil {
		return errors.New("không tìm thấy hồ sơ")
	}

	if app.ApplicationStatus != "Submitted" {
		return errors.New("chỉ có thể từ chối hồ sơ khi đang ở trạng thái đã gửi (Submitted)")
	}

	app.ApplicationStatus = "Rejected"
	app.RejectReason = &req.RejectReason

	// TODO: Audit log reject application

	if err := s.repo.Update(app); err != nil {
		return errors.New("lỗi từ chối hồ sơ")
	}

	return nil
}
