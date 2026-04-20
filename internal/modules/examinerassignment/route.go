package examinerassignment

import (
	"go_be_enrollment/internal/config"
	"go_be_enrollment/internal/middleware"
	"go_be_enrollment/internal/modules/examinerassignment/handler"
	"go_be_enrollment/internal/modules/examinerassignment/repository"
	"go_be_enrollment/internal/modules/examinerassignment/service"
	examiner_repo "go_be_enrollment/internal/modules/examiner/repository"
	room_repo "go_be_enrollment/internal/modules/examroom/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterExaminerAssignmentRoutes(api fiber.Router, db *gorm.DB, cfg *config.Config) {
	repo := repository.NewExaminerAssignmentRepository(db)
	examinerRepo := examiner_repo.NewExaminerRepository(db)
	roomRepo := room_repo.NewExamRoomRepository(db)
	svc := service.NewExaminerAssignmentService(repo, examinerRepo, roomRepo)
	h := handler.NewExaminerAssignmentHandler(svc)

	admin := api.Group("/admin/examiner-assignments")
	admin.Use(middleware.AdminAuth(cfg))
	admin.Get("/", h.GetList)
	admin.Post("/", h.Create)
	admin.Put("/:id", h.Update)
	admin.Delete("/:id", h.Delete)
	
	// GET /api/v1/admin/exam-rooms/:id/examiners
	adminRoom := api.Group("/admin/exam-rooms/:id/examiners")
	adminRoom.Use(middleware.AdminAuth(cfg))
	adminRoom.Get("/", h.GetListByRoomID)
}
