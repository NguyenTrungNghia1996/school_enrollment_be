package main

import (
	"log"

	"go_be_enrollment/internal/config"
	"go_be_enrollment/internal/database"
	"go_be_enrollment/internal/modules/adminauth/entity"
	"go_be_enrollment/pkg/logger"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger.InitLogger(cfg.AppEnv)
	defer logger.Sync()

	db, err := database.ConnectMySQL(cfg)
	if err != nil {
		logger.Log.Fatal("Could not connect to database", zap.Error(err))
	}
	defer database.Close()

	// Ensure table exists
	if err := db.AutoMigrate(&entity.AdminUser{}); err != nil {
		logger.Log.Fatal("Failed to migrate admin user table", zap.Error(err))
	}

	// Setup user credentials
	username := "admin"
	password := "adminpwd"
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.Fatal("Failed to hash password", zap.Error(err))
	}

	admin := entity.AdminUser{
		Username:     username,
		PasswordHash: string(hashedBytes),
		FullName:     "System Administrator",
		IsSuperAdmin: true,
		IsActive:     true,
	}

	// Check if table is empty
	var count int64
	if countErr := db.Model(&entity.AdminUser{}).Count(&count).Error; countErr != nil {
		logger.Log.Fatal("Failed to check admin users count", zap.Error(countErr))
	}

	if count > 0 {
		logger.Log.Info("Admin users table is not empty. Skipping seeding.")
		return
	}

	if createErr := db.Create(&admin).Error; createErr != nil {
		logger.Log.Fatal("Failed to seed admin", zap.Error(createErr))
	}

	logger.Log.Info("Successfully seeded super admin user account.")
}
