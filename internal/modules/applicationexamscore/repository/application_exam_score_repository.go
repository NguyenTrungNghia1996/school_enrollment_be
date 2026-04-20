package repository

import (
	"go_be_enrollment/internal/modules/applicationexamscore/entity"

	"gorm.io/gorm"
)

type ApplicationExamScoreRepository interface {
	GetByApplicationID(appID uint) ([]entity.ApplicationExamScore, error)
	ReplaceScores(appID uint, scores []entity.ApplicationExamScore) error
}

type applicationExamScoreRepository struct {
	db *gorm.DB
}

func NewApplicationExamScoreRepository(db *gorm.DB) ApplicationExamScoreRepository {
	return &applicationExamScoreRepository{db: db}
}

func (r *applicationExamScoreRepository) GetByApplicationID(appID uint) ([]entity.ApplicationExamScore, error) {
	var list []entity.ApplicationExamScore
	if err := r.db.Preload("Subject").Where("application_id = ?", appID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *applicationExamScoreRepository) ReplaceScores(appID uint, scores []entity.ApplicationExamScore) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("application_id = ?", appID).Delete(&entity.ApplicationExamScore{}).Error; err != nil {
			return err
		}

		if len(scores) > 0 {
			if err := tx.Create(&scores).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
