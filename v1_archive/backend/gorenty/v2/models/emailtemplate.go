package models

import "errors"

// EmailTemplate is a templated email
type EmailTemplate struct {
	Base
	Label        string                 `json:"label"`
	Subject      string                 `json:"value"`
	From         string                 `json:"from"`
	Body         Asset                  `json:"body"`
	TemplateVars map[string]interface{} `json:"template_vars"`
}

// EmailTemplateItem is the db model for a template stored in bucket
type EmailTemplateItem struct {
	Base
	Label        string                 `json:"label"`
	Subject      string                 `json:"value"`
	Body         string                 `json:"body"`
	From         string                 `json:"from"`
	TemplateVars map[string]interface{} `json:"template_vars"`
}

// Validate an email template before saving it
func (e EmailTemplate) Validate() error {
	if len(e.Label) == 0 {
		return errors.New("missing label element")
	}
	if len(e.Subject) == 0 {
		return errors.New("missing subject element")
	}
	return nil
}
