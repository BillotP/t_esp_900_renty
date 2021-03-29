package resolvers

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm/clause"
)

func (r *QueryResolver) Tenants(ctx context.Context) ([]*models.Tenant, error) {
	var (
		tenants []models.Tenant

		err error
	)

	if err = r.DB.Preload(clause.Associations).Find(&tenants).Error; err == nil {
		var tenantsfmt []*models.Tenant

		for i := range tenants {
			tenantsfmt = append(tenantsfmt, &tenants[i])
		}
		return tenantsfmt, nil
	}
	lib.LogError("mutation/GetTenants", err.Error())
	return nil, err
}
