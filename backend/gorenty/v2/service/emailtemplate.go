package service

import (
	"context"
	"errors"

	driver "github.com/arangodb/go-driver"
	"github.com/BillotP/gorenty/v2/models"
)

// EmailTemplateItem is the dbo for email template model
type EmailTemplateItem struct {
	Repository
	driver.Collection
}

// Create a new email template in database
func (e *EmailTemplateItem) Create(item models.EmailTemplate) (*models.EmailTemplateItem, error) {
	var err error
	var nitem models.EmailTemplateItem
	if err = item.Validate(); err != nil {
		return nil, err
	}
	nitem.Base = models.NewBase(nil)
	nitem.Label = item.Label
	nitem.Subject = item.Subject
	nitem.From = item.From
	var nasset *models.Asset
	var assetrepo = &Asset{
		Repository: e.Repository,
	}
	if nasset, err = assetrepo.Save(item.Body); err != nil {
		return nil, err
	}
	nitem.Body = nasset.Key
	res, err := e.Save(nitem)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Save a new email template item in database (create or update)
func (e *EmailTemplateItem) Save(item models.EmailTemplateItem) (*models.EmailTemplateItem, error) {
	var err error
	ctx := context.Background()
	var col driver.Collection
	var meta driver.DocumentMeta
	if e == nil || (e != nil && e.Db == nil) {
		return nil, errors.New("missing db config")
	}
	if col, err = e.Db.Collection(ctx, "email_templates"); err != nil {
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
	item.Base = models.Base{
		DocumentMeta: meta,
	}
	return &item, nil
}

// GetByLabel return an email template by its label
func (e *EmailTemplateItem) GetByLabel(label string) (*models.EmailTemplate, error) {
	var err error
	ctx := context.Background()
	var doc models.EmailTemplate
	if e == nil || (e != nil && e.Db == nil) {
		return nil, errors.New("missing db config")
	}
	q := `
	FOR tpl IN email_templates
	let a = (
		FOR asset IN assets
		FILTER asset._key == tpl.body
	)
	FILTER tpl.label == @label
	RETURN merge(tpl, {
		body: FIRST(a)
	})
	`
	cursor, err := e.Db.Query(ctx, q, map[string]interface{}{
		"label": label,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close()
	if _, err = cursor.ReadDocument(ctx, &doc); err != nil {
		return nil, err
	}
	return &doc, nil
}
