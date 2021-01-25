package resolvers

import (
	"context"
	"fmt"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"golang.org/x/crypto/bcrypt"
)

func (r *MutationResolver) SignupAsAdmin(ctx context.Context, input models.AdminInput) (*models.Credential, error) {
	var (
		token   string
		admin *models.Admin
		pwdHash []byte

		err error
	)

	//verified := false

	admin = &models.Admin{
		User: &models.User{
			Username: input.User.Username,
			Password: "",
			Role:     models.RoleAdmin,
		},
	}
	if err = r.DB.Where("user_id = ?", admin.User.ID).First(&admin).Error; err == nil {
		return nil, fmt.Errorf("admin seems already register")
	}
	if pwdHash, err = bcrypt.GenerateFromPassword([]byte(input.User.Password), getPseudoRandomCost()); err != nil {
		lib.LogError("mutation/Register", err.Error())
		return nil, err
	}

	admin.User.Password = string(pwdHash)
	if err = r.DB.Create(&admin).Error; err != nil {
		lib.LogError("mutation/Register/Admin", err.Error())
		return nil, err
	}
	if token, err = createToken(admin.User.Username, admin.User.Role); err != nil {
		return nil, err
	}
	return &models.Credential{
		User:  admin.User,
		Token: &token,
	}, nil
}

