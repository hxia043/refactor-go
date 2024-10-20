package main

type Customer struct {
	name     string
	location string
}

func reportLines(aCustomer Customer) []string {
	lines := []string{}
	gatherCustomerData(lines, aCustomer)
	return lines
}

func gatherCustomerData(out []string, aCustomer Customer) {
	out = append(out, "name: "+aCustomer.name)
	out = append(out, "location: "+aCustomer.location)
}

func main() {
	reportLines(Customer{name: "hxia", location: "HangZhou"})
}
