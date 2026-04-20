package wardunit

import (
	"go_be_enrollment/internal/config"
	"go_be_enrollment/internal/middleware"
	"go_be_enrollment/internal/modules/wardunit/handler"
	"go_be_enrollment/internal/modules/wardunit/repository"
	"go_be_enrollment/internal/modules/wardunit/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterWardUnitRoutes(router fiber.Router, db *gorm.DB, cfg *config.Config) {
	repo := repository.NewWardUnitRepository(db)
	svc := service.NewWardUnitService(repo)
	hdl := handler.NewWardUnitHandler(svc)

	// Admin routes
	adminGroup := router.Group("/admin/ward-units")
	adminGroup.Use(middleware.AdminAuth(cfg))
	adminGroup.Get("/", hdl.GetList)
	adminGroup.Get("/:id", hdl.GetDetail)
	adminGroup.Post("/", hdl.Create)
	adminGroup.Put("/:id", hdl.Update)
	adminGroup.Patch("/:id/status", hdl.UpdateStatus)

	// Public routes
	publicGroup := router.Group("/public/ward-units")
	publicGroup.Get("/", hdl.GetPublicList)
}
