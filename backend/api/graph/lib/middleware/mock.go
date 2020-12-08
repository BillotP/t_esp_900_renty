package middleware

import (
	"database/sql"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/exec"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/resolvers"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib/directive"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	Mock   sqlmock.Sqlmock
	Server *client.Client
)

func InitMockDB() {
	var (
		db    *gorm.DB
		sqlDb *sql.DB
		err   error
	)

	sqlDb, Mock, err = sqlmock.New()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color,

		},
	)

	if db, err = gorm.Open(postgres.Dialector{Config: &postgres.Config{Conn: sqlDb}}, &gorm.Config{Logger: newLogger}); err != nil {
		panic(err.Error())
	}
	c := exec.Config{
		Resolvers: &resolvers.Resolver{DB: db},
		Directives: exec.DirectiveRoot{
			HasRole: directive.HasRole,
		}}
	Server = client.New(handler.NewDefaultServer(exec.NewExecutableSchema(c)))
}
