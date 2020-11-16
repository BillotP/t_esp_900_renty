package models

// Source is an external website
type Source struct {
	Base
	URL string `json:"url"`
}

// Asset is an asset like a document or a picture
type Asset struct {
	Base
	URL     string `json:"url"`
	Type    string `json:"type"`
	Storage string `json:"storage"`
}

// Label is a text element with locale and fulltext search index on it
type Label struct {
	Base
	Value  string `json:"value"`
	Locale string `json:"locale"`
}

// Price is a price element
type Price struct {
	Base
	Value     float64 `json:"value"`
	Condition string  `json:"condition"`
	Currency  string  `json:"currency"`
}

// Surface is a surface element
type Surface struct {
	Base
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
	Rooms int64   `json:"rooms"`
}
