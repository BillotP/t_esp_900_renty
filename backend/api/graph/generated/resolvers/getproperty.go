package resolvers

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm/clause"
)

func (r *QueryResolver) Property(ctx context.Context, id int64) (*models.Property, error) {
	var (
		property models.Property
		err error
	)

	if err = r.DB.Preload(clause.Associations).Where("id = ?", id).First(&property).Error; err == nil {
		return &property, nil
	}
	lib.LogError("mutation/GetProperty", err.Error())
	return nil, err
}
