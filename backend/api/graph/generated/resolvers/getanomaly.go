package resolvers

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm/clause"
)

func (r *QueryResolver) Anomaly(ctx context.Context, id int64) (*models.Anomaly, error) {
	var (
		anomaly models.Anomaly
		err error
	)

	if err = r.DB.Preload(clause.Associations).Preload("CreateBy.User").Preload("AssignedTo.User").Where("id = ?", id).First(&anomaly).Error; err != nil {
		lib.LogError("mutation/GetAnomaly", err.Error())
		return nil, err
	}
	return &anomaly, nil
}
