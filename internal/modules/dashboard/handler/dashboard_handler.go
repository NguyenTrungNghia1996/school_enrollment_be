package handler

import (
	"go_be_enrollment/internal/common/httpresponse"
	"go_be_enrollment/internal/modules/dashboard/dto"
	"go_be_enrollment/internal/modules/dashboard/service"

	"github.com/gofiber/fiber/v2"
)

type DashboardHandler struct {
	svc service.DashboardService
}

func NewDashboardHandler(svc service.DashboardService) *DashboardHandler {
	return &DashboardHandler{svc: svc}
}

func (h *DashboardHandler) GetSummary(c *fiber.Ctx) error {
	var filter dto.DashboardSummaryFilter
	if err := c.QueryParser(&filter); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_QUERY", "Tham số không hợp lệ", nil)
	}

	res, err := h.svc.GetSummary(&filter)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Lấy thông tin tổng quan thành công", res)
}
