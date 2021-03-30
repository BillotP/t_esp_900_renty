package resolvers

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm/clause"
)

func (r *MutationResolver) UpdateAnomaly(ctx context.Context, id int64, input *models.AnomalyUpdateInput) (*models.Anomaly, error) {
	var (
		anomaly models.Anomaly
		err     error
	)

	if err = r.DB.Where("id = ?", id).First(&anomaly).Error; err != nil {
		lib.LogError("mutation/UpdateAnomaly", err.Error())
		return nil, err
	}

	anomaly.AssignedToID = input.AssignedTo
	anomaly.State = input.State

	if err = r.DB.Preload(clause.Associations).Updates(anomaly).First(&anomaly).Error; err != nil {
		lib.LogError("mutation/UpdateAnomaly", err.Error())
		return nil, err
	}

	return &anomaly, nil
}
