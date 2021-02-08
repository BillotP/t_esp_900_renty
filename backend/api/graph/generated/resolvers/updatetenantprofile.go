package resolvers

import (
	"context"
	"strconv"

	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
)

// UpdateTenantProfile udpate a tenant item in db based on input datas
func (r *MutationResolver) UpdateTenantProfile(ctx context.Context, input *models.TenantUpdateInput) (*models.Tenant, error) {
	var (
		tenant models.Tenant
		nprops []*models.Property
		ndocs  []*models.Asset
		err    error
	)

	id := ctx.Value(lib.ContextKey("username")).(string)
	idVal, _ := strconv.ParseInt(id, 10, 64)
	tenant.ID = &idVal
	if err = r.DB.First(&tenant).Error; err != nil {
		lib.LogError("mutation/UpdateTenantProfile", err.Error())
		return nil, err
	}
	for el := range input.Properties {
		if input.Properties[el] != nil {
			nprops = append(nprops, &models.Property{
				ID: input.Properties[el],
			})
		}
	}
	for el := range input.Documents {
		if input.Documents[el] != nil {
			ndocs = append(ndocs, &models.Asset{
				ID: input.Properties[el],
			})
		}
	}
	tenant.Documents = ndocs
	tenant.Properties = nprops
	if err = r.DB.Updates(&tenant).Error; err != nil {
		lib.LogError("mutation/UpdateTenantProfile", err.Error())
		return nil, err
	}
	return &tenant, nil
}
