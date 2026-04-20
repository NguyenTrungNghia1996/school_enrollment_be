package handler

import (
	"strconv"

	"go_be_enrollment/internal/common/httpresponse"
	"go_be_enrollment/internal/modules/examinerassignment/dto"
	"go_be_enrollment/internal/modules/examinerassignment/service"

	"github.com/gofiber/fiber/v2"
)

type ExaminerAssignmentHandler struct {
	svc service.ExaminerAssignmentService
}

func NewExaminerAssignmentHandler(svc service.ExaminerAssignmentService) *ExaminerAssignmentHandler {
	return &ExaminerAssignmentHandler{svc: svc}
}

func (h *ExaminerAssignmentHandler) GetList(c *fiber.Ctx) error {
	var filter dto.ExaminerAssignmentFilter
	if err := c.QueryParser(&filter); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_QUERY", "Tham số không hợp lệ", nil)
	}

	res, err := h.svc.GetList(&filter)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Lấy danh sách phân công coi thi thành công", res)
}

func (h *ExaminerAssignmentHandler) Create(c *fiber.Ctx) error {
	var req dto.ExaminerAssignmentCreateReq
	if err := c.BodyParser(&req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_BODY", "Dữ liệu không hợp lệ", nil)
	}

	res, err := h.svc.Create(&req)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "CREATE_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Phân công coi thi thành công", res)
}

func (h *ExaminerAssignmentHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID không hợp lệ", nil)
	}

	var req dto.ExaminerAssignmentUpdateReq
	if err := c.BodyParser(&req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_BODY", "Dữ liệu không hợp lệ", nil)
	}

	res, err := h.svc.Update(uint(id), &req)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "UPDATE_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Cập nhật phân công coi thi thành công", res)
}

func (h *ExaminerAssignmentHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID không hợp lệ", nil)
	}

	if err := h.svc.Delete(uint(id)); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "DELETE_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Xóa phân công coi thi thành công", nil)
}

func (h *ExaminerAssignmentHandler) GetListByRoomID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID không hợp lệ", nil)
	}

	res, err := h.svc.GetListByRoomID(uint(id))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "GET_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Lấy danh sách cán bộ trong phòng thi thành công", res)
}
