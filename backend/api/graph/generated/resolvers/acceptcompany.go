package resolvers

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
)

func (r *MutationResolver) AcceptCompany(ctx context.Context, id int64) (*models.Company, error) {
	var (
		company  *models.Company
		verified = true

		err error
	)

	company = &models.Company{
		ID:       &id,
		Verified: &verified,
	}
	if err = r.DB.Updates(&company).Error; err == nil {
		return company, nil
	}
	lib.LogError("mutation/AcceptCompany", err.Error())
	return nil, err
}
