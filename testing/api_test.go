package testing

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/afsxt/simple-vote/routers"
)

func TestPing(t *testing.T) {
	setup()

	router := routers.InitRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

var themeTests = []struct {
	init           func(*http.Request)
	url            string
	method         string
	bodyData       string
	expectedCode   int
	responseRegexg string
	msg            string
}{
	{
		func(req *http.Request) {
			resetThemeTable()
		},
		"/admin/v1/vote/theme",
		"POST",
		`{"name": "test_name", "description": "test description"}`,
		http.StatusOK,
		`{"code":200,"msg":"ok","data":null}`,
		"",
	},
	{
		func(req *http.Request) {},
		"/admin/v1/vote/theme",
		"POST",
		`{"name": "test_name", "description": "test description"}`,
		http.StatusOK,
		`{"code":40001,"msg":"该主题已经存在","data":null}`,
		"",
	},
	{
		func(req *http.Request) {},
		"/admin/v1/vote/theme",
		"POST",
		`{"description": "test description"}`,
		http.StatusBadRequest,
		`{"code":400,"msg":"请求参数错误","data":null}`,
		"",
	},
	{
		func(req *http.Request) {
			resetCandidateTable()
		},
		"/admin/v1/vote/theme/1/state",
		"POST",
		`{"state": 1}`,
		http.StatusOK,
		`{"code":40006,"msg":"该主题候选人不够","data":null}`,
		"",
	},
	{
		func(req *http.Request) {},
		"/admin/v1/vote/candidates",
		"POST",
		`{"name":"candidate 1","description":"test candidate 1","themeID":1}`,
		http.StatusOK,
		`{"code":200,"msg":"ok","data":null}`,
		"",
	},
	{
		func(req *http.Request) {},
		"/admin/v1/vote/candidates",
		"POST",
		`{"name":"candidate 2","description":"test candidate 2","themeID":1}`,
		http.StatusOK,
		`{"code":200,"msg":"ok","data":null}`,
		"",
	},
	{
		func(req *http.Request) {},
		"/admin/v1/vote/theme/1/state",
		"POST",
		`{"state": 1}`,
		http.StatusOK,
		`{"code":200,"msg":"ok","data":null}`,
		"",
	},
	{
		func(req *http.Request) {
			resetUserTable()
		},
		"/api/v1/vote/verify",
		"POST",
		`{"email": "test"}`,
		http.StatusBadRequest,
		`{"code":400,"msg":"请求参数错误","data":null}`,
		"",
	},
	{
		func(req *http.Request) {},
		"/api/v1/vote/verify",
		"POST",
		`{"email": "test@126.com","idCard":"A12345688(7)"}`,
		http.StatusOK,
		`{"code":50003,"msg":"非法用户","data":null}`,
		"",
	},
	{
		func(req *http.Request) {},
		"/api/v1/vote/verify",
		"POST",
		`{"email": "test@126.com","idCard":"A123456(7)"}`,
		http.StatusOK,
		`{"code":200,"msg":"ok","data":null}`,
		"",
	},
	{
		func(req *http.Request) {},
		"/api/v1/vote/verify",
		"POST",
		`{"email": "test2@126.com","idCard":"B123456(7)"}`,
		http.StatusOK,
		`{"code":200,"msg":"ok","data":null}`,
		"",
	},
	{
		func(req *http.Request) {},
		"/api/v1/vote",
		"POST",
		`{"themeID": 3,"userID": 1, "candidateID": 1}`,
		http.StatusOK,
		`{"code":40004,"msg":"该主题不存在","data":null}`,
		"",
	},
	{
		func(req *http.Request) {
			resetVoteTable()
		},
		"/api/v1/vote",
		"POST",
		`{"themeID": 1,"userID": 1, "candidateID": 1}`,
		http.StatusOK,
		`{"code":200,"msg":"ok","data":null}`,
		"",
	},
	{
		func(req *http.Request) {},
		"/api/v1/vote",
		"POST",
		`{"themeID": 1,"userID": 1, "candidateID": 1}`,
		http.StatusOK,
		`{"code":60002,"msg":"该用户已经对该主题投过票","data":null}`,
		"",
	},
	{
		func(req *http.Request) {},
		"/api/v1/vote",
		"POST",
		`{"themeID": 1,"userID": 2, "candidateID": 1}`,
		http.StatusOK,
		`{"code":200,"msg":"ok","data":null}`,
		"",
	},
	{
		func(req *http.Request) {},
		"/api/v1/vote/theme/1",
		"GET",
		``,
		http.StatusOK,
		`{"code":200,"msg":"ok","data":{"lists":[{"CandidateID":1,"VoteCount":2}]}}`,
		"",
	},
	{
		func(req *http.Request) {},
		"/admin/v1/vote/theme/1",
		"GET",
		``,
		http.StatusOK,
		`{"code":200,"msg":"ok","data":{"lists":[{"CandidateID":1,"VoteCount":2}]}}`,
		"",
	},
}

func TestAddTheme(t *testing.T) {
	router := routers.InitRouter()
	asserts := assert.New(t)

	for _, tc := range themeTests {
		bodyData := tc.bodyData
		req, err := http.NewRequest(tc.method, tc.url, bytes.NewBufferString(bodyData))
		req.Header.Set("Content-Type", "application/json")
		asserts.NoError(err)

		tc.init(req)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		asserts.Equal(tc.expectedCode, w.Code, "Response Status - "+tc.msg)
		// asserts.Regexp(tc.responseRegexg, w.Body.String(), "Response Content - "+tc.msg)
		asserts.Equal(tc.responseRegexg, w.Body.String(), "Response Content - "+tc.msg)
	}
}
