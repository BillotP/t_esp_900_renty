package resolvers_test

import (
	"fmt"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/exec"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/resolvers"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/directive"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	Db *gorm.DB
	Server *client.Client
)

func init() {
	var (
		username string
		password string
		dbName   string
		dbHost   string

		err error
	)

	username = "postgres"
	password = "AFEC62CDD545955405CC3BCBD00CE85D"
	dbName = "rentydb"
	dbHost = "127.0.0.1"

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, username, dbName, password, "5432") //Build connection string
	fmt.Println(dbUri)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color,

		},
	)

	Db, err = gorm.Open(postgres.Open(dbUri), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Print(err)
	}

	if err = Db.AutoMigrate(&models.Admin{}, &models.Anomaly{}, models.Asset{}, &models.Company{}, &models.EstateAgent{}, &models.Property{}, &models.Tenant{}, &models.User{}); err != nil {
		os.Exit(84)
	}
	c := exec.Config{
		Resolvers: &resolvers.Resolver{DB: Db},
		Directives: exec.DirectiveRoot{
			HasRole: directive.HasRole,
		}}
	Server = client.New(handler.NewDefaultServer(exec.NewExecutableSchema(c)))
}

