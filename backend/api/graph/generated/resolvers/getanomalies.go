package resolvers

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm/clause"
)

func (r *QueryResolver) Anomalies(ctx context.Context) ([]*models.Anomaly, error) {
	var (
		estateAgent models.EstateAgent
		tenant      models.Tenant

		anomalies []*models.Anomaly
		err       error
	)

	username := ctx.Value(lib.ContextKey("username")).(string)

	if err = r.DB.Joins("User").Where("username = ?", username).First(&tenant).Error; err == nil {
		if err = r.DB.Preload(clause.Associations).Where("create_by_id = ?", tenant.ID).Find(&anomalies).Error; err != nil {
			lib.LogError("mutation/GetAnomalies", err.Error())
			return nil, err
		}
		return anomalies, nil
	}

	lib.LogError("mutation/GetAnomalies", err.Error())

	if err = r.DB.Joins("User").Where("username = ?", username).First(&estateAgent).Error; err == nil {
		if err = r.DB.Preload(clause.Associations).Where("assigned_to_id = ?", estateAgent.ID).Find(&anomalies).Error; err != nil {
			lib.LogError("mutation/GetAnomalies", err.Error())
			return nil, err
		}
		return anomalies, nil
	}

	lib.LogError("mutation/GetAnomalies", err.Error())
	return nil, err
}
