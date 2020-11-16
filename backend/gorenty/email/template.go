package email

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/BillotP/gorenty/v2/models"
)

// FillTemplate return a template string filled with the required variables
func FillTemplate(tofill string, params interface{}) (*string, error) {
	var err error
	var tpl bytes.Buffer
	var t *template.Template
	if t, err = template.New("email_template").Parse(tofill); err != nil {
		fmt.Printf("Failed to parse template : %s\n", err.Error())
		return nil, err
	}
	if err = t.Execute(&tpl, params); err != nil {
		return nil, err
	}
	result := tpl.String()
	return &result, nil
}

// GetEmail fill an email template with params and return an email ready to send
func GetEmail(tpl models.EmailTemplate, tplbody string, params interface{}) (*Email, error) {
	var err error
	var nemail Email
	var filled *string
	if strings.Contains(tpl.Subject, "{{") {
		if filled, err = FillTemplate(tpl.Subject, params); err != nil {
			return nil, err
		}
		nemail.Subject = *filled
	} else {
		nemail.Subject = tpl.Subject
	}

	if strings.Contains(tplbody, "{{") {
		if filled, err = FillTemplate(tplbody, params); err != nil {
			return nil, err
		}
		nemail.Body = *filled
	} else {
		nemail.Body = tplbody
	}
	if tpl.From != "" && strings.Contains(tpl.From, "{{") {
		if filled, err = FillTemplate(tpl.From, params); err != nil {
			return nil, err
		}
		nemail.From = []string{*filled}
	}
	return &nemail, nil
}
