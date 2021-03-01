package resolvers

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm/clause"
)

func (r *QueryResolver) GetCompany(ctx context.Context, id int64) (*models.Company, error) {
	var (
		company models.Company

		err error
	)

	if err = r.DB.Preload(clause.Associations).Where("id = ?", id).First(&company).Error; err == nil {
		return &company, nil
	}
	lib.LogError("mutation/GetCompany", err.Error())
	return nil, err
}
