package models

// Geometry is a coordinates element
//	coordinates[0] is longitude , coordinates[1] is latitude
type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// Properties is used to store metadatas in GeoJson object
type Properties struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

// GeoJSON is an RFC7946 GeoJSON object
type GeoJSON struct {
	Type       string `json:"type"`
	Geometry   `json:"geometry"`
	Properties `json:"properties"`
}

// Location is a geo location element
type Location struct {
	Base
	GeoJSON
}
