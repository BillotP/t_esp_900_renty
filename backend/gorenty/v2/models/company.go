package models

// Company is registered company
type Company struct {
	Base
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	Email        *Email    `json:"email"`
	Phone        *Phone    `json:"phone"`
	Country      string    `json:"country"`
	StreetName   *string   `json:"street_name"`
	StreetNumber *string   `json:"street_number"`
	City         *string   `json:"city"`
	PostalCode   *string   `json:"postal_code"`
	Location     *Location `json:"location"`
	SIRENNumber  *string   `json:"siren_number"`
	VATNumber    *string   `json:"vat_number"`
	DUNSNumber   *string   `json:"duns_number"`
}

// CompanyItem is registered company in database
type CompanyItem struct {
	Base
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	Email        *string `json:"email"`
	Phone        *string `json:"phone"`
	Country      string  `json:"country"`
	StreetName   *string `json:"street_name"`
	StreetNumber *string `json:"street_number"`
	City         *string `json:"city"`
	PostalCode   *string `json:"postal_code"`
	Location     *string `json:"location"`
	SIRENNumber  *string `json:"siren_number"`
	VATNumber    *string `json:"vat_number"`
	DUNSNumber   *string `json:"duns_number"`
}
