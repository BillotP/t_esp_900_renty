package function

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/BillotP/gorenty"
	"github.com/BillotP/gorenty/api"
	"github.com/BillotP/gorenty/v2/models"
	"github.com/BillotP/gorenty/v2/service"
)

// Query for user by id
type Query struct {
	ID string `json:"id"`
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
	var query Query
	var DB = dbservice.Db
	var results models.User
	var ctx = context.Background()
	api.OPENFAASGetBody(r, &query)
	if goscrappy.Debug {
		fmt.Printf("Query : %+v\n", query)
	}
	var qtemplate = `
	LET usr = DOCUMENT("users", "%s")
	LET a = (
		FOR comp IN companies FILTER comp == usr.company 
		RETURN comp
	)
	LET b = (
		FOR email IN emails FILTER email == usr.email 
		RETURN email
	)
	LET c = (
		FOR phone IN phones FILTER phone == usr.phone 
		RETURN phone
	)
	RETURN merge(usr, { 
			company: FIRST(a),
			email: FIRST(b),
			phone: FIRST(c)
		}
	)
	`
	q := fmt.Sprintf(qtemplate, query.ID)
	fmt.Printf("query : %s\n", q)
	cursor, err := DB.Query(ctx, q, nil)
	if err != nil {
		api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer cursor.Close()
	if _, err = cursor.ReadDocument(ctx, &results); err != nil {
		api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	api.OPENFAASJsonResponse(w, http.StatusOK, results)
}
