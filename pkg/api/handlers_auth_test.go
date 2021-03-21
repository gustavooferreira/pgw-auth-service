package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gustavooferreira/pgw-auth-service/pkg/api"
	"github.com/gustavooferreira/pgw-auth-service/pkg/core/log"
	"github.com/gustavooferreira/pgw-auth-service/pkg/core/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFeeds(t *testing.T) {

	type RequestBody struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	}

	type ResponseBody struct {
		Valid bool `json:"valid"`
	}

	// Setup
	assert := assert.New(t)
	logger := log.NullLogger{}
	db := createCredsHolder()
	server := api.NewServer("", 9999, false, logger, db)
	router := server.Router

	// Table driven testing
	tests := map[string]struct {
		RequestBody          RequestBody
		expectedStatusCode   int
		expectedResponseBody ResponseBody
	}{
		"empty body": {
			RequestBody:        RequestBody{},
			expectedStatusCode: 400,
		},
		"valid user credentials": {
			RequestBody: RequestBody{
				Username: "bill",
				Password: "pass1"},
			expectedStatusCode: 200,
			expectedResponseBody: ResponseBody{
				Valid: true},
		},
		"invalid user credentials": {
			RequestBody: RequestBody{
				Username: "Bob",
				Password: "pass123"},
			expectedStatusCode: 200,
			expectedResponseBody: ResponseBody{
				Valid: false},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			requestBodyBytes, err := json.Marshal(test.RequestBody)
			require.NoError(t, err)
			responseBodyBytes, err := json.Marshal(test.expectedResponseBody)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			req, err := http.NewRequest("POST", "/api/v1/auth", bytes.NewBuffer(requestBodyBytes))
			require.NoError(t, err)
			router.ServeHTTP(w, req)

			require.Equal(t, test.expectedStatusCode, w.Code)

			if test.expectedStatusCode == 200 {
				assert.JSONEq(string(responseBodyBytes), w.Body.String())
			}
		})
	}
}

func createCredsHolder() *repository.CredsHolder {
	ch := repository.NewCredsHolder()

	ch.Credentials["bill"] = "pass1"
	ch.Credentials["adam"] = "pass2"
	ch.Credentials["john"] = "pass3"

	return &ch
}
