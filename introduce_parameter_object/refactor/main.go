package main

import "fmt"

type station struct {
	name     string
	readings []reading
}

type reading struct {
	temp int
	time string
}

type operatingPlan struct {
	temperatureFloor   int
	temperatureCeiling int
}

type Range struct {
	min int
	max int
}

func (r *Range) contain(v int) {
	if v < r.min || v > r.max {
		fmt.Println(v)
	}
}

func readingsOutsideRange(station station, r Range) {
	for _, reading := range station.readings {
		r.contain(reading.temp)
	}
}

func main() {
	s := station{
		name: "ZB1",
		readings: []reading{
			{temp: 47, time: "2016-11-10 09:10"},
			{temp: 53, time: "2016-11-10 09:20"},
			{temp: 58, time: "2016-11-10 09:30"},
			{temp: 53, time: "2016-11-10 09:40"},
			{temp: 51, time: "2016-11-10 09:50"},
		},
	}

	o := operatingPlan{temperatureFloor: 45, temperatureCeiling: 50}

	r := Range{
		min: o.temperatureFloor,
		max: o.temperatureCeiling,
	}

	readingsOutsideRange(s, r)
}
