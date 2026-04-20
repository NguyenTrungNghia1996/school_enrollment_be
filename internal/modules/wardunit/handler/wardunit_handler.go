package handler

import (
	"strconv"

	"go_be_enrollment/internal/common/httpresponse"
	"go_be_enrollment/internal/modules/wardunit/dto"
	"go_be_enrollment/internal/modules/wardunit/service"

	"github.com/gofiber/fiber/v2"
)

type WardUnitHandler struct {
	svc service.WardUnitService
}

func NewWardUnitHandler(svc service.WardUnitService) *WardUnitHandler {
	return &WardUnitHandler{svc: svc}
}

func (h *WardUnitHandler) GetList(c *fiber.Ctx) error {
	var filter dto.WardUnitFilter
	if err := c.QueryParser(&filter); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_QUERY", "Tham số không hợp lệ", nil)
	}

	res, err := h.svc.GetList(&filter)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Lấy danh sách ward unit thành công", res)
}

func (h *WardUnitHandler) GetDetail(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID không hợp lệ", nil)
	}

	res, err := h.svc.GetDetail(uint(id))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusNotFound, "NOT_FOUND", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Lấy thông tin ward unit thành công", res)
}

func (h *WardUnitHandler) Create(c *fiber.Ctx) error {
	var req dto.WardUnitCreateReq
	if err := c.BodyParser(&req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_BODY", "Dữ liệu không hợp lệ", nil)
	}

	res, err := h.svc.Create(&req)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "CREATE_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Tạo ward unit thành công", res)
}

func (h *WardUnitHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID không hợp lệ", nil)
	}

	var req dto.WardUnitUpdateReq
	if err := c.BodyParser(&req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_BODY", "Dữ liệu không hợp lệ", nil)
	}

	res, err := h.svc.Update(uint(id), &req)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "UPDATE_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Cập nhật ward unit thành công", res)
}

func (h *WardUnitHandler) UpdateStatus(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID không hợp lệ", nil)
	}

	var req dto.WardUnitStatusReq
	if err := c.BodyParser(&req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_BODY", "Dữ liệu không hợp lệ", nil)
	}

	if err := h.svc.UpdateStatus(uint(id), &req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "UPDATE_STATUS_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Cập nhật trạng thái thành công", nil)
}

func (h *WardUnitHandler) GetPublicList(c *fiber.Ctx) error {
	provinceIDStr := c.Query("province_id")
	if provinceIDStr == "" {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_QUERY", "Thiếu province_id", nil)
	}

	provinceID, err := strconv.Atoi(provinceIDStr)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_QUERY", "province_id không hợp lệ", nil)
	}

	res, err := h.svc.GetActiveListByProvince(uint(provinceID))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "GET_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Lấy danh sách ward unit thành công", res)
}
