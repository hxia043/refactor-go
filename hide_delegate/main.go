package main

type Person struct{}

type Department struct {
	manager string
}

func (p *Person) department() Department {
	return Department{}
}

func main() {
	p := Person{}
	_ = p.department().manager
}
