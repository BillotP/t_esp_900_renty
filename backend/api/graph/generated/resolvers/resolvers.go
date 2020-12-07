package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"fmt"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm"

	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/exec"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
)

type Resolver struct {
	DB *gorm.DB
}

func (r *mutationResolver) SignupAsAdmin(ctx context.Context, input models.AdminInput) (*models.Admin, error) {
	panic("not implemented")
}

func (r *mutationResolver) SignupAsCompany(ctx context.Context, input models.CompanyInput) (*models.Company, error) {
	var (
		company *models.Company

		err error
	)

	verified := false
	company = &models.Company{
		Name:        input.Name,
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
	panic("not implemented")
}

func (r *mutationResolver) CreateTenantUser(ctx context.Context, input *models.TenantInput) (*models.Tenant, error) {
	panic("not implemented")
}

func (r *mutationResolver) AcceptCompany(ctx context.Context) (*models.Company, error) {
	panic("not implemented")
}

func (r *mutationResolver) LoginAsCompany(ctx context.Context, input *models.UserInput) (*models.Company, error) {
	panic("not implemented")
}

func (r *mutationResolver) LoginAsEstateAgent(ctx context.Context, input *models.UserInput) (*models.EstateAgent, error) {
	panic("not implemented")
}

func (r *mutationResolver) LoginAsTenant(ctx context.Context, input *models.UserInput) (*models.Tenant, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateTenantProfile(ctx context.Context, input *models.TenantUpdateInput) (*models.Tenant, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateProperty(ctx context.Context, input *models.PropertyInput) (*models.Property, error) {
	var (
		property *models.Property
		err error
	)
	property = &models.Property{
		Area:       input.Area,
		Address: 	input.Address,
		CodeNumber: input.CodeNumber,
		Type:		input.Type,
	}
	if err = r.DB.Where("codeNumber = ?", property.CodeNumber).First(&property).Error; err == nil {
		return nil, fmt.Errorf("Property already registered")
	}
	if err = r.DB.Create(&property).Error; err != nil {
		lib.LogError("mutation/Register/Property", err.Error())
		return nil, err
	}
	return property, nil
}

func (r *mutationResolver) CreateAnomaly(ctx context.Context, input *models.AnomalyInput) (*models.Anomaly, error) {
	var (
		anomaly *models.Anomaly
		err error
	)
	anomaly = &models.Anomaly{
		PropertyID:     input.Property,
		Type: 			input.Type,
		Description:	input.Description,
	}

	if err = r.DB.Where("codeNumber = ?", anomaly.PropertyID).First(&anomaly).Error; err == nil {
		return nil, fmt.Errorf("Anomaly already created")
	}
	if err = r.DB.Create(&anomaly).Error; err != nil {
		lib.LogError("mutation/Register/Anomaly", err.Error())
		return nil, err
	}
	return anomaly, nil
}

func (r *mutationResolver) UpdateAnomaly(ctx context.Context, input *models.AnomalyUpdateInput) (*models.Anomaly, error) {
	var (
		anomaly *models.Anomaly
		err error
	)
	anomaly = &models.Anomaly{
		AssignedToID:   input.AssignedTo,
		State: 			input.State,
	}
	if err = r.DB.Update('', &anomaly).Error; err != nil {
		lib.LogError("mutation/Register/Anomaly", err.Error())
		return nil, err
	}
	return anomaly, nil
}

func (r *queryResolver) Anomaly(ctx context.Context, id string) (*models.Anomaly, error) {
	var (
		anomaly *models.Anomaly
		err error
	)
	anomaly = &models.Anomaly{
		PropertyID:     input.Property,
		Type: 			input.Type,
		Description:	input.Description,
	}

	if err = r.DB.Where("codeNumber = ?", anomaly.PropertyID).First(&anomaly).Error; err == nil {
		return nil, fmt.Errorf("Anomaly not found")
	}
	return anomaly, nil
}

func (r *queryResolver) Anomalies(ctx context.Context) ([]*models.Anomaly, error) {
	panic("not implemented")
}

func (r *queryResolver) Tenant(ctx context.Context, id string) (*models.Tenant, error) {
	panic("not implemented")
}

func (r *queryResolver) Tenants(ctx context.Context) ([]*models.Tenant, error) {
	panic("not implemented")
}

func (r *queryResolver) EstateAgent(ctx context.Context, id string) (*models.EstateAgent, error) {
	panic("not implemented")
}

func (r *queryResolver) EstateAgents(ctx context.Context) ([]*models.EstateAgent, error) {
	panic("not implemented")
}

func (r *queryResolver) Company(ctx context.Context, id string) (*models.Company, error) {
	panic("not implemented")
}

func (r *queryResolver) Companies(ctx context.Context) ([]*models.Company, error) {
	panic("not implemented")
}

// Mutation returns exec.MutationResolver implementation.
func (r *Resolver) Mutation() exec.MutationResolver { return &mutationResolver{r} }

// Query returns exec.QueryResolver implementation.
func (r *Resolver) Query() exec.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
