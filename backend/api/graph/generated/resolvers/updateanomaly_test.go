package resolvers_test

import (
	"github.com/99designs/gqlgen/client"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/middleware"
	"regexp"
	"testing"
)

func TestMutationResolver_UpdateAnomaly(t *testing.T) {
	// var err error
	query := `mutation updateAnomaly($input: TenantUpdateInput!){updateAnomaly(id: $id){ID}}`
	// var hopefullyID int64 = 1
	var output struct {
		updateAnomaly models.Anomaly
	}
	anomalyToUpdate := &models.AnomalyUpdateInput{
	}

	middleware.InitMockDB(models.RoleAdmin)
	t.Run("should update anomaly correctly", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta(`UPDATE "anomalies" SET "updated_at"=$1,"state"=$2`)).
			WithArgs(AnyTime{}, anomalyToUpdate.State)
		_ = middleware.Server.Post(query, &output, client.Var("input", anomalyToUpdate))
		t.Logf("Output : %+v\n", output)
	})
}
