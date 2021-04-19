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

func TestMutationResolver_CreateAnomaly(t *testing.T) {
	var (
		query string

		input  *models.AnomalyInput
		output struct {
			CreateAnomaly models.Anomaly
		}
		expectedId int64 = 1

		err error

		propertyTest    int64 = 1
		typeTest              = "Damages"
		descriptionTest       = ""
	)

	middleware.InitMockDB(models.RoleAdmin)

	query = `mutation createAnomaly($input: AnomalyInput!){createAnomaly(input: $input){ID}}`
	input = &models.AnomalyInput{
		Property:    &propertyTest,
		Type:        models.AnomalyTypes(typeTest),
		Description: descriptionTest,
	}

	t.Run("should create anomaly successfully", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(".+").
			WillReturnRows(sqlmock.NewRows(nil))
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("INSERT INTO \"anomalies\" (\"created_at\",\"updated_at\",\"property\",\"type\",\"description\") VALUES ($1,$2,$3,$4,$5) RETURNING \"id\"")).
			WithArgs(AnyTime{}, AnyTime{}, input.Property, input.Type, input.Description).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		err = middleware.Server.Post(query, &output, client.Var("input", input))

		description := output.CreateAnomaly.Description
		require.Equal(t, descriptionTest, description)
		require.NotNil(t, err.Error())
	})
}
