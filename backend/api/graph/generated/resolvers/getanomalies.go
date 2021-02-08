package resolvers

import (
	"context"
	"fmt"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
)

func (r *QueryResolver) GetAnomalies(ctx context.Context) ([]*models.Anomaly, error) {
	var (
		anomalies []*models.Anomaly
		err error
	)

	if err = r.DB.Find(&anomalies).Error; err != nil {
		return nil, fmt.Errorf("any anomaly found")
	}
	return anomalies, nil
}
