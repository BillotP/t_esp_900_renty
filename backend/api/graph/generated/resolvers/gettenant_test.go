package resolvers_test

import (
	"regexp"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/middleware"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestQueryResolver_GetTenant(t *testing.T) {
	var (
		query  string
		output struct {
			GetTenant models.Tenant
		}
		expectedId int64 = 1
		// err error
	)

	query = `query getTenant($id: Int!){ tenant(id: $id) {ID} }`

	middleware.InitMockDB(models.RoleEstateAgent)
	errRecordNotFound := `[{"message":"record not found","path":["tenant"]}]`
	t.Run("should get tenant id 1 if exist or record not found otherwise", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"tenants\" WHERE id = $1")).
			WithArgs(expectedId).
			WillReturnRows(sqlmock.NewRows([]string{"id"}))
		err := middleware.Server.Post(query, &output, client.Var("id", expectedId))
		require.Equal(t, errRecordNotFound, err.Error())
	})
}
