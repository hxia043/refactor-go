package main

type Customer struct {
	name     string
	location string
}

func reportLines(aCustomer Customer) []string {
	lines := []string{}
	lines = append(lines, "name: "+aCustomer.name)
	lines = append(lines, "location: "+aCustomer.location)
	return lines
}

func main() {
	reportLines(Customer{name: "hxia", location: "HangZhou"})
}
