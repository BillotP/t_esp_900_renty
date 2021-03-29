package resolvers

import (
	"context"
	"fmt"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
)

func (r *MutationResolver) CreateAnomaly(ctx context.Context, input *models.AnomalyInput) (*models.Anomaly, error) {
	var (
		tenant models.Tenant

		anomaly models.Anomaly
		err     error
	)

	username := ctx.Value(lib.ContextKey("username")).(string)
	tenant = models.Tenant{User: &models.User{Username: username}}
	if err = r.DB.Preload("EstateAgent").Joins("User").Where("username = ?", tenant.User.Username).First(&tenant).Error; err != nil {
		lib.LogError("mutation/CreateAnomaly", err.Error())
		return nil, err
	}

	if err = r.DB.Model(&tenant).Association("Properties").Find(&tenant.Properties); err != nil {
		lib.LogError("mutation/CreateAnomaly", err.Error())
		return nil, err
	}


	anomaly = models.Anomaly{
		PropertyID:   tenant.Properties[0].ID,
		Type:         input.Type,
		Description:  input.Description,
		CreateByID:   tenant.ID,
		AssignedToID: tenant.EstateAgent.ID,
	}

	fmt.Println(*tenant.Properties[0].ID)
	fmt.Println(tenant.Properties[0])
	fmt.Println(input.Type)
	fmt.Println(input.Description)
	fmt.Println(tenant)
	fmt.Println(*tenant.ID)
	fmt.Println(tenant.EstateAgent)
	fmt.Println(*tenant.EstateAgent.ID)

	if err = r.DB.Where(anomaly).First(&anomaly).Error; err == nil {
		return nil, fmt.Errorf("anomaly already created")
	}

	if err = r.DB.Create(&anomaly).Error; err != nil {
		lib.LogError("mutation/CreateAnomaly", err.Error())
		return nil, err
	}
	return &anomaly, nil
}
