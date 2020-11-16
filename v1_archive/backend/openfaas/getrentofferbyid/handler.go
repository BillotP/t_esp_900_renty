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

// Query for rent offer by id
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
	var ctx = context.Background()
	var query Query
	var results models.RentOffer
	var DB = dbservice.Db
	api.OPENFAASGetBody(r, &query)
	if goscrappy.Debug {
		fmt.Printf("Query : %+v\n", query)
	}
	var qtemplate = `
	LET ro = DOCUMENT("rentoffers", "%s")
	LET a = (
			FOR x IN ro.title
					FOR lb IN labels FILTER x == lb._key 
					RETURN lb
	)
	LET b = (
			FOR src IN sources FILTER ro.source == src._key
			RETURN src
	)
	LET c = (
			FOR x IN ro.description
					FOR des IN labels FILTER x == des._key 
					RETURN des
	)
	LET d = (
			FOR x IN ro.price
					FOR pr IN prices FILTER x == pr._key
					RETURN pr
	)
	LET e = (
			FOR x IN ro.assets
					FOR ass IN assets FILTER x == ass._key 
					RETURN ass
	)
	LET f = (
			FOR loc IN locations FILTER loc._key == ro.location 
			RETURN loc
	)
	LET g = (
			FOR off IN offerors FILTER off._key == ro.offeror 
			RETURN off
	)
	RETURN merge(ro, { 
		  title: a,
		  source: FIRST(b),
		  description: c,
		  price: d,
		  assets: e,
		  location: FIRST(f),
		  offeror: FIRST(g)
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
