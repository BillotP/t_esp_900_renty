package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/BillotP/gorenty"
)

var (
	// MaxPayloadSize is 1mb
	MaxPayloadSize = int(1048576)
	// ErrTooLarge is returned when req body size is > MaxPayloadSize
	ErrTooLarge = errors.New("Payload too large")
)

// ItemsBody is a response format when asking for a slice type datas
type ItemsBody struct {
	Data  interface{} `json:"data"`
	Items int         `json:"items"`
}

// ErrorBody is an error from api
type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

// LAMBDAISAUTH check authorization header field in aws lambda context
func LAMBDAISAUTH(req events.APIGatewayProxyRequest, apiKey string) bool {
	authheader := req.Headers["Authorization"]
	return authheader == apiKey
}

// LAMBDAResponse send a json response through amazon api gateway
func LAMBDAResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	var err error
	var stringBody []byte
	resp := events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Methods": "GET,POST,DELETE,PATCH,OPTIONS",
			"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
			"Access-Control-Allow-Origin":  "*",
		},
	}
	if stringBody, err = json.Marshal(body); err != nil {
		fmt.Printf("Error(APIResponse): Failed to marshal resp : %s\n", err.Error())
		return nil, err
	}
	resp.Body = string(stringBody)
	return &resp, nil
}

// OPENFAASJsonResponse send a json response
func OPENFAASJsonResponse(w http.ResponseWriter, status int, body interface{}) {
	var err error
	var bodyByte []byte
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,DELETE,PATCH,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Date,Authorization,X-Xss-Protection,Content-Length")
	if bodyByte, err = json.Marshal(body); err != nil {
		OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bodyByte)
	return
}

// OPENFAASErrorResponse write an error response in openfaas context
func OPENFAASErrorResponse(w http.ResponseWriter, status int, err error) {
	log.SetPrefix("ERROR ")
	log.Printf("%s\n", err.Error())
	w.WriteHeader(status)
	w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	return
}

// LAMBDAGetBody check and deserialze the req payload in handler
func LAMBDAGetBody(req events.APIGatewayProxyRequest, handler interface{}) error {
	if len(req.Body) > MaxPayloadSize {
		return ErrTooLarge
	}
	dec := json.NewDecoder(strings.NewReader(req.Body))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&handler); err != nil {
		fmt.Printf("Error(getbody): %s\n", err.Error())
		return err
	}
	return nil
}

// OPENFAASGetBody check and deserialze the req payload in handler
func OPENFAASGetBody(r *http.Request, handler interface{}) error {
	var err error
	var body []byte
	if r == nil || r.Body == nil {
		return errors.New("invalid or null request body")
	}
	defer r.Body.Close()
	if body, err = ioutil.ReadAll(r.Body); err != nil {
		return nil
	}
	if len(body) > MaxPayloadSize {
		return ErrTooLarge
	}
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&handler); err != nil {
		fmt.Printf("Error(OPENFAASGetBody): %s\n", err.Error())
		return err
	}
	return nil
}

// LAMBDAGetQueryParam return a query param string or nil if not founds
func LAMBDAGetQueryParam(req events.APIGatewayProxyRequest, key string) *string {
	if tmp := req.QueryStringParameters[key]; tmp != "" {
		if goscrappy.Debug {
			fmt.Printf("Info(GetQueryParam): Got query string param [%s]=[%s]\n", key, tmp)
		}
		return &tmp
	}
	return nil
}

// LAMBDAGetPathParam return an url path parameter or nil if not found
func LAMBDAGetPathParam(req events.APIGatewayProxyRequest, key string) *string {
	if tmp := req.PathParameters[key]; tmp != "" {
		if goscrappy.Debug {
			fmt.Printf("Info(GetPathParam): Got path param [%s]=[%s]\n", key, tmp)
		}
		return &tmp
	}
	return nil
}
