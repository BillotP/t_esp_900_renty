package service

import (
	"context"
	"errors"
	"fmt"

	driver "github.com/arangodb/go-driver"
	"github.com/BillotP/gorenty"
	"github.com/BillotP/gorenty/v2/models"
)

// WarrantItem is the dbo for warrant object
type WarrantItem struct {
	Repository
	driver.Collection
}

// Create a new warrant in db
func (w *WarrantItem) Create(item models.Warrant) (*models.WarrantItem, error) {
	var err error
	var nitem models.WarrantItem
	if err = item.Validate(); err != nil {
		return nil, err
	}
	nitem.Base = models.NewBase(nil)
	nitem.LastName = item.LastName
	nitem.FirstName = item.FirstName
	if item.Email != nil { // Save email item
		var nemail *models.Email
		emailrepo := Email{
			Repository: w.Repository,
		}
		if nemail, err = emailrepo.Save(*item.Email); err != nil {
			return nil, err
		}
		nitem.Phone = &nemail.Key
	}
	if item.Phone != nil { // Save phone item
		var nphone *models.Phone
		phonerepo := Phone{
			Repository: w.Repository,
		}
		if nphone, err = phonerepo.Save(*item.Phone); err != nil {
			return nil, err
		}
		nitem.Phone = &nphone.Key
	}
	nitem.IDDocuments = []string{}
	for _, el := range item.IDDocuments {
		nitem.IDDocuments = append(nitem.IDDocuments, el.Key)
	}
	nitem.IncomesProof = []string{}
	for _, el := range item.IncomesProof {
		nitem.IDDocuments = append(nitem.IncomesProof, el.Key)
	}
	if goscrappy.Debug {
		fmt.Printf("Info(Create): Will Save new Item %+v\n", nitem)
	}
	res, err := w.Save(nitem)
	if err != nil {
		return nil, err
	}
	return res, err
}

// Save a new candidatureitem in db
func (w *WarrantItem) Save(item models.WarrantItem) (*models.WarrantItem, error) {
	var err error
	ctx := context.Background()
	var col driver.Collection
	var meta driver.DocumentMeta
	var nitem models.WarrantItem
	if w == nil || (w != nil && w.Db == nil) {
		return nil, errors.New("missing db config")
	}
	if col, err = w.Db.Collection(ctx, "warrants"); err != nil {
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
