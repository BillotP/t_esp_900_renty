package resolvers

import (
	"context"
	"fmt"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"golang.org/x/crypto/bcrypt"
)

func (r *MutationResolver) LoginAsTenant(ctx context.Context, input *models.UserInput) (*models.Credential, error) {
	var (
		tenant *models.Tenant
		token  = ""

		err error
	)

	tenant = &models.Tenant{}
	if err = r.DB.Joins("User").Where("username = ?", input.Username).First(&tenant).Error; err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(tenant.User.Password), []byte(input.Password)); err != nil {
		lib.LogError("resolvers/LoginAsCompany", err.Error())
		return nil, fmt.Errorf("bad password provided")
	}
	if token, err = createToken(tenant.User.Username, tenant.User.Role); err != nil {
		return nil, err
	}
	return &models.Credential{
		User:  tenant.User,
		Token: &token,
	}, nil
}
