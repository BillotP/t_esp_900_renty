package function

// Response from the pem key server
type Response struct {
	Message string `json:"message"`
}

// Jwks is the pem datas
type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

// JSONWebKeys is a pem data
type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}
