package adminuser

import (
	"go_be_enrollment/internal/config"
	"go_be_enrollment/internal/middleware"
	"go_be_enrollment/internal/modules/adminuser/handler"
	"go_be_enrollment/internal/modules/adminuser/repository"
	"go_be_enrollment/internal/modules/adminuser/service"

	adminauth_repo "go_be_enrollment/internal/modules/adminauth/repository"
	adminauth_service "go_be_enrollment/internal/modules/adminauth/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAdminUserRoutes(router fiber.Router, db *gorm.DB, cfg *config.Config) {
	repo := repository.NewAdminUserRepository(db)
	svc := service.NewAdminUserService(repo)
	hdl := handler.NewAdminUserHandler(svc)

	permRepo := adminauth_repo.NewPermissionRepository(db)
	permSvc := adminauth_service.NewPermissionService(permRepo)

	adminUserGroup := router.Group("/admin/admin-users")

	// Protected endpoints (Requires Admin Login)
	adminUserGroup.Use(middleware.AdminAuth(cfg))

	// TODO: Thêm middleware PermissionGuard vào các route này.
	// Ví dụ: middleware.PermissionGuard(permSvc, "admin_users_menu", 0) cho quyền xem danh sách (READ bit)
	//        middleware.PermissionGuard(permSvc, "admin_users_menu", 1) cho quyền tạo (CREATE bit)
	_ = permSvc // Tạm bỏ qua unused variable cho việc import

	adminUserGroup.Get("/", hdl.GetList)
	adminUserGroup.Get("/:id", hdl.GetDetail)
	adminUserGroup.Post("/", hdl.Create)
	adminUserGroup.Put("/:id", hdl.Update)
	adminUserGroup.Patch("/:id/status", hdl.UpdateStatus)
	adminUserGroup.Patch("/:id/reset-password", hdl.ResetPassword)
	
	adminUserGroup.Get("/:id/role-groups", hdl.GetRoleGroups)
	adminUserGroup.Put("/:id/role-groups", hdl.UpdateRoleGroups)
}
