package testing

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/afsxt/simple-vote/models"
	"github.com/afsxt/simple-vote/pkg/gredis"
	"github.com/afsxt/simple-vote/pkg/logging"
	"github.com/afsxt/simple-vote/pkg/setting"
	"github.com/afsxt/simple-vote/routers"
)

func setup() {
	setting.Setup("../conf/app.ini")
	models.Setup()
	gredis.Setup()
	logging.Setup()
}

func TestPing(t *testing.T) {
	setup()

	router := routers.InitRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestAddTheme(t *testing.T) {
	setup()

	var jsonData = []byte(`{
		"name": "morpheus",
		"description": "leader"
	}`)

	router := routers.InitRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/v1/vote/theme", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
