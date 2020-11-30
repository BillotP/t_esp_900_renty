package middleware

import (
	"context"
	"fmt"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"

	"regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// SetContext get the Authorization header field, get the jwt and set the recorded value in context
func SetContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Code for the middleware...
		var claims jwt.StandardClaims
		re := regexp.MustCompile(`(?i)bearer`)
		ctx := c.Request.Context()
		authHeader := c.GetHeader("Authorization")

		userIPCtx := lib.ContextKey("userip")
		userAgentCtx := lib.ContextKey("useragent")
		userIP := c.ClientIP()
		ctx = context.WithValue(ctx, userIPCtx, userIP)
		ctx = context.WithValue(ctx, userAgentCtx, c.GetHeader("User-Agent"))
		if re.Match([]byte(authHeader)) {
			authHeader = re.ReplaceAllString(authHeader, "")
			authHeader = strings.Trim(authHeader, " ")
		}
		_, err := jwt.ParseWithClaims(authHeader, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(lib.ServerConf.JwtSigningKey), nil
		})
		if err == nil && claims.Valid() == nil {
			userInfos := strings.Split(claims.Subject, ":")
			subjectInfos := strings.Split(claims.Audience, ":")
			namekontext := lib.ContextKey("username")
			rolekontext := lib.ContextKey("userrole")
			subjectkontext := lib.ContextKey("subject")
			if len(userInfos) > 0 && len(subjectInfos) > 0 {
				ctx = context.WithValue(ctx, namekontext, userInfos[0])
				ctx = context.WithValue(ctx, rolekontext, userInfos[1])
				ctx = context.WithValue(ctx, subjectkontext, subjectInfos[1])
			} else {
				logmsg := fmt.Sprintf("Weird token : %+v\n", claims)
				lib.LogError("middleware/SetContext", logmsg)
			}
		} else if err.Error() != "token contains an invalid number of segments" {
			lib.LogError("middleware/SetContext", err.Error())
		}
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
