package models

import (
	"errors"
	"time"
)

// RentOffer is a real estate rent offer from a website
type RentOffer struct {
	Base
	Title        []Label       `json:"title"`
	Source       Source        `json:"source"`
	Description  []Label       `json:"description"`
	Price        []Price       `json:"price"`
	Assets       []Asset       `json:"assets"`
	CreatedAt    time.Time     `json:"created_at"`
	ExpiredAt    time.Time     `json:"expired_at"`
	Location     Location      `json:"location"`
	Offeror      Offeror       `json:"offeror"`
	Surface      Surface       `json:"surface"`
	Candidatures []Candidature `json:"candidatures"`
	Requirements *Requirements `json:"requirements"`
}

// Validate a rentoffer before saving
func (r RentOffer) Validate() error {
	if len(r.Title) == 0 {
		return errors.New("missing title element")
	}
	if len(r.Source.URL) == 0 {
		return errors.New("missing source element")
	}
	if len(r.Description) == 0 {
		return errors.New("missing description element")
	}
	if len(r.Price) == 0 {
		return errors.New("missing price element")
	}
	if len(r.Assets) == 0 {
		return errors.New("missing asset element")
	}
	if len(r.Location.Coordinates) == 0 {
		return errors.New("missing location element")
	}
	if len(r.Offeror.Name) == 0 {
		return errors.New("missing offeror element")
	}
	if r.Surface.Value == 0 {
		return errors.New("missing surface element")
	}
	return nil
}

// RentOfferItem is the db scheme with no embed structs
type RentOfferItem struct {
	Base
	Title        []string `json:"title"`
	Source       string   `json:"source"`
	Description  []string `json:"description"`
	Price        []string `json:"price"`
	Assets       []string `json:"assets"`
	CreatedAt    string   `json:"created_at"`
	ExpiredAt    string   `json:"expired_at"`
	Location     string   `json:"location"`
	Offeror      string   `json:"offeror"`
	Surface      string   `json:"surface"`
	Candidatures []string `json:"candidatures"`
	Requirements *string  `json:"requirements"`
}
