package handler

import (
	"strconv"

	"go_be_enrollment/internal/common/httpresponse"
	"go_be_enrollment/internal/modules/subject/dto"
	"go_be_enrollment/internal/modules/subject/service"

	"github.com/gofiber/fiber/v2"
)

type SubjectHandler struct {
	svc service.SubjectService
}

func NewSubjectHandler(svc service.SubjectService) *SubjectHandler {
	return &SubjectHandler{svc: svc}
}

func (h *SubjectHandler) GetList(c *fiber.Ctx) error {
	var filter dto.SubjectFilter
	if err := c.QueryParser(&filter); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_QUERY", "Tham số không hợp lệ", nil)
	}

	res, err := h.svc.GetList(&filter)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Lấy danh sách môn học thành công", res)
}

func (h *SubjectHandler) GetDetail(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID không hợp lệ", nil)
	}

	res, err := h.svc.GetDetail(uint(id))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusNotFound, "NOT_FOUND", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Lấy thông tin môn học thành công", res)
}

func (h *SubjectHandler) Create(c *fiber.Ctx) error {
	var req dto.SubjectCreateReq
	if err := c.BodyParser(&req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_BODY", "Dữ liệu không hợp lệ", nil)
	}

	res, err := h.svc.Create(&req)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "CREATE_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Tạo môn học thành công", res)
}

func (h *SubjectHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID không hợp lệ", nil)
	}

	var req dto.SubjectUpdateReq
	if err := c.BodyParser(&req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_BODY", "Dữ liệu không hợp lệ", nil)
	}

	res, err := h.svc.Update(uint(id), &req)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "UPDATE_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Cập nhật môn học thành công", res)
}

func (h *SubjectHandler) UpdateStatus(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID không hợp lệ", nil)
	}

	var req dto.SubjectStatusReq
	if err := c.BodyParser(&req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_BODY", "Dữ liệu không hợp lệ", nil)
	}

	if err := h.svc.UpdateStatus(uint(id), &req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "UPDATE_STATUS_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Cập nhật trạng thái thành công", nil)
}

func (h *SubjectHandler) GetPublicList(c *fiber.Ctx) error {
	res, err := h.svc.GetActiveList()
	if err != nil {
		return httpresponse.Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", err.Error(), nil)
	}
	return httpresponse.Success(c, fiber.StatusOK, "Lấy danh sách môn học thành công", res)
}
