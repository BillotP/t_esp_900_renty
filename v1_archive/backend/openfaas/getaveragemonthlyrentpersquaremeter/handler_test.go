package function

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// TestMain is the setup main before test
func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func TestHandle(t *testing.T) {
	var data io.Reader
	query := Query{
		SurfaceMin: 20,
		SurfaceMax: 200,
	}
	bodyByte, _ := json.Marshal(query)
	if blob := bytes.NewBuffer(bodyByte); blob != nil {
		data = blob
	}
	req, err := http.NewRequest("POST", "/", data)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handle)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	t.Logf("Resp : [%s]\n", rr.Body.String())
}
