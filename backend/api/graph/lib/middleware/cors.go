package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

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

// CORSMiddleware allow all frontends to reach this .server.
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
