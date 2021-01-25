package resolvers

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm/clause"
)

func (r *QueryResolver) GetEstateAgents(ctx context.Context) ([]*models.EstateAgent, error) {
	var (
		estateAgents []models.EstateAgent

		err error
	)

	if err = r.DB.Preload(clause.Associations).Find(&estateAgents).Error; err == nil {
		var estateagentsfmt []*models.EstateAgent

		for i := range estateAgents {
			estateagentsfmt = append(estateagentsfmt, &estateAgents[i])
		}
		return estateagentsfmt, nil
	}
	lib.LogError("mutation/GetEstateAgents", err.Error())
	return nil, err
}
