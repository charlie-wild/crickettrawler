package restclient

import "net/http"

var Client HTTPClient

//HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

//sets client to instance of httpclient when it initialises. Init function runs once
//when package is imported
func init() {
	Client = &http.Client{}
}
