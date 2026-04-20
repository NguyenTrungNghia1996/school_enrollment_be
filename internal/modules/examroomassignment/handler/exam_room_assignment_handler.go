package handler

import (
	"strconv"

	"go_be_enrollment/internal/common/httpresponse"
	"go_be_enrollment/internal/modules/examroomassignment/dto"
	"go_be_enrollment/internal/modules/examroomassignment/service"

	"github.com/gofiber/fiber/v2"
)

type ExamRoomAssignmentHandler struct {
	svc service.ExamRoomAssignmentService
}

func NewExamRoomAssignmentHandler(svc service.ExamRoomAssignmentService) *ExamRoomAssignmentHandler {
	return &ExamRoomAssignmentHandler{svc: svc}
}

func (h *ExamRoomAssignmentHandler) GetList(c *fiber.Ctx) error {
	var filter dto.ExamRoomAssignmentFilter
	if err := c.QueryParser(&filter); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_QUERY", "Tham số không hợp lệ", nil)
	}

	res, err := h.svc.GetList(&filter)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Lấy danh sách phân phòng thành công", res)
}

func (h *ExamRoomAssignmentHandler) Create(c *fiber.Ctx) error {
	var req dto.ExamRoomAssignmentCreateReq
	if err := c.BodyParser(&req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_BODY", "Dữ liệu không hợp lệ", nil)
	}

	res, err := h.svc.Create(&req)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "CREATE_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Phân phòng thành công", res)
}

func (h *ExamRoomAssignmentHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID không hợp lệ", nil)
	}

	var req dto.ExamRoomAssignmentUpdateReq
	if err := c.BodyParser(&req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_BODY", "Dữ liệu không hợp lệ", nil)
	}

	res, err := h.svc.Update(uint(id), &req)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "UPDATE_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Cập nhật phân phòng thành công", res)
}

func (h *ExamRoomAssignmentHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID không hợp lệ", nil)
	}

	if err := h.svc.Delete(uint(id)); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "DELETE_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Xóa phân phòng thành công", nil)
}

func (h *ExamRoomAssignmentHandler) GetListByRoomID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID không hợp lệ", nil)
	}

	res, err := h.svc.GetListByRoomID(uint(id))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "GET_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Lấy danh sách thí sinh trong phòng thành công", res)
}
