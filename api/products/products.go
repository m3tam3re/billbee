package products

import (
	"encoding/json"
	"fmt"
	"github.com/m3tam3re/billbee/api"
	"github.com/m3tam3re/errors"
	"io/ioutil"
	"net/http"
)

const path errors.Path = "github.com/m3tam3re/billbee/api/products/products.go"

// GetOne(prod string) will lookup a single product in the Billbee API and store the
// datafields of the response into the Product struct. Fields not included in the struct
// will be ignored. If you need additional fields please check and extend the models/products.proto
// file and recompile it.
func (p *Product) GetOne(prod string) error {
	const op errors.Op = "products.go|method: GetOne"

	endpoint := "products/" + prod + "?lookupBy=sku"
	resp, err := api.StartRequest("GET", endpoint, nil)
	defer resp.Body.Close()
	if err != nil {
		return errors.E(errors.Internal, op, path, err, "error while sending request")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.E(errors.NotExist, op, path, err, fmt.Sprintf("statuscode should be 200, got: %v", resp.StatusCode))
	}
	body, err := ioutil.ReadAll(resp.Body)
	var respmap map[string]json.RawMessage
	err = json.Unmarshal(body, &respmap)
	if err != nil {
		return errors.E(errors.Internal, op, path, err, "could not unmarshal response")
	}
	err = json.Unmarshal(respmap["Data"], &p)
	if err != nil {
		return errors.E(errors.Internal, op, path, err, "could not unmarshal 'Data' from response")
	}
	return nil
}
