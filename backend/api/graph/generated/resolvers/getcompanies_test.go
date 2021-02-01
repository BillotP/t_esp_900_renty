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

func TestQueryResolver_GetCompanies(t *testing.T) {

	var (
		query             string
		input  *models.CompanyInput
		output struct {
			GetCompany models.Company
		}
		expectedId int64 = 1
		err error
	)

	query = `mutation getCompanies($input: CompanyInput!){getCompanies(input: $input){ID, Name}}`
	input = &models.CompanyInput{
		Name: "test",
	}

	middleware.InitMockDB(models.RoleEstateAgent)

	t.Run("should get companies", func(t *testing.T) {

		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"companies\"")).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		err = middleware.Server.Post(query, &output, client.Var("input", &input))
		require.Equal(t, "", err.Error())
	})
}
