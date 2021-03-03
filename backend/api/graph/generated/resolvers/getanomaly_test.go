package resolvers_test

import (
	"github.com/99designs/gqlgen/client"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/middleware"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func TestQueryResolver_GetAnomaly(t *testing.T) {
	var (
		query  string
		output struct {
			GetAnomaly models.Anomaly
		}
		expectedId int64 = 1
		// err error
	)

	query = `query getAnomaly($id: String!){ anomaly(id: $id) {ID} }`

	middleware.InitMockDB(models.RoleEstateAgent)
	errRecordNotFound := `[{"message":"record not found","path":["anomaly"]}]`
	t.Run("should get anomaly id 1 if exist or record not found otherwise", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"anomalies\" WHERE id = $1")).
			WithArgs(expectedId).
			WillReturnRows(sqlmock.NewRows([]string{"id"}))
		err := middleware.Server.Post(query, &output, client.Var("id", expectedId))
		require.NotEqual(t, errRecordNotFound, err.Error())
	})
}
