package resolvers_test

import (
	"regexp"
	"testing"

	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/middleware"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestQueryResolver_Companies(t *testing.T) {

	var (
		query  string
		output struct {
			GetCompanies []models.Company
		}
		// expectedId int64 = 1
		// err error
	)

	query = `query getCompanies{ companies {ID, name} }`

	middleware.InitMockDB(models.RoleAdmin)

	t.Run("should get all companies", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"companies\"")).
			WillReturnRows(sqlmock.NewRows([]string{"id"}))
		_ = middleware.Server.Post(query, &output)
		// if err != nil {
		// 	t.Fatal(err)
		// }
		require.Equal(t, 0, len(output.GetCompanies))
		// require.Equal(t, nil, err)
	})
}
