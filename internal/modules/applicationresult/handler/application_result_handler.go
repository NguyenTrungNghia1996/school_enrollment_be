package handler

import (
	"strconv"

	"go_be_enrollment/internal/common/httpresponse"
	"go_be_enrollment/internal/modules/applicationresult/dto"
	"go_be_enrollment/internal/modules/applicationresult/service"

	"github.com/gofiber/fiber/v2"
)

type ApplicationResultHandler struct {
	svc service.ApplicationResultService
}

func NewApplicationResultHandler(svc service.ApplicationResultService) *ApplicationResultHandler {
	return &ApplicationResultHandler{svc: svc}
}

func (h *ApplicationResultHandler) GetAdminResult(c *fiber.Ctx) error {
	appID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID hồ sơ không hợp lệ", nil)
	}

	res, err := h.svc.GetAdminResult(uint(appID))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "GET_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Lấy kết quả ứng viên thành công", res)
}

func (h *ApplicationResultHandler) UpdateManual(c *fiber.Ctx) error {
	appID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID hồ sơ không hợp lệ", nil)
	}

	var req dto.ApplicationResultUpdateReq
	if err := c.BodyParser(&req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_BODY", "Dữ liệu không hợp lệ", nil)
	}

	res, err := h.svc.UpdateManual(uint(appID), &req)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "UPDATE_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Cập nhật kết quả thành công", res)
}

func (h *ApplicationResultHandler) RecalculateSingle(c *fiber.Ctx) error {
	appID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID hồ sơ không hợp lệ", nil)
	}

	res, err := h.svc.RecalculateSingle(uint(appID))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "RECALCULATE_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Tính toán kết quả thành công", res)
}

func (h *ApplicationResultHandler) RecalculateRankingByPeriod(c *fiber.Ctx) error {
	periodID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID kỳ tuyển sinh không hợp lệ", nil)
	}

	if err := h.svc.RecalculateRankingByPeriod(uint(periodID)); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "RECALCULATE_RANKING_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Tính hạng tất cả ứng viên thành công", nil)
}

func (h *ApplicationResultHandler) GetUserResult(c *fiber.Ctx) error {
	appID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID hồ sơ không hợp lệ", nil)
	}
	userID := c.Locals("user_id").(uint)

	res, err := h.svc.GetUserResult(uint(appID), userID)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "GET_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Lấy kết quả ứng viên thành công", res)
}
