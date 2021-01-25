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

func TestMutationResolver_LoginAsCompany(t *testing.T) {
	var (
		errCompanyExists error
		query            string

		input  *models.CompanyInput
		output struct {
			SignupAsCompany models.Credential
		}
		expectedId int64 = 1

		err error
	)

	middleware.InitMockDB("admin", models.RoleAdmin)

	errCompanyExists = errors.New("[{\"message\":\"company seems already register\",\"path\":[\"signupAsCompany\"]}]")
	query = `mutation signupAsCompany($input: CompanyInput!){signupAsCompany(input: $input){user{ID,username}}}`
	input = &models.CompanyInput{
		Name:        "Thales",
		Description: "Thales est un groupe d'électronique spécialisé dans l'aérospatiale",
		Tel:         "6101010101",
		User: &models.UserInput{
			Username: "thalesadmin",
			Password: "thales1234",
		},
	}

	t.Run("should login company correctly", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(".+").
			WillReturnRows(sqlmock.NewRows(nil))
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("INSERT INTO \"users\" (\"created_at\",\"updated_at\",\"username\",\"password\",\"role\") VALUES ($1,$2,$3,$4,$5) ON CONFLICT DO NOTHING RETURNING \"id\"")).
			WithArgs(AnyTime{}, AnyTime{}, input.User.Username, AvoidPassword{}, models.RoleCompany.String()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("INSERT INTO \"companies\" (\"created_at\",\"updated_at\",\"name\",\"logo_id\",\"description\",\"tel\",\"user_id\",\"verified\") VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING \"id\"")).
			WithArgs(AnyTime{}, AnyTime{}, input.Name, nil, input.Description, input.Tel, expectedId, false).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		err = middleware.Server.Post(query, &output, client.Var("input", input))

		require.Equal(t, &expectedId, output.SignupAsCompany.User.ID)
		require.Equal(t, "thalesadmin", output.SignupAsCompany.User.Username)
	})

	t.Run("should provide company login error", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"companies\" WHERE name = $1 ORDER BY \"companies\".\"id\" LIMIT 1")).
			WithArgs(input.Name).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		err = middleware.Server.Post(query, &output, client.Var("input", &input))

		require.Equal(t, errCompanyExists.Error(), err.Error())
	})
}
