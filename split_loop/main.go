package main

import "fmt"

type person struct {
	age    int
	salary int
}

func main() {
	people := []person{
		{age: 30, salary: 1000},
		{age: 29, salary: 1200},
	}

	youngest := people[0].age
	totalSalary := 0

	for _, p := range people {
		if p.age < youngest {
			youngest = p.age
		}

		totalSalary += p.salary
	}

	fmt.Println("youngest age: ", youngest, " total salary: ", totalSalary)
}
