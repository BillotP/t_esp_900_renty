package resolvers_test

import (
	"regexp"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/middleware"
	"github.com/stretchr/testify/require"
)

func TestMutationResolver_AcceptCompany(t *testing.T) {
	// var err error
	query := `mutation acceptCompany($id: Int!){acceptCompany(id: $id){verified}}`
	var verified = true
	var hopefullyID int64 = 1
	var output struct {
		acceptCompany models.Company
	}
	companytoUpdate := &models.Company{
		ID:       &hopefullyID,
		Verified: &verified,
	}

	middleware.InitMockDB(models.RoleAdmin)
	t.Run("should update company correctly", func(t *testing.T) {
		// middleware.Mock.
		// 	ExpectQuery(".+").
		// 	WillReturnRows(sqlmock.NewRows(nil))

		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta(`UPDATE "companies" SET "updated_at"=$1,"verified"=$2 WHERE "id" = $3`)).
			WithArgs(AnyTime{}, true, hopefullyID)
		_ = middleware.Server.Post(query, &output, client.Var("id", hopefullyID))
		t.Logf("Output : %+v\n", output)
		// if err != nil {
		// 	t.Fatal(err)
		// }
		require.Equal(t, &verified, companytoUpdate.Verified)

	})
}
