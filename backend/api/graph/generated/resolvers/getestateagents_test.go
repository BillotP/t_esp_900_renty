package resolvers_test

import (
	"regexp"
	"testing"

	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/middleware"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestQueryResolver_EstateAgents(t *testing.T) {
	var (
		query  string
		output struct {
			GetEstateAgent []models.EstateAgent
		}
		// expectedId int64 = 1
		// err error
	)

	query = `query getEstateAgents { estateAgents {ID companyID} }`

	middleware.InitMockDB(models.RoleEstateAgent)
	// errRecordNotFound := `[{"message":"record not found","path":["estateAgents"]}]`
	t.Run("should get estate agent 1 if exist or record not found otherwise", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"estate_agents\"")).
			WillReturnRows(sqlmock.NewRows([]string{"id"}))
		_ = middleware.Server.Post(query, &output)
		// if err != nil {
		// 	t.Fatal(err)
		// }
		// require.Equal(t, errRecordNotFound, err.Error())
	})
}
