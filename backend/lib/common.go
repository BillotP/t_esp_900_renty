package goscrappy

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	// StageDevel is the development stage
	StageDevel = "dev"
	// StageStaging is the pre-production stage
	StageStaging = "preprod"
	// StageProduction is the production stage
	StageProduction = "prod"
	// ContextAWS is the aws deployment context
	ContextAWS = "aws"
	// ContextLocal is the local deployment context
	ContextLocal = "local"
	// ContextKube is the kubernetes deployment context
	ContextKube = "kube"
	// ContextOpenfaas is the openfaas deployment context
	ContextOpenfaas = "openfaas"
)

// Stage is the stage in which this program run
var Stage = os.Getenv("STAGE")

// Context is the deployment context this programm is running in
var Context = os.Getenv("CONTEXT")

// Debug is the debug value
var Debug = os.Getenv("DEBUG") != ""

// DryRun is the dryrun value
var DryRun = os.Getenv("DRYRUN") != ""

// SecretPrefix is the prefix for ssm param name (depends on service)
var SecretPrefix = os.Getenv("SECRET_PREFIX")

// Ctx is the context
var Ctx = context.Background()

var (
	awsSession *session.Session
	ssmCli     *ssm.SSM
)

var secretPrefix = func() string {
	if Stage == StageDevel {
		return ""
	}
	return SecretPrefix + "-"
}()

// GetRedisClient return an open connection to redis key value store
func GetRedisClient() *redis.Client {
	var err error
	var pong string
	var redisURL = MustGetSecret("redis_url")
	rdb := redis.NewClient(&redis.Options{
		Addr: redisURL,
		DB:   0, // use default DB
	})
	if pong, err = rdb.Ping(Ctx).Result(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Info(GetStore): Got successfull %s from redis\n", pong)
	return rdb
}

// GetDB return an open connection to postgres database
func GetDB(migrate interface{}) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	var (
		dbHOST     = MustGetSecret("postgres_host")
		dbPORT     = MustGetSecret("postgres_port")
		dbUSER     = MustGetSecret("postgres_user")
		dbPASSWORD = MustGetSecret("postgres_password")
		dbNAME     = MustGetSecret("postgres_dbname")
		dbSSLMODE  = MustGetSecret("postgres_sslmode")
	)
	dbURL := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHOST,
		dbPORT,
		dbUSER,
		dbPASSWORD,
		dbNAME,
		dbSSLMODE,
	)
	if db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}); err != nil {
		return nil, err
	}
	if migrate != nil {
		// Migrate the schema
		if err = db.AutoMigrate(migrate); err != nil {
			return nil, err
		}
	}
	if Debug {
		fmt.Printf("Info(GetDB): Successfully connected to postgres\n")
	}
	return db, nil
}

func getOpenFaasSecret(secretName string) (secretVal string, err error) {
	var secretBytes []byte
	if strings.Contains(secretName, "_") {
		secretName = strings.ReplaceAll(secretName, "_", "-")
	}
	// read from the openfaas secrets folder
	secretBytes, err = ioutil.ReadFile("/var/openfaas/secrets/" + secretName)
	if err != nil {
		// read from the original location for backwards compatibility with openfaas <= 0.8.2
		secretBytes, err = ioutil.ReadFile("/run/secrets/" + secretName)
	}
	return string(secretBytes), err
}

func getAwsSsmSecret(secretName string) (secretVal string, err error) {
	if awsSession == nil {
		awsSession = session.Must(session.NewSession())
	}
	if ssmCli == nil {
		ssmCli = ssm.New(awsSession)
	}
	if ssmCli == nil {
		err = errors.New("failed to get ssm client")
		if Debug {
			fmt.Printf("Error(getAwsSsmSecret): %s\n", err.Error())
		}
		return "", err
	}
	secretName = secretPrefix + secretName
	if Debug {
		fmt.Printf("Info(getAwsSsmSecret): Getting secret name : %s\n", secretName)
	}
	out, err := ssmCli.GetParameter(&ssm.GetParameterInput{
		Name:           &secretName,
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		fmt.Printf("Error(getAwsSsmSecret): Can't get secret [%s]\n", secretName)
		return "", err
	}
	return *out.Parameter.Value, nil
}

// MustGetSecret from ssm param store or env var if stage dev (panic if failed)
func MustGetSecret(secretName string) (secret string) {
	var err error
	defVal := os.Getenv(secretName)
	if Context == ContextAWS {
		if defVal, err = getAwsSsmSecret(secretName); err != nil {
			log.Fatal(err)
		}
	} else if Context == ContextOpenfaas {
		if defVal, err = getOpenFaasSecret(secretName); err != nil {
			log.Fatal(err)
		}
	}
	if defVal == "" {
		log.Fatalf("missing secret %s", secretName)
	}
	return defVal
}

// GetSecret from ssm param store or env var if stage dev (handle error if any)
func GetSecret(secretName string) (secret string, err error) {
	defVal := os.Getenv(secretName)
	if Context == ContextAWS {
		if defVal, err = getAwsSsmSecret(secretName); err != nil {
			return "", err
		}
	} else if Context == ContextOpenfaas {
		if defVal, err = getOpenFaasSecret(secretName); err != nil {
			return "", err
		}
	}
	if defVal == "" {
		return "", errors.New("missing secret " + secretName)
	}
	return defVal, nil
}
