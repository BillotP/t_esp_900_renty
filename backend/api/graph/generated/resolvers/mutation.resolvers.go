package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/exec"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
)

func (r *mutationResolver) SignupAsCompany(ctx context.Context, input models.CompanyInput) (*models.Company, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SignupAsEstateAgent(ctx context.Context, input *models.EstateAgentInput) (*models.EstateAgent, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SignupAsTenant(ctx context.Context, input *models.TenantInput) (*models.Tenant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTenantProfile(ctx context.Context, input *models.TenantUpdateInput) (*models.Tenant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateProperty(ctx context.Context, input *models.PropertyInput) (*models.Property, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAnomaly(ctx context.Context, input *models.AnomalyInput) (*models.Anomaly, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAnomaly(ctx context.Context, input *models.AnomalyUpdateInput) (*models.Anomaly, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns exec.MutationResolver implementation.
func (r *Resolver) Mutation() exec.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
