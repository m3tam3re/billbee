package wayfair

import (
	"github.com/m3tam3re/billbee/api"
	"reflect"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	want := &api.Order{
		OrderNumber:   "UK12345678",
		CreatedAt:     "2019-10-18",
		PaymentMethod: 22,
		VatMode:       2,
		Customer: &api.Customer{
			Id:     73771437,
			Number: 293,
		},
		ShippingAddress: &api.ShippingAddress{
			LastName:    "Max Mustermann",
			Street:      "1 Muster Street",
			Line2:       "This is line2",
			City:        "Musterhausen",
			Zip:         "RG12 3YA",
			Country:     "GB",
			CountryISO2: "GB",
		},
		OrderItems: []*api.OrderItem{
			{
				Product: &api.Product{
					SKU: "glamour-g_120x80-1-FG",
				},
				Quantity:   1,
				TotalPrice: 112.9905,
			},
			{
				Product: &api.Product{
					SKU: "weltkarte-graphit-g_120x80-1-FG",
				},
				Quantity:   2,
				TotalPrice: 9.401001,
			},
		},
	}
	order, err := CreateOrder("../../testdata/wayfair_sample.csv")
	if err != nil {
		t.Errorf("Function CreateOrder failed: %s", err)
	}
	if !reflect.DeepEqual(want, order) {
		t.Errorf("Did not get expected values:\nWanted: %v\n\nGot: %v", want, order)
	}
}
