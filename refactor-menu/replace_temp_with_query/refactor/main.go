package main

import "fmt"

type bill struct {
	_quantity int
	_item     item
}

type item struct {
	price int
}

func (b *bill) basePrice() float64 {
	return float64(b._quantity * b._item.price)
}

func (b *bill) discountFactor() float64 {
	var discountFactor = 0.98
	if b.basePrice() > 1000 {
		discountFactor -= 0.03
	}

	return discountFactor
}

func (b *bill) getPrice() float64 {
	return b.basePrice() * b.discountFactor()
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
