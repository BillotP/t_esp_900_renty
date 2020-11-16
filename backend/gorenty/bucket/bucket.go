package bucket

import (
	"log"
	"strconv"

	"github.com/minio/minio-go"
	"github.com/BillotP/gorenty"
)

// Config is the needed param to get minio client
type Config struct {
	Endpoint string
	keyID    string
	secret   string
	withSSL  string
}

// Client is the bucket client
type Client struct {
	config Config
	C      *minio.Client
}

// New return a bucket client, panic if error
func New() *Client {
	var ncli Client
	var err error
	var ssl bool
	ncli.config.Endpoint = goscrappy.MustGetSecret("minio_host")
	ncli.config.keyID = goscrappy.MustGetSecret("minio_keyid")
	ncli.config.secret = goscrappy.MustGetSecret("minio_secret")
	ncli.config.withSSL = goscrappy.MustGetSecret("minio_withssl")
	if ssl, err = strconv.ParseBool(ncli.config.withSSL); err != nil {
		panic(err)
	}
	if ncli.C, err = minio.New(ncli.config.Endpoint, ncli.config.keyID, ncli.config.secret, ssl); err != nil {
		panic(err)
	}
	return &ncli
}

// CreateUserBucket set up a bucket for a user
func (c *Client) CreateUserBucket(userID string) (bucketAddress *string, err error) {
	// var ctx = context.Background()
	err = c.C.MakeBucket(userID, "")
	bucketaddr := c.config.Endpoint + "/" + userID
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := c.C.BucketExists(userID)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", userID)
			return &bucketaddr, nil
		}
		return nil, err
	} else if goscrappy.Debug {
		log.Printf("Info(CreateUserBucket): Successfully created %s\n", bucketaddr)
	}
	return &bucketaddr, nil
}

// GetObject return an object blob from bucket
func (c *Client) GetObject(bucketName, objName string) ([]byte, error) {
	var bytearray []byte
	res, err := c.C.GetObject(bucketName, objName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	if _, err = res.Read(bytearray); err != nil {
		return nil, err
	}
	return bytearray, nil
}
