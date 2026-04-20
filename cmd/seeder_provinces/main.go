package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"gorm.io/gorm/clause"

	"go_be_enrollment/internal/config"
	"go_be_enrollment/internal/database"
	province_entity "go_be_enrollment/internal/modules/province/entity"
	wardunit_entity "go_be_enrollment/internal/modules/wardunit/entity"
	"go_be_enrollment/pkg/logger"

	"go.uber.org/zap"
)

type ApiWard struct {
	Name         string `json:"name"`
	Code         int    `json:"code"`
	DivisionType string `json:"division_type"`
}

type ApiProvince struct {
	Name         string    `json:"name"`
	Code         int       `json:"code"`
	DivisionType string    `json:"division_type"`
	Wards        []ApiWard `json:"wards"`
}

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

	if err := db.AutoMigrate(&province_entity.Province{}, &wardunit_entity.WardUnit{}); err != nil {
		logger.Log.Fatal("Failed to migrate province/wardunit tables", zap.Error(err))
	}

	logger.Log.Info("Fetching data from open-api.vn API v2 ...")
	client := http.Client{Timeout: 60 * time.Second}
	resp, err := client.Get("https://provinces.open-api.vn/api/v2/?depth=2")
	if err != nil {
		logger.Log.Fatal("Failed to fetch API", zap.Error(err))
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Fatal("Failed to read body", zap.Error(err))
	}

	var data []ApiProvince
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		logger.Log.Fatal("Failed to unmarshal JSON", zap.Error(err))
	}

	logger.Log.Info(fmt.Sprintf("Fetched %d provinces. Starting import...", len(data)))

	tx := db.Begin()
	if tx.Error != nil {
		logger.Log.Fatal("Failed to start transaction", zap.Error(tx.Error))
	}

	// Disable foreign key checks for truncation
	tx.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	tx.Exec("TRUNCATE TABLE ward_units;")
	tx.Exec("TRUNCATE TABLE provinces;")
	tx.Exec("SET FOREIGN_KEY_CHECKS = 1;")

	for i, p := range data {
		provCode := fmt.Sprintf("%02d", p.Code)

		prov := province_entity.Province{
			Code:     provCode,
			Name:     p.Name,
			IsActive: true,
		}

		if err := tx.Create(&prov).Error; err != nil {
			logger.Log.Error("Failed to insert province", zap.String("name", p.Name), zap.Error(err))
			tx.Rollback()
			return
		}

		// Insert wards for this province
		// Using batch insert to be fast
		var wardEntities []wardunit_entity.WardUnit
		seenWardNames := make(map[string]bool)

		for _, w := range p.Wards {
			cleanName := strings.TrimSpace(strings.ToLower(w.Name))
			if seenWardNames[cleanName] {
				// Prevent database duplicate unique error on idx_province_name restriction
				logger.Log.Warn("Skipping duplicate ward name in same province", zap.String("name", w.Name))
				continue
			}
			seenWardNames[cleanName] = true

			wardCode := fmt.Sprintf("%05d", w.Code)
			
			// Map API division type to ENUM('Ward', 'Commune', 'SpecialZone')
			unitType := "Ward"
			if w.DivisionType == "xã" || w.DivisionType == "thị trấn" {
				unitType = "Commune"
			} else if w.DivisionType == "đặc khu" {
				unitType = "SpecialZone"
			}

			wardEntities = append(wardEntities, wardunit_entity.WardUnit{
				ProvinceID: prov.ID,
				Code:       wardCode,
				Name:       w.Name,
				UnitType:   unitType,
				IsActive:   true,
			})
		}

		if len(wardEntities) > 0 {
			// Batch insert with OnConflict DoNothing to bypass unique constraint dynamically
			if err := tx.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(wardEntities, 100).Error; err != nil {
				logger.Log.Error("Failed to insert wards", zap.String("province", p.Name), zap.Error(err))
				tx.Rollback()
				return
			}
		}

		if (i+1)%10 == 0 {
			logger.Log.Info(fmt.Sprintf("Processed %d/%d provinces...", i+1, len(data)))
		}
	}

	if err := tx.Commit().Error; err != nil {
		logger.Log.Fatal("Failed to commit transaction", zap.Error(err))
	}

	logger.Log.Info(fmt.Sprintf("Import successful! Imported %d provinces.", len(data)))
}
