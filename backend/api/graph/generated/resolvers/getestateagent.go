package resolvers

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm/clause"
)

func (r *QueryResolver) EstateAgent(ctx context.Context, id int64) (*models.EstateAgent, error) {
	var (
		estateAgent models.EstateAgent

		err error
	)

	if err = r.DB.Preload(clause.Associations).Where("id = ?", id).First(&estateAgent).Error; err == nil {
		return &estateAgent, nil
	}
	lib.LogError("mutation/GetEstateAgent", err.Error())
	return nil, err
}
