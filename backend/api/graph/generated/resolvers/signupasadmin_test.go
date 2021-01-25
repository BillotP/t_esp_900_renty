package resolvers_test

import (
	"errors"
	"github.com/99designs/gqlgen/client"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/middleware"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func TestMutationResolver_SignupAsAdmin(t *testing.T) {
	var (
		errAdminUserExists	error
		query				string

		input  *models.UserInput
		output struct {
			SignupAsAdmin models.Credential
		}
		expectedId int64 = 1

		err error
	)

	middleware.InitMockDB(models.RoleAdmin)

	errAdminUserExists = errors.New("[{\"message\":\"user seems already register\",\"path\":[\"signupAsAdmin\"]}]")
	query = `mutation signupAsAdmin($input: UserInput!){signupAsAdmin(input: $input){user{ID,username}}}`
	input = &models.UserInput{
			Username: "adminusertest",
			Password: "aut1234",
	}

	t.Run("should signup admin user correctly", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(".+").
			WillReturnRows(sqlmock.NewRows(nil))
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("INSERT INTO \"users\" (\"created_at\",\"updated_at\",\"username\",\"password\",\"role\") VALUES ($1,$2,$3,$4,$5) ON CONFLICT DO NOTHING RETURNING \"id\"")).
			WithArgs(AnyTime{}, AnyTime{}, input.Username, AvoidPassword{}, models.RoleAdmin.String()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		err = middleware.Server.Post(query, &output, client.Var("input", input))

		require.Equal(t, &expectedId, output.SignupAsAdmin.User.ID)
		require.Equal(t, "adminusertest", output.SignupAsAdmin.User.Username)
	})

	t.Run("should provide admin user already register error", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"users\" WHERE username = $1 ORDER BY \"users\".\"id\" LIMIT 1")).
			WithArgs(input.Username).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		err = middleware.Server.Post(query, &output, client.Var("input", &input))

		require.Equal(t, errAdminUserExists.Error(), err.Error())
	})
}
