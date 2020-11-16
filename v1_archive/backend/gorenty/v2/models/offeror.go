package models

// Offeror is a rent offer offeror
type Offeror struct {
	Base
	Name                 string `json:"name"`
	Email                *Email `json:"email"`
	Phone                *Phone `json:"phone"`
	Type                 string `json:"type"`
	PreferredContactMode string `json:"preferred_contact_mode"`
}

// OfferorItem is the db representation of rent offer offeror
type OfferorItem struct {
	Base
	Name                 string  `json:"name"`
	Email                *string `json:"email"`
	Phone                *string `json:"phone"`
	Type                 string  `json:"type"`
	PreferredContactMode string  `json:"preferred_contact_mode"`
}
