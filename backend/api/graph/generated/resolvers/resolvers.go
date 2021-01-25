package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"

	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/exec"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm"
)

type Resolver struct {
	DB *gorm.DB
}

func getPseudoRandomCost() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn((20 - bcrypt.MinCost) + bcrypt.MinCost)
}

func createToken(username string, userRole models.Role) (string, error) {
	var (
		token = ""

		err error
	)

	atClaims := jwt.MapClaims{}

	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["userrole"] = userRole.String()
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	if token, err = at.SignedString([]byte(lib.ServerConf.JwtSigningKey)); err != nil {
		return "", err
	}
	return token, nil
}

// Mutation returns exec.MutationResolver implementation.
func (r *Resolver) Mutation() exec.MutationResolver { return &MutationResolver{r} }

// Query returns exec.QueryResolver implementation.
func (r *Resolver) Query() exec.QueryResolver { return &QueryResolver{r} }

type MutationResolver struct{ *Resolver }
type QueryResolver struct{ *Resolver }
