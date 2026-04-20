package database

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"go_be_enrollment/internal/config"
	zaplogger "go_be_enrollment/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbInstance *gorm.DB

// ConnectMySQL initializes the MySQL connection using GORM and creates a connection pool
func ConnectMySQL(cfg *config.Config) (*gorm.DB, error) {
	// Parse Timezone correctly for DSN
	tz := url.QueryEscape(cfg.AppTimezone)
	
	// Create DSN
	// Note: You can customize MYSQL_PARAMS in env, but by default we make sure it embeds loc properly
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		cfg.MySQLUser,
		cfg.MySQLPassword,
		cfg.MySQLHost,
		cfg.MySQLPort,
		cfg.MySQLDB,
		tz,
	)

	// Set GORM Logger Level based on Environment
	gormLogLevel := logger.Info
	if cfg.AppEnv == "production" {
		gormLogLevel = logger.Warn
	}

	gormLogger := logger.Default.LogMode(gormLogLevel)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		zaplogger.Log.Error("Failed to connect to MySQL database: " + err.Error())
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		zaplogger.Log.Error("Failed to extract sql.DB from gorm: " + err.Error())
		return nil, err
	}

	// Setup Connection Pool
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	zaplogger.Log.Info("Connected to MySQL database via GORM with pooling enabled")
	
	dbInstance = db
	return db, nil
}

// GetDB returns the initialized *gorm.DB globally, can be bypassed by Dependency Injection in Repository
func GetDB() *gorm.DB {
	return dbInstance
}

// Close explicitly closes the database pool
func Close() error {
	if dbInstance != nil {
		sqlDB, err := dbInstance.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}

// Ping performs a lightweight query to check if the DB is responding
func Ping(ctx context.Context) error {
	if dbInstance == nil {
		return fmt.Errorf("database not initialized")
	}
	sqlDB, err := dbInstance.DB()
	if err != nil {
		return err
	}
	return sqlDB.PingContext(ctx)
}
