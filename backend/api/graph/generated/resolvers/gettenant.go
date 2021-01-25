package resolvers

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm/clause"
)

func (r *QueryResolver) GetTenant(ctx context.Context, id int64) (*models.Tenant, error) {
	var (
		tenant models.Tenant

		err error
	)

	if err = r.DB.Preload(clause.Associations).Where("id = ?", id).First(&tenant).Error; err == nil {
		return &tenant, nil
	}
	lib.LogError("mutation/GetTenant", err.Error())
	return nil, err
}
