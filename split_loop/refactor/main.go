package main

import "fmt"

type person struct {
	age    int
	salary int
}

func totalSalary(people []person) int {
	result := 0
	for _, p := range people {
		result += p.salary
	}

	return result
}

func youngestAge(people []person) int {
	result := people[0].age
	for _, p := range people {
		if p.age < result {
			result = p.age
		}
	}

	return result
}

func main() {
	people := []person{
		{age: 30, salary: 1000},
		{age: 29, salary: 1200},
	}

	fmt.Println("youngest age: ", youngestAge(people), " total salary: ", totalSalary(people))
}
