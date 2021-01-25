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

func TestMutationResolver_CreateEstateAgentUser(t *testing.T) {
	var (
		errEstateAgentExists error
		query                string

		input  *models.EstateAgentInput
		output struct {
			CreateEstateAgentUser models.Credential
		}

		expectedId int64 = 1

		err error
	)

	middleware.InitMockDB(models.RoleCompany)

	errEstateAgentExists = errors.New("[{\"message\":\"estate agent seems already register\",\"path\":[\"createEstateAgentUser\"]}]")
	query = `mutation createEstateAgentUser($input: EstateAgentInput){createEstateAgentUser(input: $input){user{ID,username}}}`
	input = &models.EstateAgentInput{
		User: &models.UserInput{
			Username: "toto@foncia.com",
			Password: "Toto1234",
		},
	}

	t.Run("should create estate agent user correctly", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"companies\" ORDER BY \"companies\".\"id\" LIMIT 1")).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		middleware.Mock.
			ExpectQuery(".+").
			WillReturnRows(sqlmock.NewRows(nil))
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("INSERT INTO \"users\" (\"created_at\",\"updated_at\",\"username\",\"password\",\"role\") VALUES ($1,$2,$3,$4,$5) ON CONFLICT DO NOTHING RETURNING \"id\"")).
			WithArgs(AnyTime{}, AnyTime{}, input.User.Username, AvoidPassword{}, models.RoleEstateAgent.String()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("INSERT INTO \"estate_agents\" (\"created_at\",\"updated_at\",\"company_id\",\"user_id\") VALUES ($1,$2,$3,$4) RETURNING \"id\"")).
			WithArgs(AnyTime{}, AnyTime{}, expectedId, expectedId).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		err = middleware.Server.Post(query, &output, client.Var("input", input))

		require.Equal(t, &expectedId, output.CreateEstateAgentUser.User.ID)
		require.Equal(t, "toto@foncia.com", output.CreateEstateAgentUser.User.Username)
	})

	t.Run("should provide estate agent user already register error", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"companies\" ORDER BY \"companies\".\"id\" LIMIT 1")).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("SELECT")).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		err = middleware.Server.Post(query, &output, client.Var("input", &input))
		require.Equal(t, errEstateAgentExists.Error(), err.Error())
	})
}