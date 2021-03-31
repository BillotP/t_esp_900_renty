package resolvers

import (
	"context"
	"fmt"

	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"golang.org/x/crypto/bcrypt"
)

func (r *MutationResolver) SignupAsCompany(ctx context.Context, input models.CompanyInput) (*models.Credential, error) {
	var (
		token    string
		company  *models.Company
		pwdHash  []byte
		verified bool

		err error
	)
	switch {
	case input.User == nil:
		return nil, fmt.Errorf("a company user must be provided")
	case input.User.Password == "":
		return nil, fmt.Errorf("user password can't be empty")
	}
	company = &models.Company{
		Name: input.Name,
		User: &models.User{
			Username: input.User.Username,
			Password: "",
			Role:     models.RoleCompany,
		},
		Description: &input.Description,
		Tel:         input.Tel,
		Verified:    &verified,
	}
	if err = r.DB.Where("name = ?", company.Name).First(&company).Error; err == nil {
		return nil, fmt.Errorf("company seems already register")
	}
	if pwdHash, err = bcrypt.GenerateFromPassword([]byte(input.User.Password), getPseudoRandomCost()); err != nil {
		lib.LogError("mutation/Register", err.Error())
		return nil, err
	}

	company.User.Password = string(pwdHash)
	if err = r.DB.Create(&company).Error; err != nil {
		lib.LogError("mutation/Register/Company", err.Error())
		return nil, err
	}
	if token, err = createToken(company.User.Username, company.User.Role); err != nil {
		return nil, err
	}
	return &models.Credential{
		User:  company.User,
		Token: &token,
	}, nil
}
