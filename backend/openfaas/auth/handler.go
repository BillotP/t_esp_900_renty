package function

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/BillotP/renty/backend/lib/api"

	"github.com/arangodb/go-driver"
	"github.com/dgrijalva/jwt-go"
	"github.com/BillotP/renty/backend/lib"
	"github.com/BillotP/renty/backend/lib/v2/models"
	"github.com/BillotP/renty/backend/lib/v2/service"
)

// Handle a serverless request
func Handle(w http.ResponseWriter, r *http.Request) {
	var err error
	var response []byte
	var user models.UserItem
	var claims jwt.StandardClaims
	var ctx = context.Background()
	var dbservice *service.Repository
	var dbName = goscrappy.MustGetSecret("arango_dbname")
	w.Header().Add("Content-Type", "application/json")
	if dbservice, err = service.New(dbName); err != nil {
		api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusOK)
		return
	} else if goscrappy.Debug {
		fmt.Printf("Token : %s\n", tokenString)
	}
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	if _, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		// Verify 'aud' claim
		aud := "myapi"
		if claims.Audience != aud {
			return token, errors.New("invalid audience")
		}
		// Verify 'iss' claim
		iss := "https://api.192.168.1.20.nip.io/"
		if claims.Issuer != iss {
			return token, errors.New("invalid issuer")
		}

		cert, err := goscrappy.GetSecret("jwt-public-key")
		if err != nil {
			return token, errors.New("server error")
		}

		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	}); err != nil {
		api.OPENFAASErrorResponse(w, http.StatusUnauthorized, errors.New("bad token"))
		return
	}
	if goscrappy.Debug {
		fmt.Printf("Claims : %+v\n", claims)
	}
	q := `
	FOR user IN users
	FILTER user.pseudo == @pseudo
	RETURN user
	`
	cursor, err := dbservice.Db.Query(ctx, q, map[string]interface{}{
		"pseudo": claims.Subject,
	})
	if err != nil {
		api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer cursor.Close()
	_, err = cursor.ReadDocument(ctx, &user)
	if driver.IsNoMoreDocuments(err) {
		api.OPENFAASErrorResponse(w, http.StatusUnauthorized, errors.New("invalid user"))
		return
	} else if err != nil {
		api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	if response, err = json.Marshal(user); err != nil {
		api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Add("X-Auth-User", string(response))
	w.WriteHeader(http.StatusOK)
}
