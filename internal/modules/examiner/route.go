package examiner

import (
	"go_be_enrollment/internal/config"
	"go_be_enrollment/internal/middleware"
	"go_be_enrollment/internal/modules/examiner/handler"
	"go_be_enrollment/internal/modules/examiner/repository"
	"go_be_enrollment/internal/modules/examiner/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterExaminerRoutes(api fiber.Router, db *gorm.DB, cfg *config.Config) {
	repo := repository.NewExaminerRepository(db)
	svc := service.NewExaminerService(repo)
	h := handler.NewExaminerHandler(svc)

	admin := api.Group("/admin/examiners")
	admin.Use(middleware.AdminAuth(cfg))
	admin.Get("/", h.GetList)
	admin.Get("/:id", h.GetDetail)
	admin.Post("/", h.Create)
	admin.Put("/:id", h.Update)
	admin.Delete("/:id", h.Delete)
}
