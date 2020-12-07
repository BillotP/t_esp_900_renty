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

func (r *mutationResolver) SignupAsAdmin(ctx context.Context, input models.AdminInput) (*models.Credential, error) {
	panic("not implemented")
}

func (r *mutationResolver) SignupAsCompany(ctx context.Context, input models.CompanyInput) (*models.Credential, error) {
	var (
		token   string
		company *models.Company
		pwdHash []byte

		err error
	)

	verified := false

	company = &models.Company{
		Name: input.Name,
		User: &models.User{
			Username: input.User.Username,
			Password: "",
			Role:     models.RoleCompany,
		},
		Description: &input.Description,
		Tel:         input.Tel,
		Verified:    &verified,
	}
	if err = r.DB.Where("name = ?", company.Name).First(&company).Error; err == nil {
		return nil, fmt.Errorf("company seems already register")
	}
	if pwdHash, err = bcrypt.GenerateFromPassword([]byte(input.User.Password), getPseudoRandomCost()); err != nil {
		lib.LogError("mutation/Register", err.Error())
		return nil, err
	}

	company.User.Password = string(pwdHash)
	if err = r.DB.Create(&company).Error; err != nil {
		lib.LogError("mutation/Register/Company", err.Error())
		return nil, err
	}
	if token, err = createToken(company.User.Username, company.User.Role); err != nil {
		return nil, err
	}
	return &models.Credential{
		User:  company.User,
		Token: &token,
	}, nil
}

func (r *mutationResolver) CreateEstateAgentUser(ctx context.Context, input *models.EstateAgentInput) (*models.EstateAgent, error) {
	var (
		usernameCtx = lib.ContextKey("username")

		company     *models.Company
		estateAgent *models.EstateAgent
		pwdHash     []byte

		err error
	)

	companyUsername := ctx.Value(usernameCtx).(string)

	company = &models.Company{
		User: &models.User{
			Username: companyUsername,
		},
	}
	if err = r.DB.Where(&company).First(&company).Error; err != nil {
		return nil, err
	}

	estateAgent = &models.EstateAgent{
		Company: company,
		User: &models.User{
			Username: input.User.Username,
			Password: "",
			Role:     models.RoleEstateAgent,
		},
	}
	if err = r.DB.Joins("User").Where("username = ?", input.User.Username).First(&estateAgent).Error; err == nil {
		return nil, fmt.Errorf("estate agent seems already register")
	}
	if pwdHash, err = bcrypt.GenerateFromPassword([]byte(input.User.Password), getPseudoRandomCost()); err != nil {
		lib.LogError("mutation/Register", err.Error())
		return nil, err
	}

	estateAgent.User.Password = string(pwdHash)
	if err = r.DB.Create(&estateAgent).Error; err != nil {
		lib.LogError("mutation/Register/EstateAgent", err.Error())
		return nil, err
	}
	return estateAgent, nil
}

func (r *mutationResolver) CreateTenantUser(ctx context.Context, input *models.TenantInput) (*models.Tenant, error) {
	var (
		usernameCtx = lib.ContextKey("username")

		estateAgent *models.EstateAgent
		tenant      *models.Tenant
		pwdHash     []byte

		err error
	)

	estateAgentUsername := ctx.Value(usernameCtx).(string)

	estateAgent = &models.EstateAgent{
		User: &models.User{
			Username: estateAgentUsername,
		},
	}
	if err = r.DB.Where(&estateAgent).First(&estateAgent).Error; err != nil {
		return nil, err
	}

	tenant = &models.Tenant{
		EstateAgent: estateAgent,
		User: &models.User{
			Username: input.User.Username,
			Password: "",
			Role:     models.RoleEstateAgent,
		},
	}
	if err = r.DB.Joins("User").Where("username = ?", input.User.Username).First(&tenant).Error; err == nil {
		return nil, fmt.Errorf("tenant seems already register")
	}
	if pwdHash, err = bcrypt.GenerateFromPassword([]byte(input.User.Password), getPseudoRandomCost()); err != nil {
		lib.LogError("mutation/Register", err.Error())
		return nil, err
	}

	tenant.User.Password = string(pwdHash)
	if err = r.DB.Create(&tenant).Error; err != nil {
		lib.LogError("mutation/Register/Tenant", err.Error())
		return nil, err
	}
	return tenant, nil
}

func (r *mutationResolver) AcceptCompany(ctx context.Context) (*models.Company, error) {
	panic("Wilfried : not implemented")
}

func (r *mutationResolver) LoginAsCompany(ctx context.Context, input *models.UserInput) (*models.Credential, error) {
	var (
		company *models.Company
		token   = ""

		err error
	)

	company = &models.Company{}
	if err = r.DB.Joins("User").Where("username = ?", input.Username).First(&company).Error; err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(company.User.Password), []byte(input.Password)); err != nil {
		lib.LogError("resolvers/LoginAsCompany", err.Error())
		return nil, fmt.Errorf("bad password provided")
	}
	if token, err = createToken(company.User.Username, company.User.Role); err != nil {
		return nil, err
	}
	return &models.Credential{
		User:  company.User,
		Token: &token,
	}, nil
}

func (r *mutationResolver) LoginAsEstateAgent(ctx context.Context, input *models.UserInput) (*models.Credential, error) {
	var (
		estateAgent *models.EstateAgent
		token   = ""

		err error
	)

	estateAgent = &models.EstateAgent{}
	if err = r.DB.Joins("User").Where("username = ?", input.Username).First(&estateAgent).Error; err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(estateAgent.User.Password), []byte(input.Password)); err != nil {
		lib.LogError("resolvers/LoginAsCompany", err.Error())
		return nil, fmt.Errorf("bad password provided")
	}
	if token, err = createToken(estateAgent.User.Username, estateAgent.User.Role); err != nil {
		return nil, err
	}
	return &models.Credential{
		User:  estateAgent.User,
		Token: &token,
	}, nil
}

func (r *mutationResolver) LoginAsTenant(ctx context.Context, input *models.UserInput) (*models.Credential, error) {
	var (
		tenant *models.Tenant
		token   = ""

		err error
	)

	tenant = &models.Tenant{}
	if err = r.DB.Joins("User").Where("username = ?", input.Username).First(&tenant).Error; err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(tenant.User.Password), []byte(input.Password)); err != nil {
		lib.LogError("resolvers/LoginAsCompany", err.Error())
		return nil, fmt.Errorf("bad password provided")
	}
	if token, err = createToken(tenant.User.Username, tenant.User.Role); err != nil {
		return nil, err
	}
	return &models.Credential{
		User:  tenant.User,
		Token: &token,
	}, nil
}

func (r *mutationResolver) UpdateTenantProfile(ctx context.Context, input *models.TenantUpdateInput) (*models.Tenant, error) {
	panic("Wilfried : not implemented")
}

func (r *mutationResolver) CreateProperty(ctx context.Context, input *models.PropertyInput) (*models.Property, error) {
	var (
		property *models.Property
		err error
	)
	property = &models.Property{
		Area:       input.Area,
		Address: 	input.Address,
		CodeNumber: input.CodeNumber,
		Type:		input.Type,
	}
	if err = r.DB.Where("address = ?", property.Address).First(&property).Error; err == nil {
		return nil, fmt.Errorf("There is already a property at this address")
	}
	if err = r.DB.Create(&property).Error; err != nil {
		lib.LogError("mutation/Register/Property", err.Error())
		return nil, err
	}
	return property, nil
}

func (r *mutationResolver) CreateAnomaly(ctx context.Context, input *models.AnomalyInput) (*models.Anomaly, error) {
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
		return nil, fmt.Errorf("Anomaly already created")
	}
	if err = r.DB.Create(&anomaly).Error; err != nil {
		lib.LogError("mutation/Register/Anomaly", err.Error())
		return nil, err
	}
	return anomaly, nil
}

func (r *mutationResolver) UpdateAnomaly(ctx context.Context, input *models.AnomalyUpdateInput) (*models.Anomaly, error) {
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

func (r *queryResolver) Anomaly(ctx context.Context, id string) (*models.Anomaly, error) {
	var (
		anomaly *models.Anomaly
		err error
	)

	if err = r.DB.Where("ID = ?", id).First(&anomaly).Error; err == nil {
		return nil, fmt.Errorf("Anomaly not found")
	}
	return anomaly, nil
}

func (r *queryResolver) Anomalies(ctx context.Context) ([]*models.Anomaly, error) {
	var (
		anomalies []*models.Anomaly
		err error
	)

	if err = r.DB.Find(&anomalies).Error; err == nil {
		return nil, fmt.Errorf("No anomalies found")
	}
	return anomalies, nil
}

func (r *queryResolver) Tenant(ctx context.Context, id string) (*models.Tenant, error) {
	panic("Wilfried : not implemented")
}

func (r *queryResolver) Tenants(ctx context.Context) ([]*models.Tenant, error) {
	panic("Wilfried : not implemented")
}

func (r *queryResolver) EstateAgent(ctx context.Context, id string) (*models.EstateAgent, error) {
	panic("Wilfried : not implemented")
}

func (r *queryResolver) EstateAgents(ctx context.Context) ([]*models.EstateAgent, error) {
	panic("Wilfried : not implemented")
}

func (r *queryResolver) Company(ctx context.Context, id string) (*models.Company, error) {
	panic("Wilfried : not implemented")
}

func (r *queryResolver) Companies(ctx context.Context) ([]*models.Company, error) {
	panic("Wilfried : not implemented")
}

// Mutation returns exec.MutationResolver implementation.
func (r *Resolver) Mutation() exec.MutationResolver { return &mutationResolver{r} }

// Query returns exec.QueryResolver implementation.
func (r *Resolver) Query() exec.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
