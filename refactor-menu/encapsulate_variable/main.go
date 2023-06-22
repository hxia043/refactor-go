package main

import "fmt"

type Owner struct {
	firstName string
	lastName  string
}

var defaultOwner = Owner{firstName: "Martin", lastName: "Fowler"}

func main() {
	owner := defaultOwner
	defaultOwner = Owner{firstName: "Rebecca", lastName: "Parsons"}
	fmt.Println(owner, defaultOwner)
}
