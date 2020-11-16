package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	driver "github.com/arangodb/go-driver"
	"github.com/BillotP/gorenty"
	"github.com/BillotP/gorenty/v2/models"
)

// Create a new rentofferitem from the rent offer with embed structs model
func (r *RentOfferItem) Create(item models.RentOffer) (*models.RentOfferItem, error) {
	var err error
	var nitem models.RentOfferItem
	var nlabel *models.Label
	if err = item.Validate(); err != nil {
		return nil, err
	}
	nitem.Base = models.NewBase(&item.Base.DocumentMeta.Key)
	var labelrepo = &Label{
		Repository: r.Repository,
	}
	if nlabel, err = labelrepo.Save(item.Title[0]); err != nil {
		return nil, err
	}
	nitem.Title = []string{nlabel.Key}
	var nsource *models.Source
	var sourcerepo = &Source{
		Repository: r.Repository,
	}
	if nsource, err = sourcerepo.Save(item.Source); err != nil {
		return nil, err
	}
	nitem.Source = nsource.Key
	var ndesc *models.Label
	if ndesc, err = labelrepo.Save(item.Description[0]); err != nil {
		return nil, err
	}
	nitem.Description = []string{ndesc.Key}
	var nprice *models.Price
	var pricerepo = &Price{
		Repository: r.Repository,
	}
	if nprice, err = pricerepo.Save(item.Price[0]); err != nil {
		return nil, err
	}
	nitem.Price = []string{nprice.Key}
	var nasset []models.Asset
	var assetrepo = &Asset{
		Repository: r.Repository,
	}
	if nasset, err = assetrepo.SaveMany(item.Assets); err != nil {
		return nil, err
	}
	for el := range nasset {
		nitem.Assets = append(nitem.Assets, nasset[el].Key)
	}
	var nloc *models.Location
	var locationrepo = &Location{
		Repository: r.Repository,
	}
	if nloc, err = locationrepo.Save(item.Location); err != nil {
		return nil, err
	}
	nitem.Location = nloc.Key
	var nofferor *models.OfferorItem
	var offerorrepo = &Offeror{
		Repository: r.Repository,
	}
	if nofferor, err = offerorrepo.Create(item.Offeror); err != nil {
		return nil, err
	}
	nitem.Offeror = nofferor.Key
	var nsurface *models.Surface
	var surfacerepo = &Surface{
		Repository: r.Repository,
	}
	if nsurface, err = surfacerepo.Save(item.Surface); err != nil {
		return nil, err
	}
	nitem.Surface = nsurface.Key
	nitem.CreatedAt = item.CreatedAt.Format(time.RFC3339)
	nitem.ExpiredAt = item.ExpiredAt.Format(time.RFC3339)
	nitem.Base.DocumentMeta.Key = item.Key
	if goscrappy.Debug {
		fmt.Printf("Info(Create): Will Save new Item %+v\n", nitem)
	}
	res, err := r.Save(nitem)
	if err != nil {
		return nil, err
	}
	return res, err
}

// Save a new rentoffer item in database (create or update)
func (r *RentOfferItem) Save(item models.RentOfferItem) (*models.RentOfferItem, error) {
	var err error
	ctx := context.Background()
	var col driver.Collection
	var meta driver.DocumentMeta
	var nitem models.RentOfferItem
	if r == nil || (r != nil && r.Db == nil) {
		return nil, errors.New("missing db config")
	}
	if col, err = r.Db.Collection(ctx, "rentoffers"); err != nil {
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
