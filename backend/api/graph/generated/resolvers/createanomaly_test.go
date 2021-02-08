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

func TestMutationResolver_CreateAnomaly(t *testing.T) {
	var (
		errAnomalyExists error
		query            string

		input  *models.AnomalyInput
		output struct {
			CreateAnomaly models.Anomaly
		}
		expectedId int64 = 1

		err error

		propertyTest int64 = 1
		typeTest = "Damages"
		descriptionTest = "VMC cuisine hors-service"
	)

	middleware.InitMockDB(models.RoleAdmin)

	errAnomalyExists = errors.New("[{\"message\":\"there is already an anomaly with this \",\"path\":[\"createAnomaly\"]}]")
	query = `mutation createAnomaly($input: AnomalyInput!){createAnomaly(input: $input){ID}}`
	input = &models.AnomalyInput{
		Property: &propertyTest,
		Type: typeTest,
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
	})

	t.Run("should raise anomaly already registered error", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"anomalies\"")).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		err = middleware.Server.Post(query, &output, client.Var("input", &input))

		require.Equal(t, errAnomalyExists.Error(), err.Error())
	})
}
