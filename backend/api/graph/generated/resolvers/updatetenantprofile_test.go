package resolvers_test

import (
	"regexp"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/middleware"
)

func TestMutationResolver_UpdateTenantProfile(t *testing.T) {
	// var err error
	query := `mutation updatetTenant($input: TenantUpdateInput!){updateTenantProfile(id: $id){ID}}`
	// var hopefullyID int64 = 1
	var output struct {
		updateTenantProfile models.Tenant
	}
	nprop := int64(0)
	fproperties := []*int64{&nprop}
	tenanttoUpdate := &models.TenantUpdateInput{
		Properties: fproperties,
	}

	middleware.InitMockDB(models.RoleAdmin)
	t.Run("should update tenant correctly", func(t *testing.T) {
		// middleware.Mock.
		// 	ExpectQuery(".+").
		// 	WillReturnRows(sqlmock.NewRows(nil))

		middleware.Mock.
			ExpectQuery(regexp.QuoteMeta(`UPDATE "tenants" SET "updated_at"=$1,"properties"=$2 WHERE "id" = $3`)).
			WithArgs(AnyTime{}, true, tenanttoUpdate.Properties)
		_ = middleware.Server.Post(query, &output, client.Var("input", tenanttoUpdate))
		t.Logf("Output : %+v\n", output)
		// if err != nil {
		// 	t.Fatal(err)
		// }
		// require.Equal(t, &verified, companytoUpdate.Verified)

	})
}
