package subject

import (
	"go_be_enrollment/internal/config"
	"go_be_enrollment/internal/middleware"
	"go_be_enrollment/internal/modules/subject/handler"
	"go_be_enrollment/internal/modules/subject/repository"
	"go_be_enrollment/internal/modules/subject/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterSubjectRoutes(api fiber.Router, db *gorm.DB, cfg *config.Config) {
	repo := repository.NewSubjectRepository(db)
	svc := service.NewSubjectService(repo)
	h := handler.NewSubjectHandler(svc)

	// Admin routes
	admin := api.Group("/admin/subjects")
	admin.Use(middleware.AdminAuth(cfg))
	admin.Get("/", h.GetList)
	admin.Get("/:id", h.GetDetail)
	admin.Post("/", h.Create)
	admin.Put("/:id", h.Update)
	admin.Patch("/:id/status", h.UpdateStatus)

	// Public routes
	public := api.Group("/public/subjects")
	public.Get("/", h.GetPublicList)
}
