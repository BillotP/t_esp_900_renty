package main

import (
	"api/graph"
	"api/graph/generated"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

var basePath = os.Getenv("BASE_PATH")

const (
	headerSep = ", "
)

var allowMethods = []string{
	"GET",
	"PUT",
	"POST",
	"OPTIONS",
} // I think we don't need more as we're in GQL API, not REST (dave)
var allowHeaders = []string{
	"Content-Type",
	"Content-Length",
	"Accept-Encoding",
	"X-CSRF-Token",
	"Authorization",
	"Accept",
	"Origin",
	"Cache-Control",
	"X-Requested-With",
} // TODO(cors): Review this list to get only what's needed.
var allowOrigins = []string{
	"*",
} // TODO(cors): To be updated when we'll got all our definitives hostnames !!

var allowHeaderString = strings.Join(allowHeaders, headerSep)
var allowMethodsString = strings.Join(allowMethods, headerSep)
var allowOriginString = strings.Join(allowOrigins, headerSep)

// CORSMiddleware allow all frontends to reach this server.
//
// It set CORS protocol headers variables in http response
// and return HTTP 200 for OPTIONS preflight request.
//
// Shamessly inspired by https://stackoverflow.com/questions/29418478/go-gin-framework-cors.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOriginString)
		c.Writer.Header().Set("Access-Control-Allow-Headers", allowHeaderString)
		c.Writer.Header().Set("Access-Control-Allow-Methods", allowMethodsString)

		// ALlowing preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{},
	}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", basePath+"/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Setting up Gin
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.POST(basePath+"/query", graphqlHandler())
	r.GET(basePath+"/", playgroundHandler())
	fmt.Printf("Info(server): Server listening on %s:8080\n", basePath)
	r.Run()
}
