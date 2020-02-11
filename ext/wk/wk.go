package wk

import (
	"fmt"
	"github.com/m3tam3re/billbee/api/orders"
	"github.com/m3tam3re/csvhelper"
	//"github.com/m3tam3re/billbee/api/products"
	"github.com/m3tam3re/errors"
	"github.com/spf13/viper"
)

const path errors.Path = "github.com/m3tam3re/ext/wk"

func CreateOrder(file string) (*orders.Order, error) {
	const op errors.Op = "wk.go|func: CreateOrder()"

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.E(errors.NotExist, op, path, err, "could not read config file")
	}

	order := orders.Order{}
	csv, err := csvhelper.GetLines(file, ',', true)
	if err != nil {
		return nil, errors.E(errors.Internal, op, path, err, "could not read csv file")

	}
	fmt.Println(csv)
	return &order, nil
}
