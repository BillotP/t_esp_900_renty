// Package models contains all db schemes
package models

import (
	"time"

	driver "github.com/arangodb/go-driver"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base contains common columns for all tables with uuid as primary key.
type Base struct {
	*driver.DocumentMeta
	UUID      uuid.UUID  `gorm:"type:uuid;primaryKey;" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`
}

// BeforeCreateArango set createdAt field before saving in arangodb
func (base *Base) BeforeCreateArango() {
	base.CreatedAt = time.Now()
	base.UpdatedAt = time.Now()
}

// BeforeUpdateArango set updatedAt field before saving in arangodb
func (base *Base) BeforeUpdateArango() {
	base.UpdatedAt = time.Now()
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.UUID = uuid.New()
	return nil
}
