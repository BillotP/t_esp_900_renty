package resolvers

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm/clause"
)

func (r *QueryResolver) GetAnomalies(ctx context.Context) ([]*models.Anomaly, error) {
	var (
		anomalies []models.Anomaly
		err error
	)

	if err = r.DB.Preload(clause.Associations).Find(&anomalies).Error; err == nil {
		var anomaliesfmt []*models.Anomaly

		for i := range anomalies {
			anomaliesfmt = append(anomaliesfmt, &anomalies[i])
		}
		return anomaliesfmt, nil
	}
	lib.LogError("mutation/GetAnomalies", err.Error())
	return nil, err
}
