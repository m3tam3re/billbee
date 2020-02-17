package wk

import (
	"github.com/m3tam3re/billbee/api/orders"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	want := &orders.Order{
		OrderNumber:   "#1234",
		PaymentMethod: 3,
		VatMode:       0,
		Customer: &orders.Customer{
			Id:     73771747,
			Name:   "Weltkarten24.com",
			Email:  "peter.gruetzner@weltkarten24.com",
			Tel1:   "17621230980",
			Number: 602,
			Type:   0,
		},
		ShippingAddress: &orders.ShippingAddress{
			LastName:    "Max Mustermann",
			Street:      "Mustergasse 23",
			City:        "Musterhausen",
			Zip:         "12345",
			Country:     "DE",
			CountryISO2: "DE",
		},
		OrderItems: []*orders.OrderItem{
			{
				Product: &orders.Product{
					SKU: "weltkarte-sw-kork-l_60x40-1-S",
				},
				Quantity:   1,
				TotalPrice: 47.92,
			},
		},
	}
	order, err := CreateOrder("../../testdata/wk_sample.csv")
	if err != nil {
		t.Errorf("function CreateOrder() failed: %s", err)
	}
	t.Log(want, order)
}
