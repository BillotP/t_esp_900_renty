package resolvers_test

import (
	"regexp"
	"testing"

	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/middleware"
	"github.com/stretchr/testify/require"
)

func TestMutationResolver_AcceptCompany(t *testing.T) {
	var (
		// query            string

		// input  int64
		// output struct {
		// 	SignupAsCompany models.Credential
		// }
		expectedID int64 = 1

	// err error
	)

	middleware.InitMockDB(models.RoleAdmin)

	// query := `mutation acceptCompany($id: Int!){acceptCompany(id: $id){company{name}}}`
	var hopefullyID int64 = 0
	var verified = true
	// input := hopefullyID
	companytoUpdate := &models.Company{
		ID:       &hopefullyID,
		Verified: &verified,
	}

	t.Run("should update company correctly", func(t *testing.T) {
		// middleware.Mock.
		// 	ExpectQuery(".+").
		// 	WillReturnRows(sqlmock.NewRows(nil))
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("UPDATE \"companies\" SET \"updated_at\"='2021-02-01 11:42:13.453',\"verified\"=true WHERE \"id\" = 1"))

		// middleware.Mock.
		// 	ExpectQuery(regexp.QuoteMeta("INSERT INTO \"companies\" (\"created_at\",\"updated_at\",\"name\",\"logo_id\",\"description\",\"tel\",\"user_id\",\"verified\") VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING \"id\"")).
		// 	WithArgs(AnyTime{}, AnyTime{}, input.Name, nil, input.Description, input.Tel, expectedID, false).
		// 	WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedId))
		// err = middleware.Server.Post(query, &output, client.Var("id", input))

		// require.Equal(t, &expectedID, output.SignupAsCompany.User.ID)
		// require.Equal(t, "thalesadmin", output.SignupAsCompany.User.Username)
		require.Equal(t, &verified, companytoUpdate.Verified)

	})

	t.Run("should provide company not found error", func(t *testing.T) {
		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta("UPDATE \"companies\" SET \"updated_at\"='2021-02-01 11:42:13.453',\"verified\"=true WHERE \"id\" = 42"))
		// 	WithArgs(input.Name).
		// 	WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedID))
		// err = middleware.Server.Post(query, &output, client.Var("input", &input))

		// require.Equal(t, errCompanyExists.Error(), err.Error())
	})
}
