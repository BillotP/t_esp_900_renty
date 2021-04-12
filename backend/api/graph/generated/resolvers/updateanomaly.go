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

	if input.AssignedTo != nil {
		anomaly.AssignedToID = input.AssignedTo
	}

	if input.State != nil {
		anomaly.State = input.State
	}

	if input.Priority != nil {
		anomaly.Priority = input.Priority
	}

	if err = r.DB.Preload(clause.Associations).Updates(anomaly).First(&anomaly).Error; err != nil {
		lib.LogError("mutation/UpdateAnomaly", err.Error())
		return nil, err
	}

	return &anomaly, nil
}
