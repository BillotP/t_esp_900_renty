package function

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/arangodb/go-driver"
	"github.com/BillotP/gorenty"
	"github.com/BillotP/gorenty/api"
	"github.com/BillotP/gorenty/v2/models"
	"github.com/BillotP/gorenty/v2/service"
)

// Query for offeror by id
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
	var ctx = context.Background()
	var dbservice *service.Repository
	var DB = dbservice.Db
	var col driver.Collection
	var result models.Offeror
	api.OPENFAASGetBody(r, &query)
	if goscrappy.Debug {
		fmt.Printf("Query : %+v\n", query)
	}
	if col, err = DB.Collection(ctx, "offerors"); err != nil {
		api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	if _, err = col.ReadDocument(ctx, query.ID, &result); err != nil {
		api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	api.OPENFAASJsonResponse(w, http.StatusOK, result)
}
