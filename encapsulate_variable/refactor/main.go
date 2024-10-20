package main

import (
	"fmt"
	"owner/owner"
)

func main() {
	o := owner.GetDefaultOwner()
	owner.SetDefaultOwner(owner.New("huyun", "xia"))
	fmt.Println(o, owner.GetDefaultOwner())
}
