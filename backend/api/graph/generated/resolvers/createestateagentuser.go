package resolvers

import (
	"context"
	"fmt"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"golang.org/x/crypto/bcrypt"
)

func (r *MutationResolver) CreateEstateAgentUser(ctx context.Context, input *models.EstateAgentInput) (*models.EstateAgent, error) {
	var (
		usernameCtx = lib.ContextKey("username")

		company     *models.Company
		estateAgent *models.EstateAgent
		pwdHash     []byte

		err error
	)

	companyUsername := ctx.Value(usernameCtx).(string)

	if err = r.DB.Joins("User").Where("username = ?", companyUsername).First(&company).Error; err != nil {
		return nil, err
	}

	estateAgent = &models.EstateAgent{
		CompanyID: company.ID,
		User: &models.User{
			Username: input.User.Username,
			Password: "",
			Role:     models.RoleEstateAgent,
		},
	}
	if err = r.DB.Joins("User").Where("username = ?", input.User.Username).First(&estateAgent).Error; err == nil {
		return nil, fmt.Errorf("estate agent seems already register")
	}
	if pwdHash, err = bcrypt.GenerateFromPassword([]byte(input.User.Password), getPseudoRandomCost()); err != nil {
		lib.LogError("mutation/Register", err.Error())
		return nil, err
	}

	estateAgent.User.Password = string(pwdHash)
	if err = r.DB.Create(&estateAgent).Error; err != nil {
		lib.LogError("mutation/Register/GetEstateAgent", err.Error())
		return nil, err
	}
	return estateAgent, nil
}
