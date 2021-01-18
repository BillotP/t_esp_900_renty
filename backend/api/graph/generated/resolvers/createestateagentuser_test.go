package resolvers_test

import (
	"errors"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/middleware"
	"testing"
)

func init()  {
	middleware.InitMockDB()
}

func TestMutationResolver_CreateEstateAgentUser(t *testing.T) {
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

	print(errCompanyExists)
	print(query)
	print(input)
	print(output.SignupAsCompany.Token)
	print(expectedId)
	print(err)
}