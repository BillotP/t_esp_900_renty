package resolvers

import (
	"context"
	"fmt"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
)

func (r *QueryResolver) GetAnomaly(ctx context.Context, id string) (*models.Anomaly, error) {
	var (
		anomaly *models.Anomaly
		err error
	)

	if err = r.DB.Where("ID = ?", id).First(&anomaly).Error; err != nil {
		return nil, fmt.Errorf("anomaly not found")
	}
	return anomaly, nil
}
