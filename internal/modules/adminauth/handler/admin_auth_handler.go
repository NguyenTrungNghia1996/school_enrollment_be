package handler

import (
	"go_be_enrollment/internal/common/httpresponse"
	"go_be_enrollment/internal/modules/adminauth/dto"
	"go_be_enrollment/internal/modules/adminauth/service"
	"github.com/gofiber/fiber/v2"
)

type AdminAuthHandler struct {
	svc service.AdminAuthService
}

func NewAdminAuthHandler(svc service.AdminAuthService) *AdminAuthHandler {
	return &AdminAuthHandler{svc: svc}
}

func (h *AdminAuthHandler) Login(c *fiber.Ctx) error {
	var req dto.AdminLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_REQUEST", "Dữ liệu yêu cầu không hợp lệ", nil)
	}

	resp, err := h.svc.Login(&req)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusUnauthorized, "LOGIN_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Đăng nhập thành công", resp)
}

func (h *AdminAuthHandler) GetMe(c *fiber.Ctx) error {
	adminIDVal := c.Locals("admin_id")
	if adminIDVal == nil {
		return httpresponse.Error(c, fiber.StatusUnauthorized, "UNAUTHORIZED", "Không tìm thấy token", nil)
	}

	adminID, ok := adminIDVal.(uint)
	if !ok {
		return httpresponse.Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", "Lỗi lấy thông tin session", nil)
	}

	resp, err := h.svc.GetMe(adminID)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusNotFound, "NOT_FOUND", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Lấy thông tin thành công", resp)
}
