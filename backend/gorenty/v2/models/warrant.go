package models

import "errors"

// Warrant is a candidature warrant
type Warrant struct {
	Base
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Civility     string  `json:"civility"`
	Email        *Email  `json:"email"`
	Phone        *Phone  `json:"phone"`
	IDDocuments  []Asset `json:"id_documents"`
	IncomesProof []Asset `json:"incomes_proof"`
}

// WarrantItem is the db model of a candidature warrant
type WarrantItem struct {
	Base
	FirstName    string   `json:"first_name"`
	LastName     string   `json:"last_name"`
	Civility     string   `json:"civility"`
	Email        *string  `json:"email"`
	Phone        *string  `json:"phone"`
	IDDocuments  []string `json:"id_documents"`
	IncomesProof []string `json:"incomes_proof"`
}

// Validate a warrant item
func (w Warrant) Validate() error {
	if len(w.FirstName) == 0 {
		return errors.New("missing first name element")
	}
	if len(w.LastName) == 0 {
		return errors.New("missing last name element")
	}
	if len(w.Civility) == 0 {
		return errors.New("missing civility element")
	}
	return nil
}
