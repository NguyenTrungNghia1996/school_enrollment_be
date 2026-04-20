package examroomassignment

import (
	"go_be_enrollment/internal/config"
	"go_be_enrollment/internal/middleware"
	"go_be_enrollment/internal/modules/examroomassignment/handler"
	"go_be_enrollment/internal/modules/examroomassignment/repository"
	"go_be_enrollment/internal/modules/examroomassignment/service"
	app_repo "go_be_enrollment/internal/modules/application/repository"
	room_repo "go_be_enrollment/internal/modules/examroom/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterExamRoomAssignmentRoutes(api fiber.Router, db *gorm.DB, cfg *config.Config) {
	repo := repository.NewExamRoomAssignmentRepository(db)
	appRepo := app_repo.NewApplicationRepository(db)
	roomRepo := room_repo.NewExamRoomRepository(db)
	svc := service.NewExamRoomAssignmentService(repo, appRepo, roomRepo)
	h := handler.NewExamRoomAssignmentHandler(svc)

	admin := api.Group("/admin/exam-room-assignments")
	admin.Use(middleware.AdminAuth(cfg))
	admin.Get("/", h.GetList)
	admin.Post("/", h.Create)
	admin.Put("/:id", h.Update)
	admin.Delete("/:id", h.Delete)
	
	// GET /api/v1/admin/exam-rooms/:id/applications
	adminRoom := api.Group("/admin/exam-rooms/:id/applications")
	adminRoom.Use(middleware.AdminAuth(cfg))
	adminRoom.Get("/", h.GetListByRoomID)
}
