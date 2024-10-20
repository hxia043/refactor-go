package main

import (
	"fmt"
	"time"
)

type Plan struct {
	summerStart          time.Time
	summerEnd            time.Time
	summerRate           float64
	regularRate          float64
	regularServiceCharge float64
	quantity             float64
}

func (plan *Plan) isSummer(t time.Time) bool {
	if !t.Before(plan.summerStart) && !t.After(plan.summerEnd) {
		return true
	}

	return false
}

func (plan *Plan) summerCharge() float64 {
	return plan.quantity * plan.summerRate
}

func (plan *Plan) regularCharge() float64 {
	return plan.quantity * plan.regularRate * plan.regularServiceCharge
}

func main() {
	plan := Plan{
		summerStart:          time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC),
		summerEnd:            time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC),
		summerRate:           0.85,
		regularRate:          0.65,
		regularServiceCharge: 0.95,
		quantity:             0.75,
	}

	data := time.Now()

	charge := 0.0
	if plan.isSummer(data) {
		charge = plan.summerCharge()
	} else {
		charge = plan.regularCharge()
	}

	fmt.Println("charge: ", charge)
}
