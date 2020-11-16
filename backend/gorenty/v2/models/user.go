package models

import "errors"

// User is a renty user
type User struct {
	Base
	Pseudo    string   `json:"pseudo"`
	Gender    string   `json:"gender"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Company   *Company `json:"company"`
	Email     *Email   `json:"email"`
	Phone     *Phone   `json:"phone"`
	Password  *string  `json:"password,omitempty"`
	OTP       *string  `json:"otp"`
}

// UserItem is a renty user in database
type UserItem struct {
	Base
	Pseudo       string  `json:"pseudo"`
	Gender       string  `json:"gender"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	PasswordHash string  `json:"password_hash"`
	Company      *string `json:"company"`
	Email        *string `json:"email"`
	Phone        *string `json:"phone"`
	OTP          *string `json:"otp"`
}

// Validate a user before saving
func (u User) Validate() error {
	if len(u.Pseudo) == 0 {
		return errors.New("missing pseudo element")
	}
	if len(u.Gender) == 0 {
		return errors.New("missing gender element")
	}
	if len(u.FirstName) == 0 {
		return errors.New("missing first name element")
	}
	if len(u.LastName) == 0 {
		return errors.New("missing last name element")
	}
	if u.Password == nil || (u.Password != nil && len(*u.Password) < 4) {
		return errors.New("missing or too short password element")
	}
	return nil
}
