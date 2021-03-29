package resolvers

import (
	"context"
	"fmt"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"golang.org/x/crypto/bcrypt"
)

func (r *MutationResolver) CreateTenantUser(ctx context.Context, input *models.TenantInput) (*models.Tenant, error) {
	var (
		estateAgent models.EstateAgent
		tenant      *models.Tenant
		pwdHash     []byte

		err error
	)

	username := ctx.Value(lib.ContextKey("username")).(string)

	fmt.Println(username)
	fmt.Println(username)
	fmt.Println(username)
	fmt.Println(username)
	fmt.Println(username)
	fmt.Println(username)
	fmt.Println(username)
	fmt.Println(username)
	fmt.Println(username)
	fmt.Println(username)
	fmt.Println(username)
	fmt.Println(username)
	fmt.Println(username)

	if err = r.DB.Joins("User").Where("username = ?", username).First(&estateAgent).Error; err != nil {
		return nil, err
	}

	tenant = &models.Tenant{
		EstateAgentID: estateAgent.ID,
		User: &models.User{
			Username: input.User.Username,
			Password: "",
			Role:     models.RoleTenant,
		},
	}
	if err = r.DB.Joins("User").Where("username = ?", input.User.Username).First(&tenant).Error; err == nil {
		return nil, fmt.Errorf("tenant seems already register")
	}
	if pwdHash, err = bcrypt.GenerateFromPassword([]byte(input.User.Password), getPseudoRandomCost()); err != nil {
		lib.LogError("mutation/Register", err.Error())
		return nil, err
	}

	tenant.User.Password = string(pwdHash)
	if err = r.DB.Create(&tenant).Error; err != nil {
		lib.LogError("mutation/Register/GetTenant", err.Error())
		return nil, err
	}
	return tenant, nil
}
