package handler

import (
	"strconv"

	"go_be_enrollment/internal/common/httpresponse"
	"go_be_enrollment/internal/modules/adminuser/dto"
	"go_be_enrollment/internal/modules/adminuser/service"

	"github.com/gofiber/fiber/v2"
)

type AdminUserHandler struct {
	svc service.AdminUserService
}

func NewAdminUserHandler(svc service.AdminUserService) *AdminUserHandler {
	return &AdminUserHandler{svc: svc}
}

func (h *AdminUserHandler) GetList(c *fiber.Ctx) error {
	var filter dto.AdminUserFilter
	if err := c.QueryParser(&filter); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_QUERY", "Tham số không hợp lệ", nil)
	}

	res, err := h.svc.GetList(&filter)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Lấy danh sách thành công", res)
}

func (h *AdminUserHandler) GetDetail(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID không hợp lệ", nil)
	}

	res, err := h.svc.GetDetail(uint(id))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusNotFound, "NOT_FOUND", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Lấy thông tin thành công", res)
}

func (h *AdminUserHandler) Create(c *fiber.Ctx) error {
	var req dto.AdminUserCreateReq
	if err := c.BodyParser(&req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_BODY", "Dữ liệu không hợp lệ", nil)
	}

	// TODO: Thêm Validator layer

	res, err := h.svc.Create(&req)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "CREATE_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Tạo tài khoản thành công", res)
}

func (h *AdminUserHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID không hợp lệ", nil)
	}

	var req dto.AdminUserUpdateReq
	if err := c.BodyParser(&req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_BODY", "Dữ liệu không hợp lệ", nil)
	}

	// TODO: Thêm Validator layer

	res, err := h.svc.Update(uint(id), &req)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "UPDATE_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Cập nhật thành công", res)
}

func (h *AdminUserHandler) UpdateStatus(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID không hợp lệ", nil)
	}

	var req dto.AdminUserStatusReq
	if err := c.BodyParser(&req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_BODY", "Dữ liệu không hợp lệ", nil)
	}

	if err := h.svc.UpdateStatus(uint(id), &req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "UPDATE_STATUS_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Cập nhật trạng thái thành công", nil)
}

func (h *AdminUserHandler) ResetPassword(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID không hợp lệ", nil)
	}

	var req dto.AdminUserResetPassReq
	if err := c.BodyParser(&req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_BODY", "Dữ liệu không hợp lệ", nil)
	}

	// TODO: Thêm Validator layer

	if err := h.svc.ResetPassword(uint(id), &req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "RESET_PASS_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Khôi phục mật khẩu thành công", nil)
}
