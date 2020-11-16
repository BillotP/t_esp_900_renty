package service

import (
	"context"
	"errors"

	driver "github.com/arangodb/go-driver"
	"github.com/BillotP/gorenty/v2/models"
)

// Phone is the dbo for phone model
type Phone struct {
	Repository
	driver.Collection
}

// Email is the dbo for email model
type Email struct {
	Repository
	driver.Collection
}

// Save a new email object in db (create or update)
func (e *Email) Save(item models.Email) (*models.Email, error) {
	var err error
	ctx := context.Background()
	var nitem models.Email
	var meta driver.DocumentMeta
	if e == nil || (e != nil && e.Db == nil) {
		return nil, errors.New("missing db config")
	}
	item.Base = models.NewBase(nil)
	if e.Collection, err = e.Db.Collection(ctx, "emails"); err != nil {
		return nil, err
	}
	if meta, err = e.Collection.CreateDocument(ctx, item); err != nil &&
		!driver.IsConflict(err) {
		return nil, err
	} else if driver.IsConflict(err) {
		ctx = driver.WithReturnNew(ctx, item)
		if meta, err = e.Collection.UpdateDocument(ctx, item.Key, item); err != nil {
			return nil, err
		}
	}
	nitem.Base = models.Base{
		DocumentMeta: meta,
	}
	return &nitem, nil
}

// Save a new phone number object in db (create or update)
func (p *Phone) Save(item models.Phone) (*models.Phone, error) {
	var err error
	ctx := context.Background()
	var nitem models.Phone
	var meta driver.DocumentMeta
	if p == nil || (p != nil && p.Db == nil) {
		return nil, errors.New("missing db config")
	}
	item.Base = models.NewBase(nil)
	if p.Collection, err = p.Db.Collection(ctx, "phones"); err != nil {
		return nil, err
	}
	if meta, err = p.Collection.CreateDocument(ctx, item); err != nil &&
		!driver.IsConflict(err) {
		return nil, err
	} else if driver.IsConflict(err) {
		ctx = driver.WithReturnNew(ctx, item)
		if meta, err = p.Collection.UpdateDocument(ctx, item.Key, item); err != nil {
			return nil, err
		}
	}
	nitem.Base = models.Base{
		DocumentMeta: meta,
	}
	return &nitem, nil
}
