package resolvers

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
)

// AssignProperty update a tenant item in db based on input data
func (r *MutationResolver) AssignProperty(ctx context.Context, tenantId int64, propertyId int64) (*models.Tenant, error) {
	var (
		tenant models.Tenant
		property models.Property

		err    error
	)

	if err = r.DB.Where("id = ?", tenantId).First(&tenant).Error; err != nil {
		lib.LogError("mutation/AssignProperty", err.Error())
		return nil, err
	}

	if err = r.DB.Where("id = ?", propertyId).First(&property).Error; err != nil {
		lib.LogError("mutation/AssignProperty", err.Error())
		return nil, err
	}

	if err = r.DB.Model(&tenant).Association("Properties").Append(&property); err != nil {
		lib.LogError("mutation/AssignProperty", err.Error())
		return nil, err
	}

	return &tenant, nil
}
