package function

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/BillotP/gorenty"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get("https://keys.192.168.1.20.nip.io/.well-known/jwks.json")

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		// Verify 'aud' claim
		aud := "myapi"
		checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
		if !checkAud {
			return token, errors.New("invalid audience")
		}
		// Verify 'iss' claim
		iss := "https://api.192.168.1.20.nip.io/"
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return token, errors.New("invalid issuer")
		}

		cert, err := getPemCert(token)
		if err != nil {
			panic(err.Error())
		}

		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	},
	SigningMethod: jwt.SigningMethodRS256,
})

func createToken(user string) (string, error) {
	var err error
	var signString string
	var signKey *rsa.PrivateKey
	if signString, err = goscrappy.GetSecret("jwt_private_key"); err != nil {
		return "", err
	}
	if signKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(signString)); err != nil {
		return "", err
	}

	// create a signer for rsa 256
	t := jwt.New(jwt.SigningMethodRS256)

	// set our claims
	t.Claims = &jwt.StandardClaims{
		// set the expire time
		// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
		ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		Issuer:    "https://api.192.168.1.20.nip.io/",
		Audience:  "myapi",
		Subject:   user,
	}
	t.Header["kid"] = "web-key"
	// Creat token string
	return t.SignedString(signKey)
}
