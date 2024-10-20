package bird

type EuropeanSwallow struct {
	name             string
	kind             string
	numberOfCoconuts int
	voltage          int
	isNailed         bool
}

func (es *EuropeanSwallow) plumage() string {
	return "average"
}

func (es *EuropeanSwallow) airSpeedVelocity() int {
	return 35
}

func NewEuropeanSwallow(bird Bird) bird {
	return &AfricanSwallow{
		name:             bird.Name,
		kind:             bird.Kind,
		numberOfCoconuts: bird.NumberOfCoconuts,
		voltage:          bird.Voltage,
		isNailed:         bird.IsNailed,
	}
}
