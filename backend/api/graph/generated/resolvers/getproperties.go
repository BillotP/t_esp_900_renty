package resolvers

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm/clause"
)

func (r *QueryResolver) Properties(ctx context.Context) ([]*models.Property, error) {
	var (
		estateAgent models.EstateAgent
		properties  []models.Property
		err         error
	)

	username := ctx.Value(lib.ContextKey("username")).(string)
	estateAgent = models.EstateAgent{User: &models.User{Username: username}}
	if err = r.DB.Joins("User").Where("username = ?", estateAgent.User.Username).First(&estateAgent).Error; err != nil {
		lib.LogError("mutation/GetProperties", err.Error())
		return nil, err
	}

	if err = r.DB.Joins("Company").Preload(clause.Associations).Where("company_id = ?", estateAgent.CompanyID).Find(&properties).Error; err == nil {
		var propertiesfmt []*models.Property

		for i := range properties {
			propertiesfmt = append(propertiesfmt, &properties[i])
		}
		return propertiesfmt, nil
	}
	lib.LogError("mutation/GetProperties", err.Error())
	return nil, err
}
