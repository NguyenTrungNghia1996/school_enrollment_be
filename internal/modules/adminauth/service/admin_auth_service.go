package service

import (
	"errors"
	"go_be_enrollment/internal/config"
	"go_be_enrollment/internal/modules/adminauth/dto"
	"go_be_enrollment/internal/modules/adminauth/repository"
	"go_be_enrollment/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type AdminAuthService interface {
	Login(req *dto.AdminLoginRequest) (*dto.AdminLoginResponse, error)
	GetMe(adminID uint) (*dto.AdminMeResponse, error)
}

type adminAuthService struct {
	repo repository.AdminUserRepository
	cfg  *config.Config
}

func NewAdminAuthService(repo repository.AdminUserRepository, cfg *config.Config) AdminAuthService {
	return &adminAuthService{repo: repo, cfg: cfg}
}

func (s *adminAuthService) Login(req *dto.AdminLoginRequest) (*dto.AdminLoginResponse, error) {
	admin, err := s.repo.FindByUsername(req.Username)
	if err != nil {
		return nil, errors.New("tên đăng nhập hoặc mật khẩu không đúng")
	}

	if !admin.IsActive {
		return nil, errors.New("tài khoản quản trị viên đã bị khóa")
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("tên đăng nhập hoặc mật khẩu không đúng")
	}

	token, err := utils.GenerateAdminToken(admin.ID, admin.Username, admin.IsSuperAdmin, s.cfg.JWTSecret, s.cfg.JWTExpiresIn)
	if err != nil {
		return nil, errors.New("lỗi tạo token")
	}

	return &dto.AdminLoginResponse{
		Token: token,
	}, nil
}

func (s *adminAuthService) GetMe(adminID uint) (*dto.AdminMeResponse, error) {
	admin, err := s.repo.FindByID(adminID)
	if err != nil {
		return nil, errors.New("không tìm thấy thông tin quản trị viên")
	}

	return &dto.AdminMeResponse{
		ID:           admin.ID,
		Username:     admin.Username,
		FullName:     admin.FullName,
		Email:        admin.Email,
		PhoneNumber:  admin.PhoneNumber,
		IsSuperAdmin: admin.IsSuperAdmin,
		IsActive:     admin.IsActive,
	}, nil
}
