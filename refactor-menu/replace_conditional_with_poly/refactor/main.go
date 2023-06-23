package main

import (
	"bird/bird"
	"fmt"
)

func plumages(birds map[string]bird.Bird) {
	for _, b := range birds {
		fmt.Println(bird.Plumage(birds[b.Name]))
	}
}

func speeds(birds map[string]bird.Bird) {
	for _, b := range birds {
		fmt.Println(bird.AirSpeedVelocity(birds[b.Name]))
	}
}

func main() {
	birds := map[string]bird.Bird{
		"xiaohei":  bird.New("xiaohei", "AfricanSwallow", 4, 120, false),
		"xiaohong": bird.New("xiaohong", "NorwegianBlueParrot", 2, 100, true),
	}

	plumages(birds)
	speeds(birds)
}
