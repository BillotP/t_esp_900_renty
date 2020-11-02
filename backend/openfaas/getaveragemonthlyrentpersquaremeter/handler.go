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

// Query ...
type Query struct {
	SurfaceMin float64 `json:"surface_min"`
	SurfaceMax float64 `json:"surface_max"`
}

// AveragePricePerSquareMeter ...
type AveragePricePerSquareMeter struct {
	Value      float64 `json:"value"`
	Unit       string  `json:"unit"`
	TotalElems int     `json:"total_elements"`
	Error      *string `json:"error"`
}

// ReducedRentOffer ...
type ReducedRentOffer struct {
	ID      string         `json:"_key"`
	Surface models.Surface `json:"surface"`
	Price   []models.Price `json:"price"`
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
	var results []*ReducedRentOffer
	var DB = dbservice.Db
	var query Query
	var resp AveragePricePerSquareMeter
	if err = api.OPENFAASGetBody(r, &query); err != nil {
		api.OPENFAASErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if goscrappy.Debug {
		fmt.Printf("Query : %+v\n", query)
	}
	var q = `
	FOR ro IN rentoffers
	LET d = (
		FOR x IN ro.price
		FOR pr IN prices FILTER x == pr._key
		RETURN pr
	)
	LET i = (
		FOR sur IN surfaces FILTER sur._key == ro.surface
		RETURN sur
	)
	RETURN merge(ro, { 
		price: d,
		surface: FIRST(i)
  	})
	`
	cursor, err := DB.Query(ctx, q, nil)
	if err != nil {
		api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	defer cursor.Close()
	for {
		var doc ReducedRentOffer
		_, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
			return
		}
		if doc.Surface.Value >= query.SurfaceMin &&
			doc.Surface.Value <= query.SurfaceMax {
			fmt.Printf("Adding surface %v for price %v\n",
				doc.Surface.Value,
				doc.Price[0].Value)
			results = append(results, &doc)
		}
	}
	if len(results) == 0 {
		errmsg := "not enough data to get acurate estimation"
		resp.Error = &errmsg
		api.OPENFAASJsonResponse(w, http.StatusOK, resp)
		return
	}
	serieSum, serieLen := float64(0), float64(0)
	for el := range results {
		serieSum += results[el].Price[0].Value
		serieLen++
	}
	resp.Value = serieSum / serieLen
	resp.Unit = "€ / month"
	resp.TotalElems = int(serieLen)
	api.OPENFAASJsonResponse(w, http.StatusOK, resp)
}
