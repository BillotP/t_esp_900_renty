package models

import "errors"

// Email is a renty user email
type Email struct {
	Base
	Valid bool   `json:"valid"`
	Value string `json:"value"`
}

// Validate an email before saving it
func (e Email) Validate() error {
	if len(e.Value) == 0 {
		return errors.New("missing value item")
	}
	return nil
}

// Phone is a phone number with location code
type Phone struct {
	Base
	CountryCode string `json:"country_code"`
	Value       string `json:"value"`
	Valid       bool   `json:"valid"`
	String      string `json:"-"`
}

// Validate a phone number before saving it
func (p Phone) Validate() error {
	if len(p.Value) == 0 {
		return errors.New("missing value item")
	}
	return nil
}
