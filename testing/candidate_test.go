package testing

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/afsxt/simple-vote/routers"
)

func TestAddCandidate(t *testing.T) {
	setup()

	var jsonData = []byte(`{
		"name": "morpheus",
		"description": "leader",
		"themeID": 1
	}`)

	router := routers.InitRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/v1/vote/candidates", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
