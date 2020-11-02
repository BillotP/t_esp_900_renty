package function

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BillotP/gorenty"
	"github.com/BillotP/gorenty/api"
	"github.com/BillotP/gorenty/bucket"
	"github.com/BillotP/gorenty/email"
	"github.com/BillotP/gorenty/v2/models"
	"github.com/BillotP/gorenty/v2/service"
)

// EmailQuery is the request to send a templated message to email in to field
type EmailQuery struct {
	To         string                 `json:"to"`
	TemplateID string                 `json:"template_id"`
	Datas      map[string]interface{} `json:"datas"`
}

var dbservice *service.Repository
var emailservice *email.Client
var bucketservice *bucket.Client

func init() {
	var err error
	var dbname = goscrappy.MustGetSecret("arango_assetdbname")
	if dbservice, err = service.New(dbname); err != nil {
		log.Fatal(err)
	}
	emailservice = email.New()
	bucketservice = bucket.New()
}

// Handle a serverless request
func Handle(w http.ResponseWriter, r *http.Request) {
	var err error
	var query EmailQuery
	var templatebody []byte
	var nemail *email.Email
	var tmpl *models.EmailTemplate
	api.OPENFAASGetBody(r, &query)
	if goscrappy.Debug {
		fmt.Printf("Query : %+v\n", query)
	}
	emailtemplaterepo := service.EmailTemplateItem{
		Repository: *dbservice,
	}
	if tmpl, err = emailtemplaterepo.GetByLabel(query.TemplateID); err != nil {
		api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	if templatebody, err = bucketservice.GetObject("renty-assets-dev", query.TemplateID+".html"); err != nil {
		api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	if nemail, err = email.GetEmail(*tmpl, string(templatebody), query.Datas); err != nil {
		api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	nemail.To = []string{query.To}
	go emailservice.Send(*nemail)
	w.WriteHeader(http.StatusOK)
}
