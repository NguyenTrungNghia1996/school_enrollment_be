package handler

import (
	"strconv"

	"go_be_enrollment/internal/common/httpresponse"
	"go_be_enrollment/internal/modules/applicationexamscore/dto"
	"go_be_enrollment/internal/modules/applicationexamscore/service"

	"github.com/gofiber/fiber/v2"
)

type ApplicationExamScoreHandler struct {
	svc service.ApplicationExamScoreService
}

func NewApplicationExamScoreHandler(svc service.ApplicationExamScoreService) *ApplicationExamScoreHandler {
	return &ApplicationExamScoreHandler{svc: svc}
}

func (h *ApplicationExamScoreHandler) GetList(c *fiber.Ctx) error {
	appID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID hồ sơ không hợp lệ", nil)
	}

	res, err := h.svc.GetByApplicationID(uint(appID))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "GET_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Lấy danh sách điểm thành công", res)
}

func (h *ApplicationExamScoreHandler) Replace(c *fiber.Ctx) error {
	appID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_PARAM", "ID hồ sơ không hợp lệ", nil)
	}

	var req dto.ReplaceExamScoresReq
	if err := c.BodyParser(&req); err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "INVALID_BODY", "Dữ liệu không hợp lệ", nil)
	}

	res, err := h.svc.ReplaceScores(uint(appID), &req)
	if err != nil {
		return httpresponse.Error(c, fiber.StatusBadRequest, "REPLACE_FAILED", err.Error(), nil)
	}

	return httpresponse.Success(c, fiber.StatusOK, "Cập nhật điểm thành công", res)
}
