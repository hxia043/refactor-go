package main

import (
	"fmt"
	"time"
)

type Data struct{}

func (data *Data) isBefore(t time.Time) bool {
	return time.Now().Before(t)
}

func (data *Data) isAfter(t time.Time) bool {
	return time.Now().After(t)
}

type Plan struct {
	summerStart          time.Time
	summerEnd            time.Time
	summerRate           float64
	regularRate          float64
	regularServiceCharge float64
}

func main() {
	quantity := 0.75
	data := Data{}
	plan := Plan{
		summerStart:          time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC),
		summerEnd:            time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC),
		summerRate:           0.85,
		regularRate:          0.65,
		regularServiceCharge: 0.95,
	}

	charge := 0.0
	if !data.isBefore(plan.summerStart) && !data.isAfter(plan.summerEnd) {
		charge = quantity * plan.summerRate
	} else {
		charge = quantity * plan.regularRate * plan.regularServiceCharge
	}

	fmt.Println("charge: ", charge)
}
