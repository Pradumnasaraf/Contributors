/*
Test case for deleteAContribution
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

func TestDeleteAContribution(t *testing.T) {

	r := gin.Default()
	r.POST("/graphql", handler.GraphqlHandler())

	// Test case for deleteAContribution
	testCases := []APITestCase{
		{
			Name: "Valid GraphQL Request - deleteAContribution",
			Request: func() *http.Request {
				requestBody := bytes.NewBufferString(`{"query": "mutation DeleteAContribution { deleteAContribution(userId: \"U8\", contributionId: \"C1\") { contributionId } }"}`)
				req, _ := http.NewRequest("POST", "/graphql", requestBody)
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{"data":{"deleteAContribution":{"contributionId":"C1"}}}`,
		},
		{
			Name: "Valid GraphQL Request - deleteAContribution - Invalid userId",
			Request: func() *http.Request {
				requestBody := bytes.NewBufferString(`{"query": "mutation DeleteAContribution { deleteAContribution(userId: \"U8\" contributionId: \"C1\") { contributionId } }"}`)
				req, _ := http.NewRequest("POST", "/graphql", requestBody)
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{"errors":[{"message":"document not found. Document with the given ID may not exist or contribution with the given ID may not exist","path":["deleteAContribution"]}],"data":null}`,
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

}
