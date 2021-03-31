package middleware

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/generated/models"
	"github.com/BillotP/t_esp_900_renty/v2/backend/api/graph/lib"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

// InitDB setup the global Db connection pool
func InitDB() {
	var (
		stage    string
		username string
		password string
		dbName   string
		dbHost   string

		err error
	)

	if err = godotenv.Load(".dev.env"); err != nil {
		lib.LogError("InitDb", err.Error())
	}
	stage = lib.GetDefVal("STAGE", "dev")
	username = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbName = lib.GetDefVal("DB_NAME", "rentydb")
	dbHost = lib.GetDefVal("POSTGRES_HOST", "127.0.0.1")
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, username, dbName, password, "5432") //Build connection string
	if stage == "dev" {
		lib.LogInfo("InitDB", fmt.Sprintf("will connect to db with URI : %s", dbURI))
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color,

		},
	)

	Db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Print(err)
	}

	if err = Db.AutoMigrate(&models.Admin{}, &models.Anomaly{}, models.Asset{}, &models.Company{}, &models.EstateAgent{}, &models.Property{}, &models.Tenant{}, &models.User{}); err != nil {
		os.Exit(84)
	}
}
