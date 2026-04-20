package dashboard

import (
	"go_be_enrollment/internal/config"
	"go_be_enrollment/internal/middleware"
	"go_be_enrollment/internal/modules/dashboard/handler"
	"go_be_enrollment/internal/modules/dashboard/repository"
	"go_be_enrollment/internal/modules/dashboard/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterDashboardRoutes(api fiber.Router, db *gorm.DB, cfg *config.Config) {
	repo := repository.NewDashboardRepository(db)
	svc := service.NewDashboardService(repo)
	h := handler.NewDashboardHandler(svc)

	admin := api.Group("/admin/dashboard")
	admin.Use(middleware.AdminAuth(cfg))
	admin.Get("/summary", h.GetSummary)
}
