package resolvers

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
)

func (r *MutationResolver) UpdateAnomaly(ctx context.Context, input *models.AnomalyUpdateInput) (*models.Anomaly, error) {
	var (
		anomaly *models.Anomaly
		err error
	)
	anomaly = &models.Anomaly{
		AssignedToID:   input.AssignedTo,
		State: 			input.State,
	}
	if err = r.DB.Updates(&anomaly).First(&anomaly).Error; err != nil {
		lib.LogError("mutation/Register/GetAnomaly", err.Error())
		return nil, err
	}
	return anomaly, nil
}
