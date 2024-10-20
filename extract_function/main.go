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

func printOwing(invoice Invoice) {
	outstanding := 0

	fmt.Println("***********************")
	fmt.Println("**** Customer Owes ****")
	fmt.Println("***********************")

	// calculate outstanding
	for _, o := range invoice.orders {
		outstanding += o.amount
	}

	// record due date
	today := time.Now()
	invoice.dueData = today.Format(time.RFC3339)

	//print details
	fmt.Println("name: ", invoice.customer)
	fmt.Println("amount: ", outstanding)
	fmt.Println("due: ", invoice.dueData)
}

func main() {
	invoice := Invoice{
		orders: []Order{
			{amount: 43},
			{amount: 42},
		},
	}
	printOwing(invoice)
}
