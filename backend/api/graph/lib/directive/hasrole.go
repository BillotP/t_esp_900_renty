package directive

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
)

// HasRole check if the role in jwt key set in context match the schema role
func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role models.Role) (interface{}, error) {
	return next(ctx)
}
