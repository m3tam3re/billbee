package api

import (
	"bytes"
	"github.com/m3tam3re/errors"
	"net/http"
	"net/url"
	"os"
	"time"
)

const path errors.Path = "github.com/m3tam3re/billbee/api/global.go"

// startRequest is helper function to create a HTTP request. It will return a pointer
// to a http request.
func StartRequest(method string, endpoint string, body []byte) (*http.Response, error) {
	const op errors.Op = "func: startRequest"
	baseUrl := os.Getenv("BILBBE_API_URL")
	if (os.Getenv("BILBBE_API_URL"))[len(os.Getenv("BILBBE_API_URL"))-1:] == endpoint[:1] {
		baseUrl = os.Getenv("BILBBE_API_URL")[:len(os.Getenv("BILBBE_API_URL"))-1]
	}
	client := http.Client{
		Timeout: time.Second * 120,
	}
	reqUrl, err := url.Parse(baseUrl + endpoint)
	if err != nil {
		return nil, errors.E(errors.Internal, op, path, err, "could not parse URL")
	}
	req, err := http.NewRequest(method, reqUrl.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.E(errors.Internal, op, path, err, "error building request")
	}
	req.SetBasicAuth(os.Getenv("BILBBE_API_USER"), os.Getenv("BILBBE_API_PASS"))
	req.Header.Add("X-Billbee-Api-Key", os.Getenv("BILBBE_API_KEY"))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.E(errors.Internal, op, path, err, "error executing request")
	}
	return resp, nil
}
