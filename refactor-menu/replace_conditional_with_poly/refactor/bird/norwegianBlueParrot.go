package bird

type NorwegianBlueParrot struct {
	name             string
	kind             string
	numberOfCoconuts int
	voltage          int
	isNailed         bool
}

func (nbp *NorwegianBlueParrot) plumage() string {
	if nbp.voltage > 100 {
		return "scorched"
	}

	return "beautiful"
}

func (nbp *NorwegianBlueParrot) airSpeedVelocity() int {
	if nbp.isNailed {
		return 0
	}

	return 10 + nbp.voltage/10
}

func NewNorwegianBlueParrot(bird Bird) bird {
	return &AfricanSwallow{
		name:             bird.Name,
		kind:             bird.Kind,
		numberOfCoconuts: bird.NumberOfCoconuts,
		voltage:          bird.Voltage,
		isNailed:         bird.IsNailed,
	}
}
