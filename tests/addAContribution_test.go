/*
Test case for addAContribution
*/
package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Pradumnasaraf/Contributors/config"
	"github.com/Pradumnasaraf/Contributors/graph"
	"github.com/Pradumnasaraf/Contributors/handler"
	"github.com/Pradumnasaraf/Contributors/mongo"
	"github.com/Pradumnasaraf/Contributors/redis"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

func TestAddAContribution(t *testing.T) {
	// Config setup
	config.Config()

	// Database connection
	redis.RedisInit()
	mongoClient := mongo.MongoInit()
	graph.GetMongoClient(mongoClient)
	defer redis.RedisClose()

	r := gin.Default()
	r.POST("/graphql", handler.GraphqlHandler())

	// Test case for adding a contribution
	testCases := []APITestCase{
		{
			Name: "Valid GraphQL Request - addAContribution",
			Request: func() *http.Request {
				requestBody := bytes.NewBufferString(`{"query": "mutation AddAContribution { addAContribution( userId: \"U6\" input: {projectName: \"test/test\", type: \"code\", date: \"2020-01-01\" }) { contributionId } }"}`)
				req, _ := http.NewRequest("POST", "/graphql", requestBody)
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   `{"data": {"addAContribution": {"contributionId": "Ctest/test"}}}`,
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
