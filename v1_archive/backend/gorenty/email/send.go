package email

import (
	"fmt"
	"net/smtp"
	"strings"
	"time"

	"github.com/BillotP/gorenty"

	"github.com/dgryski/trifles/uuid"
)

// Client is the email client with the smtp server credentials
type Client struct {
	Host     string
	Port     string
	User     string
	Password string
	Domain   string
}

// Email is an email to send
type Email struct {
	From    []string `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

// New setup smtp config or exit if failed
func New() *Client {
	var client Client
	client.Host = goscrappy.MustGetSecret("smtp_host")
	client.Port = goscrappy.MustGetSecret("smtp_port")
	client.User = goscrappy.MustGetSecret("smtp_user")
	client.Password = goscrappy.MustGetSecret("smtp_password")
	client.Domain = goscrappy.MustGetSecret("smtp_domain")
	return &client
}

// Send  a complete text or html email with go smtp package
func (c *Client) Send(email Email) chan error {
	var err error
	const CRLF = "\r\n"
	const rfc2822 = "Mon, 02 Jan 2006 15:04:05 -0700"
	var errs = make(chan error, 1)
	smtpAuth := New()
	defer close(errs)
	auth := smtp.CRAMMD5Auth(c.User, c.Password)
	mime := "MIME-version: 1.0;" + CRLF
	tomsg := "To: " + strings.Join(email.To, ",") + CRLF
	frommsg := "From: " + strings.Join(email.From, ",") + CRLF
	subjectmsg := "Subject: " + email.Subject + CRLF
	date := "Date: " + time.Now().Format(rfc2822) + CRLF
	messageID := "Message-ID: <" + uuid.UUIDv4() + "@" + c.Domain + ">" + CRLF
	contentTransferEncoding := "Content-Transfer-Encoding: 8bit" + CRLF
	contentType := "Content-Type: text/html; charset=\"UTF-8\";" + CRLF
	fullBody := fmt.Sprintf("%s%s%s%s%s%s%s%s%s",
		date,
		tomsg,
		frommsg,
		subjectmsg,
		messageID,
		mime,
		contentTransferEncoding,
		contentType,
		email.Body,
	)
	logmsg := fmt.Sprintf("Will send mail size %v via %s:%s from %s to %s",
		len(email.Body),
		smtpAuth.Host,
		smtpAuth.Port,
		email.From,
		email.To,
	)
	fmt.Printf("Info(SendEmail): %s\n", logmsg)
	if err = smtp.SendMail(
		c.Host+":"+c.Port,
		auth,
		c.User,
		email.To,
		[]byte(fullBody)); err != nil {
		fmt.Printf("Error(SendEmail): %s\n", err.Error())
		errs <- err
	}
	return errs
}
