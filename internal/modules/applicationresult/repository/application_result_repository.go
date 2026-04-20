package repository

import (
	"go_be_enrollment/internal/modules/applicationresult/entity"

	"gorm.io/gorm"
)

type ApplicationResultRepository interface {
	FindByApplicationID(appID uint) (*entity.ApplicationResult, error)
	GetByAdmissionPeriod(periodID uint) ([]entity.ApplicationResult, error)
	Save(e *entity.ApplicationResult) error
	UpdateBatch(results []entity.ApplicationResult) error
}

type applicationResultRepository struct {
	db *gorm.DB
}

func NewApplicationResultRepository(db *gorm.DB) ApplicationResultRepository {
	return &applicationResultRepository{db: db}
}

func (r *applicationResultRepository) FindByApplicationID(appID uint) (*entity.ApplicationResult, error) {
	var a entity.ApplicationResult
	if err := r.db.Where("application_id = ?", appID).First(&a).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Trả về nil nếu k tìm thấy để có thể insert mới
		}
		return nil, err
	}
	return &a, nil
}

func (r *applicationResultRepository) GetByAdmissionPeriod(periodID uint) ([]entity.ApplicationResult, error) {
	var list []entity.ApplicationResult
	// Join với applications để lấy theo admission_period_id và Order theo final_total_score DESC
	if err := r.db.Joins("Application").Where("Application.admission_period_id = ?", periodID).Order("final_total_score DESC").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *applicationResultRepository) Save(e *entity.ApplicationResult) error {
	if e.ID == 0 {
		return r.db.Create(e).Error
	}
	return r.db.Save(e).Error
}

func (r *applicationResultRepository) UpdateBatch(results []entity.ApplicationResult) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, res := range results {
			if err := tx.Save(&res).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
