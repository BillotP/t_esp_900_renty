package resolvers_test

import (
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/exec"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/resolvers"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/directive"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/middleware"
)

var (
	Server *client.Client
)

func init() {
	middleware.InitDB()
	c := exec.Config{
		Resolvers: &resolvers.Resolver{DB: middleware.Db},
		Directives: exec.DirectiveRoot{
			HasRole: directive.HasRole,
		}}
	Server = client.New(handler.NewDefaultServer(exec.NewExecutableSchema(c)))
}
