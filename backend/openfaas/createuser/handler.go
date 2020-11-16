package function

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BillotP/gorenty"

	"github.com/BillotP/gorenty/api"
	"github.com/BillotP/gorenty/v2/models"
	"github.com/BillotP/gorenty/v2/service"
)

// EmailInput is an email input
type EmailInput struct {
	Value       string `json:"value"`
	InfoConsent bool   `json:"infoConsent"`
}

// PhoneInput is a phone number input
type PhoneInput struct {
	CountryCode string `json:"countryCode"`
	Value       string `json:"value"`
}

// UserInput is a user registration input
type UserInput struct {
	Pseudo    string      `json:"pseudo"`
	Gender    string      `json:"gender"`
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	Password  string      `json:"password"`
	Email     *EmailInput `json:"email"`
	Phone     *PhoneInput `json:"phone"`
}

var dbservice *service.Repository

func init() {
	var err error
	var dbname = goscrappy.MustGetSecret("arango_dbname")
	if dbservice, err = service.New(dbname); err != nil {
		log.Fatal(err)
	}
}

// Handle a serverless request
func Handle(w http.ResponseWriter, r *http.Request) {
	var err error
	var query UserInput
	var nuser *models.UserItem
	var userdbo = service.UserItem{
		Repository: *dbservice,
	}
	api.OPENFAASGetBody(r, &query)
	if goscrappy.Debug {
		fmt.Printf("Query : %+v\n", query)
	}
	if nuser, err = userdbo.Create(models.User{
		Pseudo:    query.Pseudo,
		Gender:    query.Gender,
		FirstName: query.FirstName,
		LastName:  query.LastName,
		Password:  &query.Password,
	}); err != nil {
		api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	api.OPENFAASJsonResponse(w, http.StatusOK, nuser)
}
