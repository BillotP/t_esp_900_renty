package resolvers

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm/clause"
)

func (r *QueryResolver) GetCompanies(ctx context.Context) ([]*models.Company, error) {
	var (
		companies []models.Company

		err error
	)

	if err = r.DB.Preload(clause.Associations).Find(&companies).Error; err == nil {
		var companiesfmt []*models.Company

		for i := range companies {
			companiesfmt = append(companiesfmt, &companies[i])
		}
		return companiesfmt, nil
	}
	lib.LogError("mutation/GetCompanies", err.Error())
	return nil, err
}
