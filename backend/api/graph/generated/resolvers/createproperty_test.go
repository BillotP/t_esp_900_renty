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

func TestMutationResolver_CreateProperty(t *testing.T) {
	var (
		errPropertyExists error
		query             string

		input  *models.PropertyInput
		output struct {
			CreateProperty models.Property
		}
		expectedId int64 = 1

		err error

		areaTest       float64 = 123
		addressTest            = "1 avenue Test, 33000, Bordeaux, apt 104B"
		codeNumberTest int64   = 33000
		typeTest               = "T3"
	)

	errPropertyExists = errors.New("[{\"message\":\"there is already a property at this address\",\"path\":[\"createProperty\"]}]")
	query = `mutation createProperty($input: PropertyInput!){createProperty(input: $input){ID,address}}`
	input = &models.PropertyInput{
		Area:       &areaTest,
		Address:    &addressTest,
		CodeNumber: &codeNumberTest,
		Type:       &typeTest,
	}

	t.Run("should create property successfully", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(".+").
			WillReturnRows(sqlmock.NewRows(nil))
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("INSERT INTO \"properties\" (\"created_at\",\"updated_at\",\"area\",\"address\",\"code_number\",\"type\") VALUES ($1,$2,$3,$4,$5,$6) RETURNING \"id\"")).
			WithArgs(AnyTime{}, AnyTime{}, input.Area, input.Address, input.CodeNumber, input.Type).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		err = middleware.Server.Post(query, &output, client.Var("input", input))

		address := output.CreateProperty.Address
		require.Equal(t, addressTest, *address)
	})

	t.Run("should raise property already registered error", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"properties\" WHERE address = $1 ORDER BY \"properties\".\"id\" LIMIT 1")).
			WithArgs(input.Address).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		err = middleware.Server.Post(query, &output, client.Var("input", &input))

		require.Equal(t, errPropertyExists.Error(), err.Error())
	})
}