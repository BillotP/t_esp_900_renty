package resolvers

import (
	"context"
	"fmt"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
)

func (r *QueryResolver) Profile(ctx context.Context) (models.Profile, error) {
	var (
		usernameCtx = lib.ContextKey("username")

		user models.User

		company     models.Company
		estateAgent models.EstateAgent
		tenant      models.Tenant

		err error
	)

	profileUsername := ctx.Value(usernameCtx).(string)

	user = models.User{
		Username: profileUsername,
	}

	if err = r.DB.Where("username = ?", profileUsername).First(&user).Error; err != nil {
		lib.LogError("query/Profile", err.Error())
		return nil, err
	}

	if err = r.DB.Where("user_id = ?", user.ID).First(&company).Error; err == nil {
		return company, nil
	}
	if err = r.DB.Where("user_id = ?", user.ID).First(&estateAgent).Error; err == nil {
		return estateAgent, nil
	}

	if err = r.DB.Where("user_id = ?", user.ID).First(&tenant).Error; err == nil {
		return tenant, nil
	}

	lib.LogError("query/Profile", "token is undefined")
	return nil, fmt.Errorf("query/Profile: token is undefined")
}
