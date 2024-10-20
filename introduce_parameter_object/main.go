package main

import "fmt"

/*
	const station = {
		name: "ZB1",
		readings: [
		  { temp: 47, time: "2016-11-10 09:10" },
		  { temp: 53, time: "2016-11-10 09:20" },
		  { temp: 58, time: "2016-11-10 09:30" },
		  { temp: 53, time: "2016-11-10 09:40" },
		  { temp: 51, time: "2016-11-10 09:50" },
		],
	  };
*/

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

func readingsOutsideRange(station station, min, max int) {
	for _, reading := range station.readings {
		if reading.temp < min || reading.temp > max {
			fmt.Println(reading.temp)
		}
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
	readingsOutsideRange(s, o.temperatureFloor, o.temperatureCeiling)
}
