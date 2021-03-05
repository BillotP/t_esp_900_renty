package resolvers

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm/clause"
)

func (r *QueryResolver) GetAnomaly(ctx context.Context, id string) (*models.Anomaly, error) {
	var (
		anomaly *models.Anomaly
		err error
	)

	/*if err = r.DB.Where("ID = ?", id).First(&anomaly).Error; err != nil {
		return nil, fmt.Errorf("anomaly not found")
	}
	return anomaly, nil*/

	if err = r.DB.Preload(clause.Associations).Where("id = ?", id).First(&anomaly).Error; err == nil {
		return anomaly, nil
	}
	lib.LogError("mutation/GetAnomaly", err.Error())
	return nil, err
}
