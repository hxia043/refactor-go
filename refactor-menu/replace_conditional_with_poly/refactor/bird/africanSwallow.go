package bird

type AfricanSwallow struct {
	name             string
	kind             string
	numberOfCoconuts int
	voltage          int
	isNailed         bool
}

func (as *AfricanSwallow) plumage() string {
	if as.numberOfCoconuts > 2 {
		return "tired"
	}

	return "average"
}

func (as *AfricanSwallow) airSpeedVelocity() int {
	return 40 - 2*as.numberOfCoconuts
}

func NewAfricanSwallow(bird Bird) bird {
	return &AfricanSwallow{
		name:             bird.Name,
		kind:             bird.Kind,
		numberOfCoconuts: bird.NumberOfCoconuts,
		voltage:          bird.Voltage,
		isNailed:         bird.IsNailed,
	}
}
