package main

import (
	"fmt"
	"math"
)

type Order struct {
	_data Record
}

type Record struct {
	quantity  float64
	itemPrice float64
}

func (o *Order) quantity() float64 {
	return o._data.quantity
}

func (o *Order) itemPrice() float64 {
	return o._data.itemPrice
}

func (o *Order) getPrice() float64 {
	return o.quantity()*o.itemPrice() - math.Max(0, o.quantity()-500)*o.itemPrice()*0.05 + math.Min(o.quantity()*o.itemPrice()*0.1, 100)
}

func main() {
	o := Order{
		_data: Record{
			quantity:  43.0,
			itemPrice: 100,
		},
	}

	fmt.Println(o.getPrice())
}
