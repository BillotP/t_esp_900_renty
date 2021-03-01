package resolvers_test

import (
	"errors"
	"github.com/99designs/gqlgen/client"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/middleware"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"testing"
	"time"
)

func TestMutationResolver_LoginAsEstateAgent(t *testing.T) {
	var (
		errBadCredentials error
		query             string

		input        *models.UserInput
		passwordHash string
		output       struct {
			LoginAsEstateAgent models.Credential
		}

		expectedId int64 = 1

		err error
	)

	middleware.InitMockDB(models.RoleCompany)

	errBadCredentials = errors.New("[{\"message\":\"bad password provided\",\"path\":[\"loginAsEstateAgent\"]}]")
	query = `mutation loginAsEstateAgent($input: UserInput){loginAsEstateAgent(input: $input){user{ID,username}}}`
	input = &models.UserInput{
		Username: "toto@foncia.com",
		Password: "Toto1234Toto1234",
	}
	rand.Seed(time.Now().UnixNano())
	cost := rand.Intn((20 - bcrypt.MinCost) + bcrypt.MinCost)
	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), cost)
	passwordHash = string(hash)
	t.Run("should login estate agent correctly", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows([]string{"User__id", "User__username", "User__password"}).AddRow(expectedId, input.Username, passwordHash))
		err = middleware.Server.Post(query, &output, client.Var("input", input))

		require.Equal(t, &expectedId, output.LoginAsEstateAgent.User.ID)
		require.Equal(t, "toto@foncia.com", output.LoginAsEstateAgent.User.Username)
	})

	t.Run("should provide bad credentials", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows([]string{"id", "User__password"}).AddRow(expectedId, input.Password))
		err = middleware.Server.Post(query, &output, client.Var("input", input))

		require.Equal(t, errBadCredentials.Error(), err.Error())
	})
}
