// package billbeeapi is intended to deliver basic operations with the Bilbee API. For more
// information on Billbee visit https://www.bilbee.io
//
// Official documentation: https://app.billbee.io//swagger/ui/index#!/
//
// Credentials for using the API // please set up environment variables accordingly.
// BILBBE_API_URL 		https://app.billbee.io/api/v1/
// BILBBE_API_USER 		Your Billbee username
// BILBBE_API_PASS 		Can be set in your Billbee account
// BILBBE_API_KEY 		The API key is provided by Billbee

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

var endpoint string

// startRequest is helper function to create a HTTP request. It will return a pointer
// to a http request.
func startRequest(method string, endpoint string, body []byte) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Second * 30,
	}
	req, err := http.NewRequest(method, os.Getenv("BILBBE_API_URL")+endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("Could not create HTPP request: %s", err)
	}
	req.SetBasicAuth(os.Getenv("BILBBE_API_USER"), os.Getenv("BILBBE_API_PASS"))
	req.Header.Add("X-Billbee-Api-Key", os.Getenv("BILBBE_API_KEY"))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Could not execute HTTP request: %s", err)
	}
	return resp, nil
}

// ByExternalRef(ref string) consumes a string / external order reference.
// It will request the given order reference, store it in the Order struct.
func (o Order) ByExternalRef(ref string) error {
	endpoint = "/orders/findbyextref/" + ref
	resp, err := startRequest("GET", endpoint, nil)
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return fmt.Errorf("Could not find order: %s", ref)
	}
	if err != nil {
		return fmt.Errorf("Request could not be executed: %w", err)
	}
	//TODO marshall into models.Order
	var bodyData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&bodyData)
	//TODO End
	if err != nil {
		return fmt.Errorf("Could not decode response: %w", err)
	}
	return nil
}

// CreateOrder() will check if the given order number already exists. If it does not exists
// a new order will be created and the http Status Code will be returned.
// Please set environt variable BILLBEE_TEST_ID to the desired Ordernumber for your test.
// If the OrderNumber exists no order will be created.
func (o Order) CreateOrder(shopId string) error {
	err := o.ByExternalRef(o.OrderNumber)
	if err == nil {
		return fmt.Errorf("Cannot create order: %s - order already exists", o.OrderNumber)
	}
	endpoint = "/orders?shopId=" + shopId
	body, err := json.Marshal(o)
	fmt.Println(string(body))
	resp, err := startRequest("POST", endpoint, body)
	if err != nil {
		return fmt.Errorf("Could not execute the request: %s", err)
	}
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("Could not create order: %s", err)
	}
	return nil
}
