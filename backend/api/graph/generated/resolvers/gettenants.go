package resolvers

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm/clause"
)

func (r *QueryResolver) Tenants(ctx context.Context) ([]*models.Tenant, error) {
	var (
		estateAgent models.EstateAgent

		tenants []*models.Tenant

		err error
	)

	username := ctx.Value(lib.ContextKey("username")).(string)

	if err = r.DB.Joins("User").Where("username = ?", username).First(&estateAgent).Error; err != nil {
		lib.LogError("mutation/GetTenants", err.Error())
		return nil, err
	}

	if err = r.DB.Preload(clause.Associations).Where("estate_agent_id = ?", estateAgent.ID).Find(&tenants).Error; err != nil {
		lib.LogError("mutation/GetTenants", err.Error())
		return nil, err
	}

	return tenants, err
}
