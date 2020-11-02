package main

import (
	"net/http"
	"net/url"

	"github.com/BillotP/renty/backend/lib/models"
)

// Context is an authenticated request context
type Context struct {
	UserID  string
	User    models.User
	Request http.Request
	Writer  http.ResponseWriter
}

// SetCookie save a cookie for an auth user
func (c *Context) SetCookie(
	name string,
	value string,
	maxAge int,
	path string,
	domain string,
	secure bool,
	httpOnly bool,
) {
	if path == "" {
		path = "/"
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(value),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
}

// Cookie return a cookie from context
func (c *Context) Cookie(name string) (string, error) {
	cookie, err := c.Request.Cookie(name)
	if err != nil {
		return "", err
	}
	val, _ := url.QueryUnescape(cookie.Value)
	return val, nil
}
