package function

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/arangodb/go-driver"
	"github.com/BillotP/renty/backend/lib"
	"github.com/BillotP/renty/backend/lib/api"
	"github.com/BillotP/renty/backend/lib/v2/models"
	"github.com/BillotP/renty/backend/lib/v2/service"
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
	var ctx = context.Background()
	var results []*models.RentOffer
	var DB = dbservice.Db
	api.OPENFAASGetBody(r, &query)
	if goscrappy.Debug {
		fmt.Printf("Query : %+v\n", query)
	}
	var qtemplate = `
	FOR ro IN rentoffers 
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
	LET ge = (
			FOR email IN emails FILTER email._key == FIRST(g).email
			RETURN email
	)
	LET he = (
			FOR phone IN phones FILTER phone._key == FIRST(g).phone
			RETURN phone
	)
	LET i = (
		FOR sur IN surfaces FILTER sur._key == ro.surface
		RETURN sur
	)
	%s
	LIMIT @offset, @limit
	RETURN merge(ro, { 
		title: a,
		source: FIRST(b),
		description: c,
		price: d,
		assets: e,
		location: FIRST(f),
		offeror: merge(FIRST(g), { 
			email: FIRST(ge), 
			phone: FIRST(he)
		}),
		surface: FIRST(i)
  })
	`
	if query != nil && query.Intext != nil {
		txtsearch = fmt.Sprintf(
			`
			LET srch = LOWER("%s")
			LET indesc = LOWER(FIRST(c).value)
			LET intitle = LOWER(FIRST(a).value)
			FILTER CONTAINS(indesc, srch) OR CONTAINS(intitle, srch)
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
		var doc models.RentOffer
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
