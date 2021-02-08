package resolvers_test

import (
	"database/sql/driver"
	"time"
)

type AnyTime struct{}
type AvoidPassword struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

// Match satisfies sqlmock.Argument interface
func (password AvoidPassword) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}
