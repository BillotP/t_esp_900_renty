package resolvers

import (
	"context"
	"fmt"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"golang.org/x/crypto/bcrypt"
)

func (r *MutationResolver) LoginAsEstateAgent(ctx context.Context, input *models.UserInput) (*models.Credential, error) {
	var (
		estateAgent *models.EstateAgent
		token   = ""

		err error
	)

	estateAgent = &models.EstateAgent{}
	if err = r.DB.Joins("User").Where("username = ?", input.Username).First(&estateAgent).Error; err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(estateAgent.User.Password), []byte(input.Password)); err != nil {
		lib.LogError("resolvers/LoginAsCompany", err.Error())
		return nil, fmt.Errorf("bad password provided")
	}
	if token, err = createToken(estateAgent.User.Username, estateAgent.User.Role); err != nil {
		return nil, err
	}
	return &models.Credential{
		User:  estateAgent.User,
		Token: &token,
	}, nil
}

