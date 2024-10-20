package main

import (
	"fmt"
	"time"
)

type Invoice struct {
	customer string
	orders   []Order
	dueData  string
}

type Order struct {
	amount int
}

func printBanner() {
	fmt.Println("***********************")
	fmt.Println("**** Customer Owes ****")
	fmt.Println("***********************")
}

func calculateAmount(invoice *Invoice) int {
	result := 0
	for _, o := range invoice.orders {
		result += o.amount
	}

	return result
}

func printDetails(invoice *Invoice) {
	fmt.Println("name: ", invoice.customer)
	fmt.Println("amount: ", calculateAmount(invoice))
	fmt.Println("due: ", invoice.dueData)
}

func recordDateTo(invoice *Invoice) {
	today := time.Now()
	invoice.dueData = today.Format(time.RFC3339)
}

func printOwing(invoice *Invoice) {
	printBanner()
	recordDateTo(invoice)
	printDetails(invoice)
}

func main() {
	invoice := &Invoice{
		customer: "hxia",
		orders: []Order{
			{amount: 43},
			{amount: 42},
		},
	}

	printOwing(invoice)
}
