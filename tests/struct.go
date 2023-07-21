package tests

import "net/http"

type APITestCase struct {
	Name           string
	Request        *http.Request
	ExpectedStatus int
	ExpectedBody   string
}
