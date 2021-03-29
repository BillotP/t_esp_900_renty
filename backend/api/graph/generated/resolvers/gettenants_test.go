package resolvers_test

import (
	"regexp"
	"testing"

	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/middleware"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestQueryResolver_Tenants(t *testing.T) {
	var (
		query  string
		output struct {
			GetTenants []models.Tenant
		}
		// expectedId int64 = 1
		// err error
	)

	query = `query getTenants{ tenants {ID} }`

	middleware.InitMockDB(models.RoleEstateAgent)
	// errRecordNotFound := `[{"message":"record not found","path":["tenants"]}]`
	t.Run("should get tenants if exist or record not found otherwise", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"tenants\"")).
			WillReturnRows(sqlmock.NewRows([]string{"id"}))
		_ = middleware.Server.Post(query, &output)
		// require.Equal(t, errRecordNotFound, err.Error())
	})
}
