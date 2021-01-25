package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"

	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/exec"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
)

type Resolver struct {
	DB *gorm.DB
}

func getPseudoRandomCost() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn((20 - bcrypt.MinCost) + bcrypt.MinCost)
}

func createToken(username string, userRole models.Role) (string, error) {
	var (
		token = ""

		err error
	)

	atClaims := jwt.MapClaims{}

	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["userrole"] = userRole.String()
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	if token, err = at.SignedString([]byte(lib.ServerConf.JwtSigningKey)); err != nil {
		return "", err
	}
	return token, nil
}

func (r *MutationResolver) SignupAsAdmin(ctx context.Context, input models.AdminInput) (*models.Credential, error) {
	panic("not implemented")
}

func (r *MutationResolver) AcceptCompany(ctx context.Context, id int64) (*models.Company, error) {
	var (
		company  *models.Company
		verified = true

		err error
	)

	company = &models.Company{
		ID:       &id,
		Verified: &verified,
	}
	if err = r.DB.Updates(&company).Error; err == nil {
		return company, nil
	}
	lib.LogError("mutation/AcceptCompany", err.Error())
	return nil, err
}

func (r *MutationResolver) UpdateTenantProfile(ctx context.Context, input *models.TenantUpdateInput) (*models.Tenant, error) {
	panic("Wilfried : not implemented")
}

func (r *MutationResolver) CreateAnomaly(ctx context.Context, input *models.AnomalyInput) (*models.Anomaly, error) {
	var (
		anomaly *models.Anomaly
		err error
	)
	anomaly = &models.Anomaly{
		PropertyID:     input.Property,
		Type: 			input.Type,
		Description:	input.Description,
	}

	if err = r.DB.Where(&anomaly).First(&anomaly).Error; err == nil {
		return nil, fmt.Errorf("anomaly already created")
	}
	if err = r.DB.Create(&anomaly).Error; err != nil {
		lib.LogError("mutation/Register/Anomaly", err.Error())
		return nil, err
	}
	return anomaly, nil
}

func (r *MutationResolver) UpdateAnomaly(ctx context.Context, input *models.AnomalyUpdateInput) (*models.Anomaly, error) {
	var (
		anomaly *models.Anomaly
		err error
	)
	anomaly = &models.Anomaly{
		AssignedToID:   input.AssignedTo,
		State: 			input.State,
	}
	if err = r.DB.Updates(&anomaly).First(&anomaly).Error; err != nil {
		lib.LogError("mutation/Register/Anomaly", err.Error())
		return nil, err
	}
	return anomaly, nil
}

func (r *QueryResolver) Anomaly(ctx context.Context, id string) (*models.Anomaly, error) {
	var (
		anomaly *models.Anomaly
		err error
	)

	if err = r.DB.Where("ID = ?", id).First(&anomaly).Error; err != nil {
		return nil, fmt.Errorf("anomaly not found")
	}
	return anomaly, nil
}

func (r *QueryResolver) Anomalies(ctx context.Context) ([]*models.Anomaly, error) {
	var (
		anomalies []*models.Anomaly
		err error
	)

	if err = r.DB.Find(&anomalies).Error; err != nil {
		return nil, fmt.Errorf("any anomaly found")
	}
	return anomalies, nil
}

func (r *QueryResolver) Tenant(ctx context.Context, id int64) (*models.Tenant, error) {
	var (
		tenant models.Tenant

		err error
	)

	if err = r.DB.Preload(clause.Associations).Where("id = ?", id).First(&tenant).Error; err == nil {
		return &tenant, nil
	}
	lib.LogError("mutation/Tenant", err.Error())
	return nil, err
}

func (r *QueryResolver) Tenants(ctx context.Context) ([]*models.Tenant, error) {
	var (
		tenants []models.Tenant

		err error
	)

	if err = r.DB.Preload(clause.Associations).Find(&tenants).Error; err == nil {
		var tenantsfmt []*models.Tenant

		for i := range tenants {
			tenantsfmt = append(tenantsfmt, &tenants[i])
		}
		return tenantsfmt, nil
	}
	lib.LogError("mutation/Tenants", err.Error())
	return nil, err
}

func (r *QueryResolver) EstateAgent(ctx context.Context, id int64) (*models.EstateAgent, error) {
	var (
		estateAgent models.EstateAgent

		err error
	)

	if err = r.DB.Preload(clause.Associations).Where("id = ?", id).First(&estateAgent).Error; err == nil {
		return &estateAgent, nil
	}
	lib.LogError("mutation/EstateAgent", err.Error())
	return nil, err
}

func (r *QueryResolver) EstateAgents(ctx context.Context) ([]*models.EstateAgent, error) {
	var (
		estateAgents []models.EstateAgent

		err error
	)

	if err = r.DB.Preload(clause.Associations).Find(&estateAgents).Error; err == nil {
		var estateagentsfmt []*models.EstateAgent

		for i := range estateAgents {
			estateagentsfmt = append(estateagentsfmt, &estateAgents[i])
		}
		return estateagentsfmt, nil
	}
	lib.LogError("mutation/EstateAgents", err.Error())
	return nil, err
}

func (r *QueryResolver) Company(ctx context.Context, id int64) (*models.Company, error) {
	var (
		company models.Company

		err error
	)

	if err = r.DB.Preload(clause.Associations).Where("id = ?", id).First(&company).Error; err == nil {
		return &company, nil
	}
	lib.LogError("mutation/Company", err.Error())
	return nil, err
}

func (r *QueryResolver) Companies(ctx context.Context) ([]*models.Company, error) {
	var (
		companies []models.Company

		err error
	)

	if err = r.DB.Preload(clause.Associations).Find(&companies).Error; err == nil {
		var companiesfmt []*models.Company

		for i := range companies {
			companiesfmt = append(companiesfmt, &companies[i])
		}
		return companiesfmt, nil
	}
	lib.LogError("mutation/Companies", err.Error())
	return nil, err
}

// Mutation returns exec.MutationResolver implementation.
func (r *Resolver) Mutation() exec.MutationResolver { return &MutationResolver{r} }

// Query returns exec.QueryResolver implementation.
func (r *Resolver) Query() exec.QueryResolver { return &QueryResolver{r} }

type MutationResolver struct{ *Resolver }
type QueryResolver struct{ *Resolver }
