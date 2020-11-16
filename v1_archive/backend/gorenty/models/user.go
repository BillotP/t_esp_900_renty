package models

import (
	"database/sql/driver"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/BillotP/gorenty"
	"gorm.io/gorm"
)

// Gender enumeration
type Gender int64

const (
	// MR is a daily report
	MR Gender = iota
	// MS is a monthly report
	MS
	// NA is a custom date range
	NA
)

var genders = []string{
	"mr",
	"ms",
	"na",
}

func (g Gender) String() string {
	return genders[g]
}

// GenderFromString get the gender type enum value from string
func GenderFromString(el string) Gender {
	val := sort.SearchStrings(genders, el)
	if val == len(el) {
		return NA
	}
	rt := Gender(val)
	if goscrappy.Debug {
		fmt.Printf("Info(FromString): Got %v gender type from %s string\n",
			rt, el)
	}
	return rt
}

// Scan a custom report type value
func (g *Gender) Scan(value interface{}) error { *g = Gender(value.(int64)); return nil }

// Value return a custom report type value
func (g Gender) Value() (driver.Value, error) { return int64(g), nil }

// Role enumeration
type Role int64

const (
	// ADMIN is an admin role
	ADMIN Role = iota
	// SERVICE is a service role
	SERVICE
	// USER is a classic role
	USER
)

var roles = []string{
	"admin",
	"service",
	"user",
}

func (r Role) String() string {
	return roles[r]
}

// RoleFromString get the gender type enum value from string
func RoleFromString(el string) Role {
	val := sort.SearchStrings(roles, el)
	if val == len(el) {
		return USER
	}
	rt := Role(val)
	if goscrappy.Debug {
		fmt.Printf("Info(FromString): Got %v role type from %s string\n",
			rt, el)
	}
	return rt
}

// Scan a custom report type value
func (r *Role) Scan(value interface{}) error { *r = Role(value.(int64)); return nil }

// Value return a custom report type value
func (r Role) Value() (driver.Value, error) { return int64(r), nil }

// User is a scrappy user
type User struct {
	Gender                Gender                 `json:"gender"`
	Role                  Role                   `json:"-"`
	FirstName             string                 `json:"firstName"`
	LastName              string                 `json:"lastName"`
	Email                 string                 `json:"email" gorm:"uniqueIndex"`
	Phone                 string                 `json:"phone"`
	APIKey                string                 `json:"apiKey"`
	ScrappingRequestItems []ScrappingRequestItem `json:"scrappingRequests"`
	ReportSubscriptions   string                 `json:"reportSub"`
	ReportsSub            []string               `json:"reportSubscriptions,omitempty" gorm:"-"`
}

// GetSubList create string array from reporttypes
func (u *User) GetSubList() {
	reportList := []string{}
	list := strings.Split(u.ReportSubscriptions, ",")
	for el := range list {
		item := FromString(list[el])
		reportList = append(reportList, item.String())
	}
	u.ReportsSub = reportList
	if goscrappy.Debug {
		fmt.Printf("Info(GetSubList): got reportsSub %+v\n", u.ReportsSub)
	}
}

// UserItem ...
type UserItem struct {
	Base
	User
}

// Valid check main template content
func (u User) Valid() bool {
	return (u.FirstName != "" && u.LastName != "" && u.Email != "")
}

// Save a new user in database
func (u User) Save(db *gorm.DB) (*UserItem, error) {
	if !u.Valid() { // Check for empty text
		return nil, ErrInvalidText
	}
	useritem := &UserItem{
		User: u,
	}
	if err := db.Create(useritem).Error; err != nil {
		return nil, err
	}
	return useritem, nil
}

// Update an existing user in database
func (u User) Update(db *gorm.DB, uuid uuid.UUID) (*UserItem, error) {
	item := UserItem{
		Base: Base{
			UUID: uuid,
		},
		User: u,
	}
	if goscrappy.Debug {
		fmt.Printf("Info(Update): Will update user [%+v]\n", item)
	}
	if err := db.Model(&item).Updates(&item).Error; err != nil {
		fmt.Printf("Error(Update): %s\n", err.Error())
		return nil, err
	}
	return &item, nil
}

// Delete a user in database
func (u *User) Delete(db *gorm.DB, uuid uuid.UUID) (*UserItem, error) {
	useritem := &UserItem{
		Base: Base{
			UUID: uuid,
		},
	}
	if err := db.Delete(useritem).Error; err != nil {
		return nil, err
	}
	return useritem, nil
}

// GetUserByID return a useritem by its id
func GetUserByID(db *gorm.DB, uuid uuid.UUID) (*UserItem, error) {
	var user UserItem
	if err := db.Where("id = ?", uuid).Find(&user).Error; err != nil {
		fmt.Printf("Error(GetUserByID): %s\n", err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	if err := db.Model(&user).
		Association("ScrappingRequestItems").
		Find(&user.ScrappingRequestItems); err != nil {
		fmt.Printf("Error(GetUserByID): %s\n", err.Error())
		return nil, err
	}
	user.GetSubList()
	return &user, nil
}

// GetUsers return all users in db
func GetUsers(db *gorm.DB, limit *string) (*[]UserItem, error) {
	var err error
	var users []UserItem
	if limit != nil {
		var limitval int
		if limitval, err = strconv.Atoi(*limit); err != nil {
			fmt.Printf("Error(GetUsers): Invalid limit param : %s\n", err.Error())
			return nil, ErrInvalidLimit
		} else if limitval < 1 || limitval > maxLimit {
			return nil, ErrInvalidLimit
		}
		if goscrappy.Debug {
			fmt.Printf("Info(GetUsers): Will looking for %v first users\n", limitval)
		}
		if err = db.Limit(limitval).Find(&users).Error; err != nil {
			fmt.Printf("Error(GetUsers): %s\n", err.Error())
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, err
		}

	} else if err = db.Find(&users).Error; err != nil {
		fmt.Printf("Error(GetUsers): %s\n", err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	for el := range users {
		users[el].GetSubList()
		if err = db.Model(&users[el]).
			Association("ScrappingRequestItems").
			Find(&users[el].ScrappingRequestItems); err != nil {
			fmt.Printf("Error(GetUsers): %s\n", err.Error())
			return nil, err
		}
		if goscrappy.Debug {
			fmt.Printf("Info(GetUsers): Found user [%+v]\n", users[el])
		}
	}
	return &users, nil
}
