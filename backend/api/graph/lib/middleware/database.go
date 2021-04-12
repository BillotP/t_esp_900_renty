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

// ExitCodeECOMM is linux exit code representing a communication error on send
const ExitCodeECOMM = 70

// ExitCodeEBADE is linux exit code representing an invalid exchange
const ExitCodeEBADE = 52

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
	finalDBURI := lib.GetDefVal("DATABASE_URL", dbURI)
	logrConf := logger.Config{
		SlowThreshold: time.Second, // Slow SQL threshold
		LogLevel:      logger.Warn, // Log level
		Colorful:      false,       // Disable color,
	}
	if stage == "dev" {
		lib.LogInfo("InitDB", fmt.Sprintf("will connect to db with URI : %s", finalDBURI))
		logrConf.LogLevel = logger.Info
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logrConf,
	)
	if Db, err = gorm.Open(postgres.Open(finalDBURI), &gorm.Config{
		Logger: newLogger,
	}); err != nil {
		lib.LogError("InitDB", "Failed to connect to database, exiting")
		os.Exit(ExitCodeECOMM)
	}

	if err = Db.AutoMigrate(
		&models.Admin{},
		&models.Anomaly{},
		&models.Asset{},
		&models.Badge{},
		&models.Company{},
		&models.EstateAgent{},
		&models.Property{},
		&models.Tenant{},
		&models.User{}); err != nil {
		lib.LogError("InitDB/AutoMigrate", "Failed to migrate db schemes, exiting")
		os.Exit(ExitCodeEBADE)
	}
}
