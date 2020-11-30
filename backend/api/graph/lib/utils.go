package lib

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/xid"
)

// ContextKey allow storing and sharing information on context.Context as strings
type ContextKey string

// GetJwtString return a signed token with exp set in hours
func GetJwtString(exp time.Duration, sub string, otp bool) *string {
	var err error
	var tokenString string
	expirationTime := time.Now().Add(exp)
	audience := ServerConf.Host + ":"
	jID := xid.NewWithTime(time.Now()).String()
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    ServerConf.JwtIssuer,
		Audience:  audience,
		Subject:   sub,
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        jID,
	})
	// Sign and get the complete encoded token as a string using the secret
	key := ServerConf.JwtSigningKey
	if tokenString, err = token.SignedString([]byte(key)); err != nil {
		LogError("lib/GetJwtString", err.Error())
	}
	return &tokenString
}
