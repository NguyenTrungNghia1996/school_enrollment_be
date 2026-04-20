package service

import (
	"errors"
	"math"

	"go_be_enrollment/internal/modules/adminauth/entity"
	"go_be_enrollment/internal/modules/adminuser/dto"
	"go_be_enrollment/internal/modules/adminuser/repository"

	"golang.org/x/crypto/bcrypt"
)

type AdminUserService interface {
	GetList(filter *dto.AdminUserFilter) (*dto.PaginatedAdminUserRes, error)
	GetDetail(id uint) (*dto.AdminUserRes, error)
	Create(req *dto.AdminUserCreateReq) (*dto.AdminUserRes, error)
	Update(id uint, req *dto.AdminUserUpdateReq) (*dto.AdminUserRes, error)
	UpdateStatus(id uint, req *dto.AdminUserStatusReq) error
	ResetPassword(id uint, req *dto.AdminUserResetPassReq) error
	GetRoleGroups(id uint) ([]dto.AssignedRoleGroupRes, error)
	UpdateRoleGroups(id uint, req *dto.AdminUserRoleGroupReq) ([]dto.AssignedRoleGroupRes, error)
}

type adminUserService struct {
	repo repository.AdminUserRepository
}

func NewAdminUserService(repo repository.AdminUserRepository) AdminUserService {
	return &adminUserService{repo: repo}
}

func mapToDto(m *entity.AdminUser) *dto.AdminUserRes {
	return &dto.AdminUserRes{
		ID:           m.ID,
		Username:     m.Username,
		FullName:     m.FullName,
		Email:        m.Email,
		PhoneNumber:  m.PhoneNumber,
		IsSuperAdmin: m.IsSuperAdmin,
		IsActive:     m.IsActive,
		CreatedAt:    m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (s *adminUserService) GetList(filter *dto.AdminUserFilter) (*dto.PaginatedAdminUserRes, error) {
	admins, total, err := s.repo.FindAll(filter)
	if err != nil {
		return nil, err
	}

	var data []dto.AdminUserRes
	for _, a := range admins {
		data = append(data, *mapToDto(&a))
	}

	page := filter.Page
	if page < 1 {page = 1}
	limit := filter.Limit
	if limit < 1 {limit = 10}

	return &dto.PaginatedAdminUserRes{
		Data:       data,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: int(math.Ceil(float64(total) / float64(limit))),
	}, nil
}

func (s *adminUserService) GetDetail(id uint) (*dto.AdminUserRes, error) {
	admin, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("không tìm thấy admin")
	}
	return mapToDto(admin), nil
}

func (s *adminUserService) Create(req *dto.AdminUserCreateReq) (*dto.AdminUserRes, error) {
	if s.repo.CheckUsernameExists(req.Username, 0) {
		return nil, errors.New("tên đăng nhập đã tồn tại")
	}
	if req.Email != nil && *req.Email != "" && s.repo.CheckEmailExists(*req.Email, 0) {
		return nil, errors.New("email đã được sử dụng")
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("không thể băm mật khẩu")
	}

	newAdmin := &entity.AdminUser{
		Username:     req.Username,
		PasswordHash: string(hashedBytes),
		FullName:     req.FullName,
		Email:        req.Email,
		PhoneNumber:  req.PhoneNumber,
		IsSuperAdmin: req.IsSuperAdmin,
		IsActive:     true,
	}

	if err := s.repo.Create(newAdmin); err != nil {
		return nil, errors.New("lỗi tạo tài khoản")
	}

	return mapToDto(newAdmin), nil
}

func (s *adminUserService) Update(id uint, req *dto.AdminUserUpdateReq) (*dto.AdminUserRes, error) {
	admin, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("không tìm thấy admin")
	}

	if req.Email != nil && *req.Email != "" && s.repo.CheckEmailExists(*req.Email, id) {
		return nil, errors.New("email đã được sử dụng")
	}

	admin.FullName = req.FullName
	admin.Email = req.Email
	admin.PhoneNumber = req.PhoneNumber
	admin.IsSuperAdmin = req.IsSuperAdmin

	if err := s.repo.Update(admin); err != nil {
		return nil, errors.New("lỗi cập nhật tài khoản")
	}

	return mapToDto(admin), nil
}

func (s *adminUserService) UpdateStatus(id uint, req *dto.AdminUserStatusReq) error {
	admin, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("không tìm thấy admin")
	}

	admin.IsActive = req.IsActive
	if err := s.repo.Update(admin); err != nil {
		return errors.New("lỗi cập nhật trạng thái")
	}
	return nil
}

func (s *adminUserService) ResetPassword(id uint, req *dto.AdminUserResetPassReq) error {
	admin, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("không tìm thấy admin")
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("không thể băm mật khẩu")
	}

	admin.PasswordHash = string(hashedBytes)
	if err := s.repo.Update(admin); err != nil {
		return errors.New("lỗi reset mật khẩu")
	}
	return nil
}

func (s *adminUserService) GetRoleGroups(id uint) ([]dto.AssignedRoleGroupRes, error) {
	roles, err := s.repo.GetAssignedRoleGroups(id)
	if err != nil {
		return nil, errors.New("lỗi lấy danh sách nhóm quyền")
	}

	var res []dto.AssignedRoleGroupRes
	for _, r := range roles {
		res = append(res, dto.AssignedRoleGroupRes{
			ID:   r.ID,
			Code: r.Code,
			Name: r.Name,
		})
	}
	if res == nil {
		res = []dto.AssignedRoleGroupRes{}
	}
	return res, nil
}

func (s *adminUserService) UpdateRoleGroups(id uint, req *dto.AdminUserRoleGroupReq) ([]dto.AssignedRoleGroupRes, error) {
	// Kiểm tra admin có tồn tại hay không
	_, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("không tìm thấy admin")
	}

	// Lọc bỏ các giá trị trùng lặp trong input list
	uniqueIDs := make(map[uint]bool)
	var finalIDs []uint
	for _, rid := range req.RoleGroupIDs {
		if !uniqueIDs[rid] {
			uniqueIDs[rid] = true
			finalIDs = append(finalIDs, rid)
		}
	}

	// Kiểm tra xem tất cả các role group ID có thực sự tồn tại hợp lệ hay không
	ok, err := s.repo.CheckRoleGroupsExist(finalIDs)
	if err != nil || !ok {
		return nil, errors.New("một hoặc nhiều nhóm quyền không hợp lệ")
	}

	// Cập nhật bằng transaction trong repo
	if err := s.repo.ReplaceRoleGroups(id, finalIDs); err != nil {
		return nil, errors.New("lỗi cập nhật nhóm quyền")
	}

	// Trả về danh sách mới nhất
	return s.GetRoleGroups(id)
}
