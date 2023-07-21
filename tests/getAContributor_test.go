/*
Test case for getAContributor
*/
package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Pradumnasaraf/Contributors/handler"
	"github.com/gin-gonic/gin"
)

func TestGetAContributor(t *testing.T) {

	r := gin.Default()
	r.POST("/graphql", handler.GraphqlHandler())

	// Test case for getAContributor
	testCases := []APITestCase{
		{
			Name: "Valid GraphQL Request - getAContributor",
			Request: func() *http.Request {
				requestBody := bytes.NewBufferString(`{"query": "query GetAContributor { getAContributor(userId: \"U1\") { userId name } }"}`)
				req, _ := http.NewRequest("POST", "/graphql", requestBody)
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{"data": {"getAContributor": {"userId": "U1","name": "user1"}}}`,
		},
		{
			Name: "Valid GraphQL Request - getAContributor - Invalid userId",
			Request: func() *http.Request {
				requestBody := bytes.NewBufferString(`{"query": "query GetAContributor { getAContributor(userId: \"U100\") { userId name } }"}`)
				req, _ := http.NewRequest("POST", "/graphql", requestBody)
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{"errors":[{"message":"error while getting the document. Document with the given ID may not exist","path":["getAContributor"]}],"data":null}`,
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

			if !JSONEqual(response, expectedResponse) {
				t.Errorf("Expected response body:\n%s\nBut got:\n%s", tc.ExpectedBody, w.Body.String())
			}
		})
	}
}
