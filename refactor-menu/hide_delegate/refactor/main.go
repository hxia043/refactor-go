package main

type Person struct{}

type Department struct {
	manager string
}

func (p *Person) manager() string {
	return Department{}.manager
}

func main() {
	p := Person{}
	_ = p.manager()
}
