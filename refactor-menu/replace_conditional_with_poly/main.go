package main

import "fmt"

type bird struct {
	name             string
	kind             string
	numberOfCoconuts int
	voltage          int
	isNailed         bool
}

func plumage(bird bird) string {
	switch bird.kind {
	case "EuropeanSwallow":
		return "average"
	case "AfricanSwallow":
		if bird.numberOfCoconuts > 2 {
			return "tired"
		}

		return "average"
	case "NorwegianBlueParrot":
		if bird.voltage > 100 {
			return "scorched"
		}

		return "beautiful"
	default:
		return "unknown"
	}
}

func plumages(birds map[string]bird) {
	for _, b := range birds {
		fmt.Println(plumage(birds[b.name]))
	}

}

func airSpeedVelocity(bird bird) int {
	switch bird.kind {
	case "EuropeanSwallow":
		return 35
	case "AfricanSwallow":
		return 40 - 2*bird.numberOfCoconuts
	case "NorwegianBlueParrot":
		if bird.isNailed {
			return 0
		}

		return 10 + bird.voltage/10
	default:
		return 0
	}
}

func speeds(birds map[string]bird) {
	for _, b := range birds {
		fmt.Println(airSpeedVelocity(birds[b.name]))
	}
}

func main() {
	birds := map[string]bird{
		"xiaohei": {
			name:             "xiaohei",
			kind:             "AfricanSwallow",
			numberOfCoconuts: 4,
			voltage:          120,
			isNailed:         false,
		},
		"xiaohong": {
			name:             "xiaohong",
			kind:             "NorwegianBlueParrot",
			numberOfCoconuts: 2,
			voltage:          100,
			isNailed:         true,
		},
	}

	plumages(birds)
	speeds(birds)
}
