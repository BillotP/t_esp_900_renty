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
	var ctx = context.Background()
	var query *SearchQuery
	var results []*models.User
	api.OPENFAASGetBody(r, &query)
	if goscrappy.Debug {
		fmt.Printf("Query : %+v\n", query)
	}
	var DB = dbservice.Db
	var qtemplate = `
	FOR usr IN users 
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
	%s
	LIMIT @offset, @limit
	RETURN merge(usr, { 
	  	company: FIRST(a),
	  	email: FIRST(b),
	  	phone: FIRST(c)
		}
	)
	`
	if query != nil && query.Intext != nil {
		txtsearch = fmt.Sprintf(
			`
			LET srch = LOWER("%s")
			LET inlastname = LOWER(usr.last_name)
			FILTER CONTAINS(inlastname, srch)
			`,
			*query.Intext,
		)
	}
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
		var doc models.User
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
