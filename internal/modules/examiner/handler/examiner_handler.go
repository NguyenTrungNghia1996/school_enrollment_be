package handler

import (
	"strconv"

	"go_be_enrollment/internal/common/httpresponse"
	"go_be_enrollment/internal/modules/examiner/dto"
	"go_be_enrollment/internal/modules/examiner/service"

	"github.com/gofiber/fiber/v2"
)

type ExaminerHandler struct {
	svc service.ExaminerService
}

func NewExaminerHandler(svc service.ExaminerService) *ExaminerHandler {
	return &ExaminerHandler{svc: svc}
}

func (h *ExaminerHandler) GetList(c *fiber.Ctx) error {
	var filter dto.ExaminerFilter
	if err := c.QueryParser(&filter); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_QUERY", "Tham số không hợp lệ", nil)
	}

	res, err := h.svc.GetList(&filter)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Lấy danh sách cán bộ coi thi thành công", res)
}

func (h *ExaminerHandler) GetDetail(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID không hợp lệ", nil)
	}

	res, err := h.svc.GetDetail(uint(id))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusNotFound, "NOT_FOUND", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Lấy thông tin cán bộ coi thi thành công", res)
}

func (h *ExaminerHandler) Create(c *fiber.Ctx) error {
	var req dto.ExaminerCreateReq
	if err := c.BodyParser(&req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_BODY", "Dữ liệu không hợp lệ", nil)
	}

	res, err := h.svc.Create(&req)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "CREATE_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Tạo cán bộ coi thi thành công", res)
}

func (h *ExaminerHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID không hợp lệ", nil)
	}

	var req dto.ExaminerUpdateReq
	if err := c.BodyParser(&req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_BODY", "Dữ liệu không hợp lệ", nil)
	}

	res, err := h.svc.Update(uint(id), &req)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "UPDATE_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Cập nhật cán bộ coi thi thành công", res)
}

func (h *ExaminerHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID không hợp lệ", nil)
	}

	if err := h.svc.Delete(uint(id)); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "DELETE_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Xóa cán bộ coi thi thành công", nil)
}
