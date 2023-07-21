/*
Test case for getAllContributors
*/
package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Pradumnasaraf/Contributors/handler"
	"github.com/gin-gonic/gin"
)

func TestGetAllContributor(t *testing.T) {

	r := gin.Default()
	r.POST("/graphql", handler.GraphqlHandler())

	// Test case for getAllContributors
	testCases := []APITestCase{
		{
			Name: "Valid GraphQL Request",
			Request: func() *http.Request {
				requestBody := bytes.NewBufferString(`{"query": "query GetAllContributors { getAllContributors { userId } }"}`)
				req, _ := http.NewRequest("POST", "/graphql", requestBody)
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{"data":{"getAllContributors":[{"userId":"U1"}, {"userId":"U2"}, {"userId":"U3"}]}}`,
		},
	}

	// Run multiple test cases
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			w := httptest.NewRecorder()
			r.ServeHTTP(w, tc.Request)

			// Check the response status code
			if w.Code != tc.ExpectedStatus {
				t.Errorf("Expected status code %d, but got %d", tc.ExpectedStatus, w.Code)
			}

			// Check the response body
			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Error parsing response body: %v", err)
			}

			expectedJSON := []byte(tc.ExpectedBody)
			expectedResponse := make(map[string]interface{})
			err = json.Unmarshal(expectedJSON, &expectedResponse)
			if err != nil {
				t.Errorf("Error parsing expected response body: %v", err)
			}

			if !JSONContains(response, expectedResponse) {
				t.Errorf("Expected response body:\n%s\nBut got:\n%s", tc.ExpectedBody, w.Body.String())
			}
		})
	}
}

func JSONContains(a, b map[string]interface{}) bool {
	aJSON, err := json.Marshal(a)
	if err != nil {
		return false
	}

	aString := string(aJSON)

	bJSON, err := json.Marshal(b)
	if err != nil {
		return false
	}

	bString := string(bJSON)

	return strings.Contains(aString[:40], bString[:40])
}
