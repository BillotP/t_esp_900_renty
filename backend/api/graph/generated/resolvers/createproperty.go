package resolvers

import (
	"context"
	"fmt"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
)

func (r *MutationResolver) CreateProperty(ctx context.Context, input *models.PropertyInput) (*models.Property, error) {
	var (
		estateAgent models.EstateAgent

		property *models.Property
		err      error
	)

	username := ctx.Value(lib.ContextKey("username")).(string)
	estateAgent = models.EstateAgent{User: &models.User{Username: username}}

	if err = r.DB.Joins("User").Joins("Company").Where("username = ?", estateAgent.User.Username).First(&estateAgent).Error; err != nil {
		lib.LogError("mutation/CreateProperty", err.Error())
		return nil, err
	}

	property = &models.Property{
		Area:       input.Area,
		Address:    input.Address,
		CodeNumber: input.CodeNumber,
		Type:       input.Type,
		Company:    estateAgent.Company,
		CompanyID:  estateAgent.CompanyID,
	}
	if err = r.DB.Where("address = ?", property.Address).First(&property).Error; err == nil {
		return nil, fmt.Errorf("there is already a property at this address")
	}
	if err = r.DB.Create(&property).Error; err != nil {
		lib.LogError("mutation/Register/Property", err.Error())
		return nil, err
	}
	return property, nil
}
