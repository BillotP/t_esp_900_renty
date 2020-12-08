package resolvers_test

import (
	"errors"
	"github.com/99designs/gqlgen/client"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMutationResolver_SignupAsCompany(t *testing.T) {
	var (
		errCompanyExists error
		query            string

		input  *models.CompanyInput
		output struct {
			SignupAsCompany models.Credential
		}
		expectedId int64 = 1

		err error
	)

	errCompanyExists = errors.New("[{\"message\":\"company seems already register\",\"path\":[\"signupAsCompany\"]}]")
	query = `mutation signupAsCompany($input: CompanyInput!){signupAsCompany(input: $input){user{ID,username}}}`
	input = &models.CompanyInput{
		Name:        "Thales",
		Description: "Thales est un groupe d'électronique spécialisé dans l'aérospatiale",
		Tel:         "6101010101",
		User: &models.UserInput{
			Username: "thalesadmin",
			Password: "thales1234",
		},
	}

	t.Run("should signup company correctly", func(t *testing.T) {
		err = Server.Post(query, &output, client.Var("input", input))

		require.Equal(t, &expectedId, output.SignupAsCompany.User.ID)
		require.Equal(t, "thalesadmin", output.SignupAsCompany.User.Username)
	})

	t.Run("should provide company already register error", func(t *testing.T) {
		err = Server.Post(query, &output, client.Var("input", &input))

		require.Equal(t, errCompanyExists.Error(), err.Error())
	})
}
