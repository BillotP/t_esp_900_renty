package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api/graph/generated"
	"api/graph/model"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/minio/minio-go"
	"github.com/BillotP/renty/backend/lib/v2/models"
)

func (r *mutationResolver) Login(ctx context.Context, input *model.AuthQuery) (*model.AuthPayload, error) {
	var data io.Reader
	var results model.AuthPayload
	bodyByte, err := json.Marshal(input)
	if err != nil {
		fmt.Printf("Error(Login): %s\n", err.Error())
		return nil, err
	}
	if blob := bytes.NewBuffer(bodyByte); blob != nil {
		data = blob
	}
	res, err := http.Post(GatewayBaseURL+"/login", "application/json", data)
	if err != nil {
		fmt.Printf("Error(Login): %s\n", err.Error())
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error(Login): Got status %v\n", res.StatusCode)
		return nil, errors.New("gateway error")
	}
	dec := json.NewDecoder(res.Body)
	if err = dec.Decode(&results); err != nil {
		return nil, err
	}
	return &results, nil
}

func (r *mutationResolver) Register(ctx context.Context, input model.UserInput) (*models.User, error) {
	var data io.Reader
	var results models.User
	bodyByte, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	if blob := bytes.NewBuffer(bodyByte); blob != nil {
		data = blob
	}
	res, err := http.Post(GatewayBaseURL+"/createuser", "application/json", data)
	if err != nil {
		fmt.Printf("Error(Register): %s\n", err.Error())
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error(Register): Got status %v\n", res.StatusCode)
		return nil, errors.New("gateway error")
	}
	dec := json.NewDecoder(res.Body)
	if err = dec.Decode(&results); err != nil {
		return nil, err
	}
	return &results, nil
}

func (r *mutationResolver) Uploadfile(ctx context.Context, input model.FileInput) (*string, error) {
	bucketName := "renty-userfiles-dev"
	endpoint := "bucket.192-168-1-34.sslip.io"
	accessKeyID := "8abd3f47a1214849c7a819b1d721533f9427ad3f"
	secretAccessKey := "fbde39d03852a99df08e31c421028007433da477"

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Will check bucket\n")
	err = minioClient.MakeBucket(bucketName, "")
	if err != nil {
		fmt.Printf("Got err : %s\n", err.Error())
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			return nil, err
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
	expiry := time.Second * 24 * 60 * 60 // 1 day.
	presignedURL, err := minioClient.PresignedPutObject(bucketName, input.UserID+"/"+input.Name, expiry)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Successfully generated presigned URL", presignedURL)
	strURL := presignedURL.String()
	return &strURL, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
