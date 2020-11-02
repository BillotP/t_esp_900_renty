package function

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/arangodb/go-driver"
	"github.com/dgrijalva/jwt-go"
	"github.com/BillotP/gorenty"
	"github.com/BillotP/gorenty/api"
	"github.com/BillotP/gorenty/v2/models"
	"github.com/BillotP/gorenty/v2/service"
)

// Config is the values for jwt token signing
type Config struct {
	Issuer   string          `json:"issuer"`
	Audience string          `json:"audience"`
	KeyID    string          `json:"key_id"`
	Expiracy time.Duration   `json:"expiracy"`
	SignKey  *rsa.PrivateKey `json:"-"`
}

// AuthQuery is the authentication payload
type AuthQuery struct {
	Pseudo   string `json:"pseudo"`
	Password string `json:"password"`
}

// AuthResponse is the authentication response
type AuthResponse struct {
	Token *string          `json:"token"`
	User  *models.UserItem `json:"user"`
	Error *string          `json:"error"`
}

var conf *Config
var dbservice *service.Repository
var (
	// ErrUnknowUser is returned when an unknow user is in the jwt
	ErrUnknowUser = errors.New("unknow user, register first")
	// ErrBadCredential is returned when submitted password is invalid
	ErrBadCredential = errors.New("bad credential")
)

func init() {
	var err error
	var signKey *rsa.PrivateKey
	var expiracy time.Duration
	var dbname = goscrappy.MustGetSecret("arango_dbname")
	if dbservice, err = service.New(dbname); err != nil {
		panic(err.Error())
	}
	signString := goscrappy.MustGetSecret("jwt_private_key")
	if signKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(signString)); err != nil {
		panic(err.Error())
	}
	jwtIssuer := goscrappy.MustGetSecret("jwt_issuer")
	jwtAudience := goscrappy.MustGetSecret("jwt_audience")
	jwtKeyID := goscrappy.MustGetSecret("jwt_key_id")
	jwtExpiracy := goscrappy.MustGetSecret("jwt_expiracy")
	if expiracy, err = time.ParseDuration(jwtExpiracy); err != nil {
		panic(err.Error())
	}
	conf = &Config{
		Issuer:   jwtIssuer,
		Audience: jwtAudience,
		KeyID:    jwtKeyID,
		Expiracy: expiracy,
		SignKey:  signKey,
	}
}

// CreateToken return a signed jwt token string
func (c *Config) CreateToken(user string) (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.SigningMethodRS256)
	// set our claims
	t.Claims = &jwt.StandardClaims{
		// set the expire time
		// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
		ExpiresAt: time.Now().Add(c.Expiracy).Unix(),
		Issuer:    c.Issuer,
		Audience:  c.Audience,
		Subject:   user,
	}
	t.Header["kid"] = c.KeyID
	// Create token string
	return t.SignedString(c.SignKey)
}

// Handle a serverless request
func Handle(w http.ResponseWriter, r *http.Request) {
	var err error
	var token string
	var query AuthQuery
	var user models.UserItem
	var ctx = context.Background()
	api.OPENFAASGetBody(r, &query)
	if goscrappy.Debug {
		fmt.Printf("Query : %+v\n", query)
	}
	q := `
	FOR user IN users
	FILTER user.pseudo == @pseudo
	RETURN user
	`
	cursor, err := dbservice.Db.Query(ctx, q, map[string]interface{}{
		"pseudo": query.Pseudo,
	})
	if err != nil {
		api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer cursor.Close()
	_, err = cursor.ReadDocument(ctx, &user)
	if driver.IsNoMoreDocuments(err) {
		msg := ErrUnknowUser.Error()
		api.OPENFAASJsonResponse(w, http.StatusOK, AuthResponse{
			Error: &msg,
		})
		return
	} else if err != nil {
		api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(query.Password)); err != nil {
		msg := ErrBadCredential.Error()
		api.OPENFAASJsonResponse(w, http.StatusOK, AuthResponse{
			Error: &msg,
		})
		return
	}
	if token, err = conf.CreateToken(user.Pseudo); err != nil {
		api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	api.OPENFAASJsonResponse(w, http.StatusOK, AuthResponse{
		Token: &token,
		User:  &user,
	})
}
