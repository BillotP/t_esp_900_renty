package resolvers

import (
	"context"
	"fmt"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
)

func (r *MutationResolver) CreateAnomaly(ctx context.Context, input *models.AnomalyInput) (*models.Anomaly, error) {
	var (
		anomaly *models.Anomaly
		err error
	)
	anomaly = &models.Anomaly{
		PropertyID:     input.Property,
		Type: 			input.Type,
		Description:	input.Description,
	}

	if err = r.DB.Where(&anomaly).First(&anomaly).Error; err == nil {
		return nil, fmt.Errorf("anomaly already created")
	}
	if err = r.DB.Create(&anomaly).Error; err != nil {
		lib.LogError("mutation/Register/GetAnomaly", err.Error())
		return nil, err
	}
	return anomaly, nil
}
