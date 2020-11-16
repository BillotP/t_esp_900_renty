package models

import (
	"time"

	driver "github.com/arangodb/go-driver"
	"github.com/google/uuid"
)

// Base is a base document on arangodb collection
type Base struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	driver.DocumentMeta
}

// NewBase is the default base document setup
func NewBase(key *string) Base {
	nKey := uuid.New().String()
	if key != nil && *key != "" {
		nKey = *key
	}
	return Base{
		DocumentMeta: driver.DocumentMeta{
			Key: nKey,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
