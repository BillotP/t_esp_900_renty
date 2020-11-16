package middleware

import (
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/exec"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/resolvers"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/directive"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// GraphqlHandler is the main graphql API handler
func GraphqlHandler() gin.HandlerFunc {
	c := exec.Config{
		Resolvers: &resolvers.Resolver{
			DB: Db,
		},
		Directives: exec.DirectiveRoot{
			HasRole:    directive.HasRole,
		}}
	srv := handler.NewDefaultServer(exec.NewExecutableSchema(c))
	srv.AddTransport(transport.POST{})
	srv.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// TODO: 🚨 Please define authorized hosts !
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}
