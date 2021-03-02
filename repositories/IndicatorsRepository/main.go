package IndicatorsRepository

import (
	"ftp/models/IndicatorModel"
	"gateway-api/database"

	"github.com/jinzhu/gorm"
)
func NewIndicatorsRepository(db *gorm.DB) *IndicatorsRepository {
	return &IndicatorsRepository{db:database.GetDB()}
}

type IndicatorsRepository struct {
	db *gorm.DB
}

func (repo IndicatorsRepository) GetById(id int64)(model IndicatorModel.Indicator, err error) {
	result := repo.db.Table("indicators").Where("id = ?", id).Scan(&model)
	err = result.Error
	return
}