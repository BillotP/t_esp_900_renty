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

func TestMutationResolver_CreateTenantUser(t *testing.T) {
	var (
		errTenantExists error
		query           string

		input  *models.TenantInput
		output struct {
			CreateTenantUser models.Credential
		}

		expectedId int64 = 1

		err error
	)

	middleware.InitMockDB(models.RoleEstateAgent)

	errTenantExists = errors.New("[{\"message\":\"tenant seems already register\",\"path\":[\"createTenantUser\"]}]")
	query = `mutation createTenantUser($input: TenantInput){createTenantUser(input: $input){user{ID,username}}}`
	input = &models.TenantInput{
		User: &models.UserInput{
			Username: "tenant@sfr.fr",
			Password: "Toto1234",
		},
	}

	t.Run("should create tenant user correctly", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"estate_agents\" ORDER BY \"estate_agents\".\"id\" LIMIT 1")).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		middleware.Mock.
			ExpectQuery(".+").
			WillReturnRows(sqlmock.NewRows(nil))
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("INSERT INTO \"users\" (\"created_at\",\"updated_at\",\"username\",\"password\",\"role\") VALUES ($1,$2,$3,$4,$5) ON CONFLICT DO NOTHING RETURNING \"id\"")).
			WithArgs(AnyTime{}, AnyTime{}, input.User.Username, AvoidPassword{}, models.RoleTenant.String()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("INSERT INTO \"tenants\" (\"created_at\",\"updated_at\",\"user_id\",\"estate_agent_id\") VALUES ($1,$2,$3,$4) RETURNING \"id\"")).
			WithArgs(AnyTime{}, AnyTime{}, expectedId, expectedId).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		err = middleware.Server.Post(query, &output, client.Var("input", input))

		require.Equal(t, &expectedId, output.CreateTenantUser.User.ID)
		require.Equal(t, "tenant@sfr.fr", output.CreateTenantUser.User.Username)
	})

	t.Run("should provide tenant user already register error", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"estate_agents\" ORDER BY \"estate_agents\".\"id\" LIMIT 1")).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("SELECT")).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		err = middleware.Server.Post(query, &output, client.Var("input", &input))
		require.Equal(t, errTenantExists.Error(), err.Error())
	})
}
