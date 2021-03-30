package resolvers

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm/clause"
)

func (r *QueryResolver) EstateAgents(ctx context.Context) ([]*models.EstateAgent, error) {
	var (
		estateAgent  models.EstateAgent
		company      models.Company
		estateAgents []*models.EstateAgent

		companyId int64

		err error
	)

	username := ctx.Value(lib.ContextKey("username")).(string)

	if err = r.DB.Joins("User").Where("username = ?", username).First(&company).Error; err != nil {
		lib.LogError("mutation/GetEstateAgents", err.Error())

		if err = r.DB.Joins("User").Where("username = ?", username).First(&estateAgent).Error; err != nil {
			lib.LogError("mutation/GetEstateAgents", err.Error())
			return nil, err
		} else {
			companyId = *estateAgent.CompanyID
		}
	} else {
		companyId = *company.ID
	}

	if err = r.DB.Preload(clause.Associations).Where("company_id = ?", companyId).Find(&estateAgents).Error; err != nil {
		lib.LogError("mutation/GetEstateAgents", err.Error())
		return nil, err
	}

	return estateAgents, nil
}
