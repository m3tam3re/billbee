package api

import (
	"bytes"
	"net/http"
	"os"
	"time"

	"github.com/m3tam3re/errors"
)

const path errors.Path = "github.com/m3tam3re/billbee/api/global.go"

// startRequest is helper function to create a HTTP request. It will return a pointer
// to a http request.
func StartRequest(method string, endpoint string, body []byte) (*http.Response, error) {
	const op errors.Op = "func: startRequest"

	client := http.Client{
		Timeout: time.Second * 30,
	}
	req, err := http.NewRequest(method, os.Getenv("BILBBE_API_URL")+endpoint, bytes.NewBuffer(body))
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
