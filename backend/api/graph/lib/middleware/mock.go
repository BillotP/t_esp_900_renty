package middleware

import (
	"context"
	"database/sql"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/exec"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/resolvers"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
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

func InitMockDB(userrole models.Role) {
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
			HasRole: func(ctx context.Context, obj interface{}, next graphql.Resolver, roles []*models.Role) (res interface{}, err error) {
				return next(ctx)
			},
		}}
	myHandler := handler.NewDefaultServer(exec.NewExecutableSchema(c))
	myHandler.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		graphql.GetOperationContext(ctx).DisableIntrospection = true
		namekontext := lib.ContextKey("username")
		rolekontext := lib.ContextKey("userrole")
		ctx = context.WithValue(ctx, namekontext, "admin")
		ctx = context.WithValue(ctx, rolekontext, userrole)
		return next(ctx)
	})
	Server = client.New(myHandler)
}
