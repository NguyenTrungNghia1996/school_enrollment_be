package main

import (
	"fmt"

	"go_be_enrollment/internal/config"
	"go_be_enrollment/internal/database"
	"go_be_enrollment/internal/middleware"
	"go_be_enrollment/internal/modules/health"
	"go_be_enrollment/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logger
	logger.InitLogger(cfg.AppEnv)
	defer logger.Sync()

	// Set Gin mode
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize database
	db, err := database.ConnectMySQL(cfg)
	if err != nil {
		logger.Log.Fatal("Could not initialize database connection", zap.Error(err))
	}
	// Prevent "declared and not used" error for now. `db` will be injected into repositories.
	_ = db

	// Create Gin router without default middlewares
	r := gin.New()

	// Apply Middlewares
	r.Use(middleware.RequestLogger())
	r.Use(middleware.Recovery())
	r.Use(middleware.CORS())

	// Register Routes
	api := r.Group("/api/v1")
	{
		health.RegisterRoutes(api)
	}

	// Start server
	addr := fmt.Sprintf(":%s", cfg.AppPort)
	logger.Log.Info(fmt.Sprintf("Server is starting on port %s...", cfg.AppPort))
	if err := r.Run(addr); err != nil {
		logger.Log.Fatal("Failed to start server", zap.Error(err))
	}
}
