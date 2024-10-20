package main

import "fmt"

func call(name string, isHappy bool) string {
	if isHappy {
		return "xiaoming"
	} else {
		return name
	}
}

func main() {
	fmt.Println(call("xiaoxia", true))
	fmt.Println(call("xiaoxia", false))
}
