package middleware

import (
	"context"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// SetContext get the Authorization header field, get the jwt and set the recorded value in context
func SetContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Code for the middleware...
		var claims jwt.MapClaims
		ctx := c.Request.Context()
		authHeader := c.GetHeader("Authorization")

		_, err := jwt.ParseWithClaims(authHeader, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(lib.ServerConf.JwtSigningKey), nil
		})
		if err == nil && claims.Valid() == nil {
			namekontext := lib.ContextKey("username")
			rolekontext := lib.ContextKey("userrole")
			ctx = context.WithValue(ctx, namekontext, claims["username"])
			ctx = context.WithValue(ctx, rolekontext, claims["userrole"])
		} else if err.Error() != "token contains an invalid number of segments" {
			lib.LogError("middleware/SetContext", err.Error())
		}
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
