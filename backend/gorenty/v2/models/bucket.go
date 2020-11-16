package models

// Bucket is a minio bucket to store many things
type Bucket struct {
	Endpoint     string  `json:"endpoint"`
	EncAccessKey *string `json:"enc_access_key,omitempty"`
	EncSecretKey *string `json:"enc_secret_key,omitempty"`
}
