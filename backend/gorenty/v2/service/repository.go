package service

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"strconv"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/BillotP/gorenty"
	"github.com/BillotP/gorenty/v2/models"
)

// DbConf is the database config object
type DbConf struct {
	Host      string `json:"db_host"`
	Port      string `json:"db_port"`
	Scheme    string `json:"db_scheme"`
	User      string `json:"db_user"`
	Password  string `json:"db_password"`
	TLSVerify string `json:"db_tlsverify"`
}

// URL return the full db url
func (d *DbConf) URL() string {
	return fmt.Sprintf(
		"%s://%s:%s",
		d.Scheme,
		d.Host,
		d.Port,
	)
}

// Repository is the database repository service
type Repository struct {
	Name string
	C    driver.Client
	Db   driver.Database
	conf *DbConf
}

func loadconfig() (*DbConf, error) {
	var err error
	var cnf DbConf
	if cnf.Host, err = goscrappy.GetSecret("arango_host"); err != nil {
		return nil, err
	}
	if cnf.Port, err = goscrappy.GetSecret("arango_port"); err != nil {
		return nil, err
	}
	if cnf.Scheme, err = goscrappy.GetSecret("arango_scheme"); err != nil {
		return nil, err
	}
	if cnf.User, err = goscrappy.GetSecret("arango_user"); err != nil {
		return nil, err
	}
	if cnf.Password, err = goscrappy.GetSecret("arango_password"); err != nil {
		return nil, err
	}
	if cnf.TLSVerify, err = goscrappy.GetSecret("arango_tlsverify"); err != nil {
		return nil, err
	}
	return &cnf, nil
}

// New setup the database repository
func New(dbname string) (*Repository, error) {
	var err error
	var nrep = Repository{
		Name: dbname,
	}
	var conn driver.Connection
	ctx := context.Background()
	if nrep.conf, err = loadconfig(); err != nil {
		return nil, err
	}
	var withTLS = func() bool {
		val, err := strconv.ParseBool(nrep.conf.TLSVerify)
		if err != nil || val == false {
			return false
		}
		return true
	}()
	if conn, err = http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{nrep.conf.URL()},
		TLSConfig: &tls.Config{
			InsecureSkipVerify: !withTLS,
		},
	}); err != nil {
		return nil, err
	}
	if nrep.C, err = driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(nrep.conf.User, nrep.conf.Password),
	}); err != nil {
		return nil, err
	}
	if nrep.Db, err = nrep.C.Database(ctx, nrep.Name); err != nil {
		return nil, err
	}
	return &nrep, err
}

// AllDatabases is a list of all renty database name
var AllDatabases = []string{
	"renty-dev",
	"assets-dev",
}

// AllCollections is a list of all renty collection names
var AllCollections = map[string][]string{
	"renty-dev": []string{
		"sources",
		"assets",
		"labels",
		"locations",
		"prices",
		"offerors",
		"rentoffers",
		"users",
		"phones",
		"emails",
		"companies",
		"surfaces",
		"candidatures",
		"requirements",
		"warrants",
	},
	"assets-dev": []string{
		"email_templates",
		"sms_templates",
		"document_templates",
	},
}

// UserItem is the dbo for user model
type UserItem struct {
	Repository
	driver.Collection
}

// RentOfferItem is the dbo for rentoffer model
type RentOfferItem struct {
	Repository
	driver.Collection
}

// Offeror is the dbo for offeror model
type Offeror struct {
	Repository
	driver.Collection
}

// Price is the dbo for price model
type Price struct {
	Repository
	driver.Collection
}

// Surface is the dbo for surface model
type Surface struct {
	Repository
	driver.Collection
}

// Label is the dbo for label model
type Label struct {
	Repository
	driver.Collection
}

// Asset is the dbo for asset model
type Asset struct {
	Repository
	driver.Collection
}

// Source is the dbo for source model
type Source struct {
	Repository
	driver.Collection
}

// Save a new offeror object in db (create or update)
func (s *Surface) Save(item models.Surface) (*models.Surface, error) {
	var err error
	ctx := context.Background()
	var col driver.Collection
	var nitem models.Surface
	var meta driver.DocumentMeta
	if s == nil || (s != nil && s.Db == nil) {
		return nil, errors.New("missing db config")
	}
	item.Base = models.NewBase(nil)
	if col, err = s.Db.Collection(ctx, "surfaces"); err != nil {
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

// Save a new price object in DB
func (p *Price) Save(item models.Price) (*models.Price, error) {
	var err error
	ctx := context.Background()
	var col driver.Collection
	var nitem models.Price
	var meta driver.DocumentMeta
	if p == nil || (p != nil && p.Db == nil) {
		return nil, errors.New("missing db config")
	}
	item.Base = models.NewBase(nil)
	if col, err = p.Db.Collection(ctx, "prices"); err != nil {
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

// Save a new label object in DB
func (l *Label) Save(item models.Label) (*models.Label, error) {
	var err error
	ctx := context.Background()
	var col driver.Collection
	var nitem models.Label
	var meta driver.DocumentMeta
	if l == nil || (l != nil && l.Db == nil) {
		return nil, errors.New("missing db config")
	}
	item.Base = models.NewBase(nil)
	if col, err = l.Db.Collection(ctx, "labels"); err != nil {
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

// Save a new asset object in DB
func (a *Asset) Save(item models.Asset) (*models.Asset, error) {
	var err error
	ctx := context.Background()
	var col driver.Collection
	var nitem models.Asset
	var meta driver.DocumentMeta
	if a == nil || (a != nil && a.Db == nil) {
		return nil, errors.New("missing db config")
	}
	item.Base = models.NewBase(nil)
	if col, err = a.Db.Collection(ctx, "assets"); err != nil {
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

// SaveMany new asset object in DB
func (a *Asset) SaveMany(item []models.Asset) ([]models.Asset, error) {
	var err error
	ctx := context.Background()
	var col driver.Collection
	var meta driver.DocumentMetaSlice
	if a == nil || (a != nil && a.Db == nil) {
		return nil, errors.New("missing db config")
	}
	if col, err = a.Db.Collection(ctx, "assets"); err != nil {
		return nil, err
	}
	for el := range item {
		if item[el].Base.Key == "" {
			item[el].Base = models.NewBase(nil)
		}
	}
	if meta, _, err = col.CreateDocuments(ctx, item); err != nil {
		return nil, err
	}
	for el := range meta {
		item[el].Base = models.Base{
			DocumentMeta: meta[el],
		}
	}
	return item, nil
}

// Save a new source object in DB
func (s *Source) Save(item models.Source) (*models.Source, error) {
	var err error
	ctx := context.Background()
	var col driver.Collection
	var nitem models.Source
	var meta driver.DocumentMeta
	if s == nil || (s != nil && s.Db == nil) {
		return nil, errors.New("missing db config")
	}
	item.Base = models.NewBase(nil)
	if col, err = s.Db.Collection(ctx, "sources"); err != nil {
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
