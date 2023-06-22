package main

import "fmt"

type bill struct {
	_quantity int
	_item     item
}

type item struct {
	price int
}

func (b *bill) getPrice() float64 {
	var basePrice = b._quantity * b._item.price
	var discountFactor = 0.98
	if basePrice > 1000 {
		discountFactor -= 0.03
	}

	return float64(basePrice) * discountFactor
}

func main() {
	b := bill{
		_quantity: 43,
		_item: item{
			price: 1000,
		},
	}

	fmt.Println(b.getPrice())
}
