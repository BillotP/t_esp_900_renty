package resolvers

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm"
)

// UpdateTenantProfile udpate a tenant item in db based on input datas
func (r *MutationResolver) UpdateTenantProfile(ctx context.Context, input *models.TenantUpdateInput) (*models.Tenant, error) {
	var (
		tenant models.Tenant
		nprops []*models.Property
		ndocs  []*models.Asset
		err    error
	)

	username := ctx.Value(lib.ContextKey("username")).(string)
	tenant = models.Tenant{User: &models.User{Username: username}}
	if err = r.DB.Joins("User").Where("username = ?", tenant.User.Username).First(&tenant).Error; err != nil {
		lib.LogError("mutation/UpdateTenantProfile", err.Error())
		return nil, err
	}
	for el := range input.Properties {
		if input.Properties[el] != nil {
			property := &models.Property{
				ID: input.Properties[el],
			}
			if err = r.DB.Where("id = ?", input.Properties[el]).First(&property).Error; err != nil {
				lib.LogError("mutation/UpdateTenantProfile", err.Error())
				return nil, err
			}
			nprops = append(nprops, property)
		}
	}
	for el := range input.Documents {
		if input.Documents[el] != nil {
			document := &models.Asset{
				URL: *input.Documents[el],
			}
			if err = r.DB.Where("url = ?", input.Documents[el]).First(&document).Error; err != nil {
				lib.LogError("mutation/UpdateTenantProfile", err.Error())
				return nil, err
			}
			ndocs = append(ndocs, document)
		}
	}
	tenant.Documents = ndocs
	tenant.Properties = nprops
	if err = r.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&tenant).Error; err != nil {
		lib.LogError("mutation/UpdateTenantProfile", err.Error())
		return nil, err
	}
	return &tenant, nil
}
