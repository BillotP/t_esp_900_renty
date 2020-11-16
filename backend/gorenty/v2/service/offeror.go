package service

import (
	"context"
	"errors"

	driver "github.com/arangodb/go-driver"
	"github.com/BillotP/gorenty/v2/models"
)

// Create save a new offeror in db including nested struct
func (o *Offeror) Create(item models.Offeror) (*models.OfferorItem, error) {
	var err error
	var nitem models.OfferorItem
	var nemail *models.Email
	var nphone *models.Phone
	// if err = item.Validate(); err != nil {
	// 	return nil, err
	// }
	nitem.Base = models.NewBase(nil)
	var emailrepo = &Email{
		Repository: o.Repository,
	}
	var phonerepo = &Phone{
		Repository: o.Repository,
	}
	if item.Email != nil {
		if err = item.Email.Validate(); err == nil {
			if nemail, err = emailrepo.Save(*item.Email); err != nil {
				return nil, err
			}
			nitem.Email = &nemail.DocumentMeta.Key
		}
	}
	if item.Phone != nil {
		if err = item.Phone.Validate(); err == nil {
			if nphone, err = phonerepo.Save(*item.Phone); err != nil {
				return nil, err
			}
			nitem.Phone = &nphone.DocumentMeta.Key
		}
	}
	nitem.Name = item.Name
	nitem.Type = item.Type
	nitem.PreferredContactMode = item.PreferredContactMode
	res, err := o.Save(nitem)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Save a new offeror object in db (create or update)
func (o *Offeror) Save(item models.OfferorItem) (*models.OfferorItem, error) {
	var err error
	ctx := context.Background()
	var col driver.Collection
	var nitem models.OfferorItem
	var meta driver.DocumentMeta
	if o == nil || (o != nil && o.Db == nil) {
		return nil, errors.New("missing db config")
	}
	item.Base = models.NewBase(nil)
	if col, err = o.Db.Collection(ctx, "offerors"); err != nil {
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
