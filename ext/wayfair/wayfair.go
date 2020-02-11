package wayfair

import (
	"github.com/m3tam3re/billbee/api/orders"
	"github.com/m3tam3re/billbee/api/products"
	"github.com/m3tam3re/csvhelper"
	"github.com/m3tam3re/errors"
	"github.com/spf13/viper"
	"strconv"
)

const path errors.Path = "github.com/m3tam3re/billbee/ext/wayfair"

// CreateOrder() will take an absolute path to a csv file and will create
// an &api.Order from it. Please note that this is function is only for
// processing csv order files from the Wayfair EDI ftp server. Orders exported
// from the Wayfair Extranet will not be working
//
// This function uses a yaml configuration file. Please hava a look at config.yaml.
func CreateOrder(file string) (*orders.Order, error) {
	const op errors.Op = "wayfair.go|func: CreateOrder()"

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.E(errors.NotExist, op, path, err, "could not read config")
	}
	order := orders.Order{}
	csv, err := csvhelper.GetLines(file, '|', false)
	if err != nil {
		return nil, errors.E(errors.Internal, op, path, err, "could not read lines from csv")
	}
	ln := 0
	for _, line := range csv {
		if ln == 0 {
			// line 0 is all about order information
			order.OrderNumber = line[viper.GetString("order.OrderNumber")]
			order.CreatedAt = line[viper.GetString("order.OrderDate")]
			order.CreatedAt = line[viper.GetString("order.OrderDate")]
			order.PaymentMethod = viper.GetInt32("order.PaymentMethod")
			if order.OrderNumber[:2] == "UK" {
				order.VatMode = 2
				order.Customer = &orders.Customer{
					Id:     73771437,
					Number: 293,
				}
			}
			if order.OrderNumber[:2] == "DE" {
				order.VatMode = 0
				order.Customer = &orders.Customer{
					Id:     73771431,
					Number: 287,
				}
			}

			order.ShippingAddress = &orders.ShippingAddress{
				LastName:    line[viper.GetString("order.ShippingAddress.LastName")],
				Street:      line[viper.GetString("order.ShippingAddress.Street")],
				Line2:       line[viper.GetString("order.ShippingAddress.Line2")],
				City:        line[viper.GetString("order.ShippingAddress.City")],
				Zip:         line[viper.GetString("order.ShippingAddress.Zip")],
				Country:     line[viper.GetString("order.ShippingAddress.Country")],
				CountryISO2: line[viper.GetString("order.ShippingAddress.Country")],
			}
			ln++
			continue
		}
		if ln > 0 {
			// lines 1+ are all about items
			quantity, err := strconv.ParseFloat(line[viper.GetString("order.OrderItems.Quantity")], 32)
			if err != nil {
				return nil, errors.E(errors.Internal, op, path, err, "could not convert Quantity to float.")
			}
			price, err := strconv.ParseFloat(line[viper.GetString("order.OrderItems.TotalPrice")], 32)
			if err != nil {
				return nil, errors.E(errors.Internal, op, path, err, "could not convert TotalPrice to float")
			}
			p := products.Product{}
			err = p.GetOne(line[viper.GetString("order.OrderItems.Product.SKU")])
			if err != nil {
				order.OrderItems = append(order.OrderItems, &orders.OrderItem{
					Product: &orders.Product{
						Title: line[viper.GetString("order.OrderItems.Product.SKU")],
					},
					Quantity:   float32(quantity),
					TotalPrice: float32(quantity * price * 1.19), // add Tax to product
				})
				ln++
				continue
			}
			order.OrderItems = append(order.OrderItems, &orders.OrderItem{
				Product: &orders.Product{
					SKU: line[viper.GetString("order.OrderItems.Product.SKU")],
				},
				Quantity:   float32(quantity),
				TotalPrice: float32(quantity * price * 1.19), // add Tax to product
			})
		}
		ln++
	}
	return &order, nil
}
