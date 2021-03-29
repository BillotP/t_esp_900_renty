package directive

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
)

// HasRole check if the role in jwt key set in context match the schema role
func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, roles []*models.Role) (interface{}, error) {
	var userroleCtx = lib.ContextKey("userrole")
	v := ctx.Value(userroleCtx)
	userRole := fmt.Sprintf("%v", v)
	for _, role := range roles {
		if userRole == string(*role) {
			// or let it pass through
			return next(ctx)
		}
		fmt.Printf("Want ROLE : %s , have : %s\n", role, userRole)
	}
	return nil, fmt.Errorf("access denied")
}
