/*
Test case for deleteAContributor
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

func TestUpdateAContributor(t *testing.T) {

	r := gin.Default()
	r.POST("/graphql", handler.GraphqlHandler())

	// Test case for getAContributor

	/*
		mutation UpdateAContributor {
		updateAContributor(
			userId: "U9"
			input: {githubUsername: "10", name: "user10", email: "test@test.com"}
		) {
			userId
			githubUsername
		}
	*/
	testCases := []APITestCase{
		{
			Name: "Valid GraphQL Request - deleteAContributor",
			Request: func() *http.Request {
				requestBody := bytes.NewBufferString(`{"query": "mutation UpdateAContributor { updateAContributor(userId: \"U9\" input: {githubUsername: \"10\", name: \"user10\", email: \"test@test.com \"}) { userId githubUsername } }"}`)
				req, _ := http.NewRequest("POST", "/graphql", requestBody)
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{"data": {"updateAContributor": {"userId": "U9", "githubUsername": "10"}}}`,
		},
		// {
		// 	Name: "Valid GraphQL Request - deleteAContributor - Invalid userId",
		// 	Request: func() *http.Request {
		// 		requestBody := bytes.NewBufferString(`{"query": "mutation DeleteAContributor { deleteAContributor(userId: \"U100\") { userId name } }"}`)
		// 		req, _ := http.NewRequest("POST", "/graphql", requestBody)
		// 		req.Header.Set("Content-Type", "application/json")
		// 		return req
		// 	}(),
		// 	ExpectedStatus: http.StatusOK,
		// 	ExpectedBody:   `{"errors":[{"message":"error while deleting the document. Document with the given ID may not exist","path":["deleteAContributor"]}],"data":null}`,
		// },
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

	// Update the contributor back to original state
	UpdateModifiedContributor()

}
