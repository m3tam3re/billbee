package api

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var shopId string = "69659"

func TestStartRequest(t *testing.T) {
	resp, err := startRequest("GET", "orders", nil)
	if err != nil {
		t.Errorf("Unexpected Error: %s", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("startRequest(\"GET\", \"orders\", nil) StatusCode: %v | Expected: %v", resp.StatusCode, 200)
	}
}

func TestOrder(t *testing.T) {
	order := Order{}
	fmt.Println(order)
}

func TestByExternalRef(t *testing.T) {
	o := Order{}
	err := o.ByExternalRef("FRR0815")
	if err != nil {
		t.Errorf("o.ByExternalRef(\"FRR0815\") failed | Error: %s", err)
	}
}

// Please set environt variable BILLBEE_TEST_ID to the desired Ordernumber for your test.
// If the OrderNumber exists no order will be created.
func TestCreateOrder(t *testing.T) {
	o := Order{
		AcceptLossOfReturnRight: false,
		Id:                      "",
		OrderNumber:             os.Getenv("BILLBEE_TEST_ID"),
		State:                   1,
		VatMode:                 0,
		CreatedAt:               time.Now().Format(time.RFC3339),
		ConfirmedAt:             time.Now().Add(time.Minute * 10).Format(time.RFC3339),
		ShippedAt:               time.Now().Add(time.Minute * 20).Format(time.RFC3339),
		PayedAt:                 time.Now().Add(time.Minute * 30).Format(time.RFC3339),
		SellerComment:           "Dies is ein Kommentar",
		InvoiceAddress: &InvoiceAddress{
			Company:     "Example Inc",
			Street:      "Musterstrasse",
			HouseNumber: "22b",
			Line2:       "Zusatz 1",
			Line3:       "Zusatz 2",
			City:        "Musterhausen",
			Zip:         "99999",
			State:       "Streng geheim",
			Country:     "DE",
			CountryISO2: "DE",
			FirstName:   "Max",
			LastName:    "Mustermann",
			Phone:       "0049 123456789",
			Email:       "dev@lanakk.com",
		},
		ShippingAddress: &ShippingAddress{
			Company:     "Example Inc",
			Street:      "Musterstrasse",
			HouseNumber: "22b",
			Line2:       "Zusatz 1",
			Line3:       "Zusatz 2",
			City:        "Musterhausen",
			Zip:         "99999",
			State:       "Streng geheim",
			Country:     "DE",
			CountryISO2: "DE",
			FirstName:   "Max",
			LastName:    "Mustermann",
			Phone:       "0049 123456789",
			Email:       "dev@lanakk.com",
		},
		ShippingCost:  4.90,
		PaymentMethod: 1,
		OrderItems: []*OrderItem{
			{
				Product: &Product{
					SKU: "weltkarte-graphit-kork-l_120x80-1-S",
					EAN: "",
				},
				Quantity:   2,
				TotalPrice: 44.66,
			},
			{
				Product: &Product{
					SKU: "weltkarte-oldstyle-l_60x40-1-S",
				},
				Quantity:   1,
				TotalPrice: 34.90,
				Discount:   20,
			},
		},
		Customer: &Customer{
			Id:     73771747,
			Number: 602,
		},
	}
	err := o.CreateOrder(shopId)
	if err != nil {
		t.Errorf("o.CreateOrder() failed | Error: %s", err)
	}
}
