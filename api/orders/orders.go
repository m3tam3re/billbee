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
//
// If you need additional fields on the Order struct please check and extend the model/orders.proto
// file and recompile it
package orders

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/m3tam3re/billbee/api"
	"github.com/m3tam3re/errors"
)

const path errors.Path = "github.com/m3tam3re/billbee/api/orders/orders.go"

// ByExternalRef(ref string) consumes a string / external order reference.
// It will request the given order reference, store it in the Order struct.
func (o Order) ByExternalRef(ref string) error {
	const op errors.Op = "orders.go|method: ByExternalRef"

	endpoint := "/orders/findbyextref/" + ref
	resp, err := api.StartRequest("GET", endpoint, nil)
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return errors.E(errors.NotExist, op, path, err, fmt.Sprintf("statuscode should be 200, got: %v", resp.StatusCode))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.E(errors.Internal, op, path, err, "could not read response body")
	}
	var respmap map[string]json.RawMessage
	err = json.Unmarshal(body, &respmap)
	if err != nil {
		return errors.E(errors.Internal, op, path, err, "could not unmarshal response")
	}
	err = json.Unmarshal(respmap["Data"], &o)
	if err != nil {
		return errors.E(errors.Internal, op, path, err, "could not unmarshal 'Data' from response")
	}
	return nil
}

// Create() will check if the given order number already exists. If it does not exists
// a new order will be created and the http Status Code will be returned.
// Please set environt variable BILLBEE_TEST_ID to the desired Ordernumber for your test.
// If the OrderNumber exists no order will be created.
func (o Order) Create(shopId string) error {
	const op errors.Op = "orders.go|method: Create"

	err := o.ByExternalRef(o.OrderNumber)
	if err == nil {
		return errors.E(errors.Exists, op, path, "order already exists")
	}
	endpoint := "/orders?shopId=" + shopId
	body, err := json.Marshal(o)
	resp, err := api.StartRequest("POST", endpoint, body)
	defer resp.Body.Close()
	if err != nil {
		return errors.E(errors.Internal, op, path, err, "POST request failed")
	}
	// TODO give a useful return | response?
	return nil
}
