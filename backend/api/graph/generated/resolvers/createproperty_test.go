package resolvers_test

import (
	"errors"
	"github.com/99designs/gqlgen/client"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/middleware"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"log"
	"regexp"
	"testing"
)

func TestMutationResolver_CreateProperty(t *testing.T) {
	var (
		errPropertyExists error
		query            string

		input  *models.PropertyInput
		output struct {
			CreateProperty models.Property
		}
		expectedId int64 = 1

		err error

		testUser = &models.UserInput{
			Username: "thalesadmin",
			Password: "thales1234",
		}

		aeraTest float64 = 123
		addressTest string = "1 avenue Test, 33000, Bordeaux, apt 104B"
		codeNumberTest int64 = 33000
		typeTest string = "T3"
	)

	errPropertyExists = errors.New("[{\"message\":\"property seems to be already registered\",\"path\":[\"createProperty\"]}]")
	query = `mutation createProperty($input: PropertyInput!){createProperty(input: $input){ID,address}}`
	input = &models.PropertyInput {
		Area:		&aeraTest,
		Address:	&addressTest,
		CodeNumber:	&codeNumberTest,
		Type: 		&typeTest,
	}
	*input.Area = aeraTest
	*input.Address = addressTest
	*input.CodeNumber = codeNumberTest
	*input.Type = typeTest
	log.Print(">>>> input = ", input)
	log.Print(">>>> &input = ", &input)
	log.Print(">>>> *input = ", *input)
	log.Print(">>>> input.Address = ", input.Address)




	t.Run("should create property successfully", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(".+").
			WillReturnRows(sqlmock.NewRows(nil))
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("INSERT INTO \"users\" (\"created_at\",\"updated_at\",\"username\",\"password\",\"role\") VALUES ($1,$2,$3,$4,$5) ON CONFLICT DO NOTHING RETURNING \"id\"")).
			WithArgs(AnyTime{}, AnyTime{}, testUser.Username, AvoidPassword{}, models.RoleEstateAgent.String()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("INSERT INTO \"properties\" (\"created_at\",\"updated_at\",\"address\",\"type\",\"area\",\"codeNumber\") VALUES ($1,$2,$3,$4,$5,$6) RETURNING \"id\"")).
			WithArgs(AnyTime{}, AnyTime{}, input.Address, input.Type, input.Area, input.CodeNumber).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		err = middleware.Server.Post(query, &output, client.Var("input", &input))

		require.Equal(t, &expectedId, output.CreateProperty.ID)
		address := output.CreateProperty.Address
		require.Equal(t, addressTest, address)
	})

	t.Run("should raise property already registered error", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"properties\" WHERE name = $1 LIMIT 1")).
			WithArgs(input.Address).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		err = middleware.Server.Post(query, &output, client.Var("input", &input))

		require.Equal(t, errPropertyExists.Error(), err.Error())
	})
}
