package applicationresult

import (
	"go_be_enrollment/internal/config"
	"go_be_enrollment/internal/middleware"
	"go_be_enrollment/internal/modules/applicationresult/handler"
	"go_be_enrollment/internal/modules/applicationresult/repository"
	"go_be_enrollment/internal/modules/applicationresult/service"
	app_repo "go_be_enrollment/internal/modules/application/repository"
	score_repo "go_be_enrollment/internal/modules/applicationexamscore/repository"
	period_repo "go_be_enrollment/internal/modules/admissionperiod/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterApplicationResultRoutes(api fiber.Router, db *gorm.DB, cfg *config.Config) {
	repo := repository.NewApplicationResultRepository(db)
	appRepo := app_repo.NewApplicationRepository(db)
	scoreRepo := score_repo.NewApplicationExamScoreRepository(db)
	periodRepo := period_repo.NewAdmissionPeriodRepository(db)
	svc := service.NewApplicationResultService(repo, appRepo, scoreRepo, periodRepo)
	h := handler.NewApplicationResultHandler(svc)

	admin := api.Group("/admin/applications/:id/result")
	admin.Use(middleware.AdminAuth(cfg))
	admin.Get("/", h.GetAdminResult)
	admin.Put("/", h.UpdateManual)
	admin.Post("/recalculate", h.RecalculateSingle)

	adminPeriod := api.Group("/admin/admission-periods/:id/results/recalculate-ranking")
	adminPeriod.Use(middleware.AdminAuth(cfg))
	adminPeriod.Post("/", h.RecalculateRankingByPeriod)

	user := api.Group("/user/me/applications/:id/result")
	user.Use(middleware.UserAuth(cfg))
	user.Get("/", h.GetUserResult)
}
