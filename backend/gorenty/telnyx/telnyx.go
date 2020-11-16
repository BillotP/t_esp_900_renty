package telnyx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/BillotP/gorenty"
)

/////////////////////
/////////////////////
///// TELNYX SMS API

// APIURL is the telnyx api url to send sms
const APIURL = "https://api.telnyx.com/v2/messages"

var (
	testMessage = `Bonjour,

Ce message est un test ou une erreur... 

Passez une agréable journée 🌞
`
	// ErrInvalidPhone is returned when phone is invalid
	ErrInvalidPhone = errors.New("phone is invalid")
	// ErrEmptyText is returned when text string is empty
	ErrEmptyText = errors.New("message is empty")
)

// Client is the interface to use telnyx api functionality
type Client struct {
	APIKey     string
	FromNumber string
}

// SMS is an sms model for Telnyx API
type SMS struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

// MessageResponseTo ...
type MessageResponseTo struct {
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
	Carrier     string `json:"carrier"`
	LineType    string `json:"line_type"`
}

// Error is an error response from telnyx api
type Error struct {
	Code   string            `json:"code"`
	Title  string            `json:"title"`
	Detail string            `json:"detail"`
	Meta   map[string]string `json:"meta"`
}

// MessageResponse is the response after a sms sending request
type MessageResponse struct {
	RecordType         string              `json:"record_type"`
	Direction          string              `json:"direction"`
	ID                 string              `json:"id"`
	Type               string              `json:"type"`
	OrganizationID     string              `json:"organization_id"`
	MessagingProfileID string              `json:"400172f0-47ba-455b-900e-543d7ee88145"`
	From               string              `json:"from"`
	To                 []MessageResponseTo `json:"to"`
	Text               string              `json:"text"`
	Media              []string            `json:"media"`
	WebhookURL         string              `json:"webhook_url"`
	WebhookFailoverURL string              `json:"webhook_failover_url"`
	Encoding           string              `json:"encoding"`
	Parts              int                 `json:"parts"`
	Tags               []string            `json:"tags"`
	Cost               *float64            `json:"cost"`
	ReceivedAt         time.Time           `json:"received_at"`
	SentAt             *time.Time          `json:"sent_at"`
	CompletedAt        *time.Time          `json:"completed_at"`
	ValidUntil         *time.Time          `json:"valid_until"`
	Errors             []Error             `json:"errors"`
}

// APIResponse is a wrapper for telnyx api response
type APIResponse struct {
	Errors []Error         `json:"errors"`
	Data   MessageResponse `json:"data"`
}

// hasErrors check if api response contains errors
func (r *APIResponse) hasErrors() bool {
	return len(r.Errors) > 0
}

// Errors return telnyx api error
func (r *APIResponse) Error() error {
	if !r.hasErrors() {
		return nil
	}
	var errfmt string
	for el := range r.Errors {
		errfmt += r.Errors[el].Title + ": " + r.Errors[el].Detail
	}
	return errors.New(errfmt)
}

// WithCountryCallingCode return a phone number with country calling code
func WithCountryCallingCode(phone string) string {
	const FRCountryCode = "33"
	if ind := strings.Index(phone, "0"); ind == 0 {
		return "+" + FRCountryCode + phone[1:10]
	} else if ind := strings.Index(phone, "+"); ind == 0 {
		return phone
	}
	return ""
}

// New return a telnyx sms client ready to use
func New(apiKey, fromNumber string) *Client {
	return &Client{
		APIKey:     apiKey,
		FromNumber: fromNumber,
	}
}

// SendSMS via telnyx api, return error if any
func (c *Client) SendSMS(to, text string) (rt *MessageResponse, err error) {
	var res *http.Response
	var req *http.Request
	var resDatas APIResponse
	var dataByte []byte
	var sms = SMS{
		From: c.FromNumber,
		To:   WithCountryCallingCode(to),
		Text: text,
	}
	if sms.To == "" {
		return nil, ErrInvalidPhone
	} else if sms.Text == "" {
		return nil, ErrEmptyText
	}
	if dataByte, err = json.Marshal(sms); err != nil {
		return nil, err
	}
	if goscrappy.Debug {
		fmt.Printf("Info(SendSMS): Will send sms %s\n", string(dataByte))
	}
	body := bytes.NewReader(dataByte)
	if req, err = http.NewRequest(http.MethodPost, APIURL, body); err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))
	if res, err = http.DefaultClient.Do(req); err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(resBytes, &resDatas); err != nil {
		return nil, err
	}
	if err := resDatas.Error(); err != nil {
		return nil, err
	}
	rt = &resDatas.Data
	return rt, err
}
