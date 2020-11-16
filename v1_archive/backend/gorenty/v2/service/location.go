package service

import (
	"context"
	"errors"

	driver "github.com/arangodb/go-driver"
	"github.com/BillotP/gorenty/v2/models"
)

// Location is the dbo for location model
type Location struct {
	Repository
	driver.Collection
}

// Save a new location object in DB
func (l *Location) Save(item models.Location) (*models.Location, error) {
	var err error
	ctx := context.Background()
	var col driver.Collection
	var nitem models.Location
	var meta driver.DocumentMeta
	if l == nil || (l != nil && l.Db == nil) {
		return nil, errors.New("missing db config")
	}
	item.Base = models.NewBase(nil)
	if col, err = l.Db.Collection(ctx, "locations"); err != nil {
		return nil, err
	}
	if meta, err = col.CreateDocument(ctx, item); err != nil &&
		!driver.IsConflict(err) {
		return nil, err
	} else if driver.IsConflict(err) {
		ctx = driver.WithReturnNew(ctx, item)
		if meta, err = col.UpdateDocument(ctx, item.Key, item); err != nil {
			return nil, err
		}
	}
	nitem.Base = models.Base{
		DocumentMeta: meta,
	}
	return &nitem, nil
}
