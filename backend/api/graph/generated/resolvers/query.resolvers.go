package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/server"
)

func (r *queryResolver) Anomaly(ctx context.Context, id string) (*models.Anomaly, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Anomalies(ctx context.Context) ([]*models.Anomaly, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tenant(ctx context.Context, id string) (*models.Tenant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tenants(ctx context.Context) ([]*models.Tenant, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns server.QueryResolver implementation.
func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
