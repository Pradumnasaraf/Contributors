/*
Test case for addAContributor
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
	"github.com/magiconair/properties/assert"
)

func TestAddAContributor(t *testing.T) {

	r := gin.Default()
	r.POST("/graphql", handler.GraphqlHandler())

	// Test case for getAContributor
	testCases := []APITestCase{
		{
			Name: "Valid GraphQL Request - getAContributor",
			Request: func() *http.Request {
				requestBody := bytes.NewBufferString(`{"query": "mutation AddAContributor { addAContributor(input: { githubUsername: \"4\" name: \"user4\" email: \"test@test.com\" }) { userId } }"}`)
				req, _ := http.NewRequest("POST", "/graphql", requestBody)
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{"data": {"addAContributor": {"userId": "U4"}}}`,
		},
		{
			Name: "Valid GraphQL Request - getAContributor - with contributions",
			Request: func() *http.Request {
				requestBody := bytes.NewBufferString(`{"query": "mutation AddAContributor { addAContributor(input: { githubUsername: \"5\", name: \"user4\", email: \"test@test.com\" contributions: { projectName: \"test/test\", type: \"code\", date: \"2020-01-01\" } }) { userId contributions { contributionId } } }"}`)
				req, _ := http.NewRequest("POST", "/graphql", requestBody)
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{"data": {"addAContributor": {"userId": "U5","contributions": [{"contributionId": "Ctest/test"}]}}}`,
		},
		{
			Name: "Invalid GraphQL Request - getAContributor - without githubUsername",
			Request: func() *http.Request {
				requestBody := bytes.NewBufferString(`{"query": "mutation AddAContributor { addAContributor(input: { name: \"user4\" email: \"test@gmail.com\" }) { userId } }"}`)
				req, _ := http.NewRequest("POST", "/graphql", requestBody)
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			ExpectedStatus: http.StatusUnprocessableEntity,
			ExpectedBody:   `{"errors":[{"message":"Field \"NewContributor.githubUsername\" of required type \"String!\" was not provided.","locations":[{"line":1,"column":51}],"extensions":{"code":"GRAPHQL_VALIDATION_FAILED"}}],"data":null}`,
		},
	}

	// Run multiple test cases
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			w := httptest.NewRecorder()
			r.ServeHTTP(w, tc.Request)

			// Check the response status code
			assert.Equal(t, tc.ExpectedStatus, w.Code, "Status code")

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

	// Dlete the user created in the test from DB
	DeleteAddedContributor("U4")
	DeleteAddedContributor("U5")
}
