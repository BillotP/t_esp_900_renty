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

// SearchQuery for rent offers
type SearchQuery struct {
	Offset *int64  `json:"offset"`
	Count  *int64  `json:"count"`
	Intext *string `json:"intext"`
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
	var txtsearch string
	var query *SearchQuery
	var DB = dbservice.Db
	var ctx = context.Background()
	var results []*models.Offeror
	api.OPENFAASGetBody(r, &query)
	if goscrappy.Debug {
		fmt.Printf("Query : %+v\n", query)
	}
	var qtemplate = `
	FOR of IN offerors
		%s
	    LIMIT @offset, @limit
	RETURN of
	`
	if query != nil && query.Intext != nil {
		txtsearch = fmt.Sprintf(
			`
			LET srch = LOWER("%s")
			LET inname = LOWER(of.name)
			FILTER CONTAINS(inname, srch)
			`,
			*query.Intext,
		)
	}
	// fmt.Printf("Got query : %s\n", q)
	limitVal := func() int64 {
		if query != nil && query.Count != nil &&
			*query.Count > 0 {
			return *query.Count
		}
		return 10
	}()
	offsetVal := func() int64 {
		if query != nil && query.Offset != nil &&
			*query.Offset > 0 {
			return *query.Offset
		}
		return 0
	}()
	q := fmt.Sprintf(qtemplate, txtsearch)
	cursor, err := DB.Query(ctx, q, map[string]interface{}{
		"offset": offsetVal,
		"limit":  limitVal,
	})
	if err != nil {
		api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer cursor.Close()
	for {
		var doc models.Offeror
		_, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
			return
		}
		results = append(results, &doc)
	}
	api.OPENFAASJsonResponse(w, http.StatusOK, results)
}
