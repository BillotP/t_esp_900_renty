package service

import (
	"context"
	"errors"
	"fmt"

	driver "github.com/arangodb/go-driver"
	"github.com/BillotP/gorenty"
	"github.com/BillotP/gorenty/v2/models"
)

// CandidatureItem is the dbo for candidature model
type CandidatureItem struct {
	Repository
	driver.Collection
}

// Create a new candidature in db
func (c *CandidatureItem) Create(item models.Candidature) (*models.CandidatureItem, error) {
	var err error
	var nitem models.CandidatureItem
	if err = item.Validate(); err != nil {
		return nil, err
	}
	nitem.Base = models.NewBase(nil)
	nitem.Candidate = item.Candidate.Key
	nitem.IDDocuments = []string{}
	for _, el := range item.IDDocuments {
		nitem.IDDocuments = append(nitem.IDDocuments, el.Key)
	}
	nitem.IncomesProof = []string{}
	for _, el := range item.IncomesProof {
		nitem.IDDocuments = append(nitem.IncomesProof, el.Key)
	}
	nitem.Warrants = []string{}
	for _, el := range item.Warrants {
		nitem.Warrants = append(nitem.Warrants, el.Key)
	}
	if goscrappy.Debug {
		fmt.Printf("Info(Create): Will Save new Item %+v\n", nitem)
	}
	res, err := c.Save(nitem)
	if err != nil {
		return nil, err
	}
	return res, err
}

// Save a new candidatureitem in db
func (c *CandidatureItem) Save(item models.CandidatureItem) (*models.CandidatureItem, error) {
	var err error
	ctx := context.Background()
	var col driver.Collection
	var meta driver.DocumentMeta
	var nitem models.CandidatureItem
	if c == nil || (c != nil && c.Db == nil) {
		return nil, errors.New("missing db config")
	}
	if col, err = c.Db.Collection(ctx, "candidatures"); err != nil {
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
