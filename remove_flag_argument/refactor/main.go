package main

import "fmt"

/*
Option 2: call with branch function
*/
func call(name string, isHappy bool) string {
	if isHappy {
		return happyCall()
	} else {
		return unhappyCall(name)
	}
}

/*
Option 1: decompose call with argument happyCall and unhappyCall
*/
func happyCall() string {
	return "xiaoming"
}

func unhappyCall(name string) string {
	return name
}

func main() {
	// option 1
	fmt.Println(happyCall())
	fmt.Println(unhappyCall("xiaoxia"))

	// option 2
	fmt.Println(call("xiaoxia", true))
	fmt.Println(call("xiaoxia", false))
}
