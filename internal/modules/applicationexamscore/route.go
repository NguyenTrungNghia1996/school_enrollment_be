package applicationexamscore

import (
	"go_be_enrollment/internal/config"
	"go_be_enrollment/internal/middleware"
	"go_be_enrollment/internal/modules/applicationexamscore/handler"
	"go_be_enrollment/internal/modules/applicationexamscore/repository"
	"go_be_enrollment/internal/modules/applicationexamscore/service"
	app_repo "go_be_enrollment/internal/modules/application/repository"
	subject_repo "go_be_enrollment/internal/modules/subject/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterApplicationExamScoreRoutes(api fiber.Router, db *gorm.DB, cfg *config.Config) {
	repo := repository.NewApplicationExamScoreRepository(db)
	appRepo := app_repo.NewApplicationRepository(db)
	subjectRepo := subject_repo.NewSubjectRepository(db)
	svc := service.NewApplicationExamScoreService(repo, appRepo, subjectRepo)
	h := handler.NewApplicationExamScoreHandler(svc)

	admin := api.Group("/admin/applications/:id/scores")
	admin.Use(middleware.AdminAuth(cfg))
	admin.Get("/", h.GetList)
	admin.Put("/", h.Replace)
}
