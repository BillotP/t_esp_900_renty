package resolvers_test

import (
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/middleware"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func TestQueryResolver_Anomalies(t *testing.T) {
	var (
		query  string
		output struct {
			GetAnomalies []models.Anomaly
		}
		// err error
	)

	query = `query getAnomalies{ anomalies {ID} }`

	middleware.InitMockDB(models.RoleEstateAgent)
	errRecordNotFound := `[{"message":"record not found","path":["anomaly"]}]`
	t.Run("should get anomalies if exist or record not found otherwise", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"anomalies\"")).
			WillReturnRows(sqlmock.NewRows([]string{"id"}))
		err := middleware.Server.Post(query, &output)
		require.NotEqual(t, errRecordNotFound, err.Error())
	})
}
