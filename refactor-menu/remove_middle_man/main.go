package main

type Person struct{}

type Department struct {
	manager string
	employ  string
}

func (p *Person) manager() string {
	return Department{}.manager
}

func (p *Person) employ() string {
	return Department{}.employ
}

func main() {
	p := Person{}
	_ = p.manager()
	_ = p.employ()
}
