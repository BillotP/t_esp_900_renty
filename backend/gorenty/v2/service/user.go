package service

import (
	"context"
	"errors"
	"fmt"

	driver "github.com/arangodb/go-driver"
	"github.com/BillotP/gorenty"
	"github.com/BillotP/gorenty/v2/models"
	"golang.org/x/crypto/bcrypt"
)

// Create a new useritem from the user embed struct model
func (u *UserItem) Create(item models.User) (*models.UserItem, error) {
	var err error
	var nitem models.UserItem
	if err = item.Validate(); err != nil {
		return nil, err
	}
	nitem.Base = models.NewBase(nil)
	hash, err := bcrypt.GenerateFromPassword([]byte(*item.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	nitem.PasswordHash = string(hash)
	nitem.Pseudo = item.Pseudo
	nitem.Gender = item.Gender
	nitem.FirstName = item.FirstName
	nitem.LastName = item.LastName
	if item.Company != nil {
		// Save company item
	}
	if item.Email != nil { // Save email item
		var nemail *models.Email
		emailrepo := Email{
			Repository: u.Repository,
		}
		if nemail, err = emailrepo.Save(*item.Email); err != nil {
			return nil, err
		}
		nitem.Phone = &nemail.Key
	}
	if item.Phone != nil {
		var nphone *models.Phone
		phonerepo := Phone{
			Repository: u.Repository,
		}
		if nphone, err = phonerepo.Save(*item.Phone); err != nil {
			return nil, err
		}
		nitem.Phone = &nphone.Key
	}
	if goscrappy.Debug {
		fmt.Printf("Info(Create): Will Save new Item %+v\n", nitem)
	}
	res, err := u.Save(nitem)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Save a new rentoffer item in database (create or update)
func (u *UserItem) Save(item models.UserItem) (*models.UserItem, error) {
	var err error
	ctx := context.Background()
	var col driver.Collection
	var meta driver.DocumentMeta
	if u == nil || (u != nil && u.Db == nil) {
		return nil, errors.New("missing db config")
	}
	if col, err = u.Db.Collection(ctx, "users"); err != nil {
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
