package graph

import (
	"fmt"

	"github.com/BillotP/gorenty"
)

// ByID is a wrapper struct for getby id style resolvers
type ByID struct {
	ID string `json:"id"`
}

var (
	// GatewayScheme is the openfaas gateway protocol scheme
	GatewayScheme = goscrappy.MustGetSecret("gateway-scheme")
	// GatewayHost is the openfaas gateway host
	GatewayHost = goscrappy.MustGetSecret("gateway-host")
	// GatewayPort is the openfaas gateway port
	GatewayPort = goscrappy.MustGetSecret("gateway-port")
	// GatewayBaseURL is the openfaas gateway base url
	GatewayBaseURL = fmt.Sprintf("%s://%s:%s/function", GatewayScheme, GatewayHost, GatewayPort)
)
