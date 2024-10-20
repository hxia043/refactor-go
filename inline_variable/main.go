package main

import "fmt"

func getPrice() int {
	return 10
}

func main() {
	price := getPrice()
	fmt.Println("price: ", price)
}
