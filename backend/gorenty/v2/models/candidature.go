package models

import (
	"errors"
	"time"
)

// Candidature is a candidature on a rent offer
type Candidature struct {
	Base
	Status       string     `json:"status"`
	Candidate    User       `json:"candidate"`
	IDDocuments  []Asset    `json:"id_documents"`
	IncomesProof []Asset    `json:"incomes_proof"`
	Warrants     []Warrant  `json:"warrants"`
	Appointment  *time.Time `json:"appointment"`
}

// CandidatureItem  is the db model for a candidature on a rent offer
type CandidatureItem struct {
	Base
	Status       string   `json:"status"`
	Candidate    string   `json:"candidate"`
	IDDocuments  []string `json:"id_documents"`
	IncomesProof []string `json:"incomes_proof"`
	Warrants     []string `json:"warrants"`
	Appointment  *string  `json:"appointment"`
}

// Validate a candidature item
func (c Candidature) Validate() error {
	if len(c.Candidate.Key) == 0 {
		return errors.New("missing candidate element")
	}
	if len(c.Status) == 0 {
		return errors.New("missing status element")
	}
	return nil
}
