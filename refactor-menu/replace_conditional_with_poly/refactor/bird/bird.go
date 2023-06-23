package bird

type bird interface {
	plumage() string
	airSpeedVelocity() int
}

type Bird struct {
	Name             string
	Kind             string
	NumberOfCoconuts int
	Voltage          int
	IsNailed         bool
}

func Plumage(bird Bird) string {
	switch bird.Kind {
	case "EuropeanSwallow":
		return NewEuropeanSwallow(bird).plumage()
	case "AfricanSwallow":
		return NewAfricanSwallow(bird).plumage()
	case "NorwegianBlueParrot":
		return NewNorwegianBlueParrot(bird).plumage()
	default:
		return "unknown"
	}
}

func AirSpeedVelocity(bird Bird) int {
	switch bird.Kind {
	case "EuropeanSwallow":
		return NewEuropeanSwallow(bird).airSpeedVelocity()
	case "AfricanSwallow":
		return NewAfricanSwallow(bird).airSpeedVelocity()
	case "NorwegianBlueParrot":
		return NewNorwegianBlueParrot(bird).airSpeedVelocity()
	default:
		return 0
	}
}

func New(name, kind string, numberOfCoconuts, voltage int, isNailed bool) Bird {
	return Bird{
		Name:             name,
		Kind:             kind,
		NumberOfCoconuts: numberOfCoconuts,
		Voltage:          voltage,
		IsNailed:         isNailed,
	}
}
