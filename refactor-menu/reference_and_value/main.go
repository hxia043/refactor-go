package main

import "fmt"

type customer struct {
	name string
	menu [10]string
}

func (c *customer) order(name string) {
	for index, food := range c.menu {
		if food == "" {
			food = name
			c.menu[index] = food
			break
		}
	}
}

func New(name string) customer {
	return customer{name: name}
}

func main() {
	c := customer{name: "hxia"}
	c.order("cola chicken")

	fmt.Println(c)

	xiaoming := New("xiaoming")
	xiaohong := New("xiaohong")

	xiaoming.order("suibian")
	xiaohong.order("Liangpi")

	fmt.Println(xiaoming)
	fmt.Println(xiaohong)
}
