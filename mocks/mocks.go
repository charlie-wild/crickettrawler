package mocks

//package for mock httpclient and mock dynamoDB

import (
	"net/http"
)

//GetDoFunc fetches the mock client 'Do' func
var GetDoFunc func(req *http.Request) (*http.Response, error)

//MockClient is the mock client
type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Do is the mock client's `Do` func
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}
