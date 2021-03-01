package resolvers

import (
	"context"
	"fmt"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"golang.org/x/crypto/bcrypt"
)

func (r *MutationResolver) LoginAsCompany(ctx context.Context, input *models.UserInput) (*models.Credential, error) {
	var (
		company *models.Company
		token   = ""

		err error
	)

	company = &models.Company{}
	if err = r.DB.Joins("User").Where("username = ?", input.Username).First(&company).Error; err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(company.User.Password), []byte(input.Password)); err != nil {
		lib.LogError("resolvers/LoginAsCompany", err.Error())
		return nil, fmt.Errorf("bad password provided")
	}
	if token, err = createToken(company.User.Username, company.User.Role); err != nil {
		return nil, err
	}
	return &models.Credential{
		User:  company.User,
		Token: &token,
	}, nil
}
