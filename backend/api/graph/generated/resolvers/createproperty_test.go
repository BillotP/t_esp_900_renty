package resolvers

import (
	"errors"
	"github.com/99designs/gqlgen/client"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMutationResolver_SignupAsCompany(t *testing.T) {
	var (
		errPropertyExists error
		query            string

		input  *models.PropertyInput
		output struct {
			CreateProperty models.Credential
		}
		expectedId int64 = 1

		err error
	)

	errPropertyExists = errors.New("[{\"message\":\"property seems to be already registered\",\"path\":[\"createProperty\"]}]")
	query = `mutation createProperty($input: PropertyInput!){createProperty(input: $input){user{ID,address}}}`
	var (
		aeraTest float64 = 123
		addressTest string = "1 avenue Test, 33000, Bordeaux, apt 104B"
		codeNumberTest int64 = 33000
		typeTest string = "T3"
	)
	input = &models.PropertyInput{
		Area:		&aeraTest,
		Address:	&addressTest,
		CodeNumber:	&codeNumberTest,
		Type: 		&typeTest,
	}


	t.Run("should create property correctly", func(t *testing.T) {
		err = Server.Post(query, &output, client.Var("input", input))

		require.Equal(t, &expectedId, output.CreateProperty.User.ID)
		require.Equal(t, "thalesadmin", output.CreateProperty.User.Username)
	})

	t.Run("should raise property already registered error", func(t *testing.T) {
		err = Server.Post(query, &output, client.Var("input", &input))

		require.Equal(t, errPropertyExists.Error(), err.Error())
	})
}
