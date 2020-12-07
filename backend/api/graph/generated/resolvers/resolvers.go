package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"fmt"

	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/exec"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
)

type Resolver struct {
	DB *gorm.DB
}

func (r *mutationResolver) SignupAsAdmin(ctx context.Context, input models.AdminInput) (*models.Admin, error) {
	panic("Miguel : not implemented")
}

func (r *mutationResolver) SignupAsCompany(ctx context.Context, input models.CompanyInput) (*models.Company, error) {
	var (
		user    *models.User
		company *models.Company

		err error
	)

	verified := false
	user = &models.User{
		Username: input.User.Username,
		Password: input.User.Password,
		Role:     models.RoleCompany,
	}
	if err = r.DB.Where("username = ?", user.Username).First(&user).Error; err == nil {
		return nil, fmt.Errorf("user seems already register")
	}
	if err = r.DB.Create(&user).Error; err != nil {
		lib.LogError("mutation/Register/User", err.Error())
		return nil, err
	}
	company = &models.Company{
		Name:        input.Name,
		UserID:      user.ID,
		Description: &input.Description,
		Tel:         input.Tel,
		Verified:    &verified,
	}
	if err = r.DB.Where("name = ?", company.Name).First(&company).Error; err == nil {
		return nil, fmt.Errorf("company seems already register")
	}
	if err = r.DB.Create(&company).Error; err != nil {
		lib.LogError("mutation/Register/Company", err.Error())
		return nil, err
	}
	return company, nil
}

func (r *mutationResolver) CreateEstateAgentUser(ctx context.Context, input *models.EstateAgentInput) (*models.EstateAgent, error) {
	panic("Miguel : not implemented")
}

func (r *mutationResolver) CreateTenantUser(ctx context.Context, input *models.TenantInput) (*models.Tenant, error) {
	panic("Miguel : not implemented")
}

func (r *mutationResolver) AcceptCompany(ctx context.Context, id int64) (*models.Company, error) {
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

func (r *mutationResolver) LoginAsCompany(ctx context.Context, input *models.UserInput) (*models.Company, error) {
	panic("Miguel : not implemented")
}

func (r *mutationResolver) LoginAsEstateAgent(ctx context.Context, input *models.UserInput) (*models.EstateAgent, error) {
	panic("Miguel : not implemented")
}

func (r *mutationResolver) LoginAsTenant(ctx context.Context, input *models.UserInput) (*models.Tenant, error) {
	panic("Miguel : not implemented")
}

func (r *mutationResolver) UpdateTenantProfile(ctx context.Context, input *models.TenantUpdateInput) (*models.Tenant, error) {
	panic("Wilfried : not implemented")
}

func (r *mutationResolver) CreateProperty(ctx context.Context, input *models.PropertyInput) (*models.Property, error) {
	panic("Remi : not implemented")
}

func (r *mutationResolver) CreateAnomaly(ctx context.Context, input *models.AnomalyInput) (*models.Anomaly, error) {
	panic("Remi : not implemented")
}

func (r *mutationResolver) UpdateAnomaly(ctx context.Context, input *models.AnomalyUpdateInput) (*models.Anomaly, error) {
	panic("Remi : not implemented")
}

func (r *queryResolver) Anomaly(ctx context.Context, id string) (*models.Anomaly, error) {
	panic("Remi : not implemented")
}

func (r *queryResolver) Anomalies(ctx context.Context) ([]*models.Anomaly, error) {
	panic("Remi : not implemented")
}

func (r *queryResolver) Tenant(ctx context.Context, id int64) (*models.Tenant, error) {
	var (
		tenant models.Tenant

		err error
	)

	if err = r.DB.Preload(clause.Associations).Where("id = ?", id).First(&tenant).Error; err == nil {
		return &tenant, nil
	}
	lib.LogError("mutation/Tenant", err.Error())
	return nil, err
}

func (r *queryResolver) Tenants(ctx context.Context) ([]*models.Tenant, error) {
	var (
		tenants []models.Tenant

		err error
	)

	if err = r.DB.Preload(clause.Associations).Find(&tenants).Error; err == nil {
		var tenantsfmt []*models.Tenant

		for i := range tenants {
			tenantsfmt = append(tenantsfmt, &tenants[i])
		}
		return tenantsfmt, nil
	}
	lib.LogError("mutation/Tenants", err.Error())
	return nil, err
}

func (r *queryResolver) EstateAgent(ctx context.Context, id int64) (*models.EstateAgent, error) {
	var (
		estateAgent models.EstateAgent

		err error
	)

	if err = r.DB.Preload(clause.Associations).Where("id = ?", id).First(&estateAgent).Error; err == nil {
		return &estateAgent, nil
	}
	lib.LogError("mutation/EstateAgent", err.Error())
	return nil, err
}

func (r *queryResolver) EstateAgents(ctx context.Context) ([]*models.EstateAgent, error) {
	var (
		estateAgents []models.EstateAgent

		err error
	)

	if err = r.DB.Preload(clause.Associations).Find(&estateAgents).Error; err == nil {
		var estateagentsfmt []*models.EstateAgent

		for i := range estateAgents {
			estateagentsfmt = append(estateagentsfmt, &estateAgents[i])
		}
		return estateagentsfmt, nil
	}
	lib.LogError("mutation/EstateAgents", err.Error())
	return nil, err
}

func (r *queryResolver) Company(ctx context.Context, id int64) (*models.Company, error) {
	var (
		company models.Company

		err error
	)

	if err = r.DB.Preload(clause.Associations).Where("id = ?", id).First(&company).Error; err == nil {
		return &company, nil
	}
	lib.LogError("mutation/Company", err.Error())
	return nil, err
}

func (r *queryResolver) Companies(ctx context.Context) ([]*models.Company, error) {
	var (
		companies []models.Company

		err error
	)

	if err = r.DB.Preload(clause.Associations).Find(&companies).Error; err == nil {
		var companiesfmt []*models.Company

		for i := range companies {
			companiesfmt = append(companiesfmt, &companies[i])
		}
		return companiesfmt, nil
	}
	lib.LogError("mutation/Companies", err.Error())
	return nil, err
}

// Mutation returns exec.MutationResolver implementation.
func (r *Resolver) Mutation() exec.MutationResolver { return &mutationResolver{r} }

// Query returns exec.QueryResolver implementation.
func (r *Resolver) Query() exec.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
