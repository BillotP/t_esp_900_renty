#!/usr/bin/env bash
set -e
if [ -z "$1" ]
  then
    echo "Missing new handler name, exiting"
    exit 1
fi
echo "[INFO] Creating new handler directory $1/handler"
mkdir -p "$1" && cd "$1"
echo "[INFO] Init go mod"
go mod init function
echo "[INFO] Creating OpenFAAS go-middleware handler template"
echo "
package function

import (
	\"fmt\"
	\"log\"
	\"net/http\"

	\"github.com/BillotP/gorenty\"

	\"github.com/BillotP/gorenty/api\"
	\"github.com/BillotP/gorenty/v2/models\"
	\"github.com/BillotP/gorenty/v2/service\"
)

// QueryType is the query received from api gateway (if any)
type QueryType struct {
    Var string \`json:\"var\"\`
}

var dbservice *service.Repository

func init() {
	var err error
	var dbname = goscrappy.MustGetSecret(\"arango_dbname\")
	if dbservice, err = service.New(dbname); err != nil {
		log.Fatal(err)
	}
}

// Handle a serverless request
func Handle(w http.ResponseWriter, r *http.Request) {
	var err error
    var response map[string]string
    var query QueryType
	api.OPENFAASGetBody(r, &query)
	if goscrappy.Debug {
		fmt.Printf(\"Query : %+v\n\", query)
	}
	if err != nil {
		api.OPENFAASErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	api.OPENFAASJsonResponse(w, http.StatusOK, response)
}
" > handler.go
echo "[INFO] Creating OpenFAAS go-middleware handler test template"
echo "
package function

import (
	\"net/http\"
	\"net/http/httptest\"
	\"os\"
	\"testing\"
)

// TestMain is the setup main before test
func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func TestHandle(t *testing.T) {
	req, err := http.NewRequest(\"POST\", \"/\", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handle)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(\"handler returned wrong status code: got %v want %v\",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	t.Logf(\"Resp : [%s]\n\", rr.Body.String())
}" > handler_test.go
echo "[INFO] Updating go mod"
go mod tidy
echo "[INFO] Update global stack.yml conf file"
cd ..
echo "  $1:
    lang: golang-middleware
    handler: ./$1
    image: repo.treescale.com/dave-lopeur/kubebeber/$1:latest
    environment:
      STAGE:   \"dev\"
      DEBUG: \"true\"
      CONTEXT: \"openfaas\"
    build_args:
      GO111MODULE: on
    secrets:
     - treescale-registry
     - arango-host
     - arango-port
     - arango-scheme
     - arango-user
     - arango-password
     - arango-tlsverify
     - arango-dbname" >> stack.yml
echo "[INFO] Done, happy hacking !"
