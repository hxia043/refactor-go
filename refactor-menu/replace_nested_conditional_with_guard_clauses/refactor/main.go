package main

var isDead = false
var isSeparated = false
var isRetired = false

func deadAmount() int {
	return 0
}

func separatedAmount() int {
	return 0
}

func retiredAmount() int {
	return 0
}

func normalPayAmount() int {
	return 0
}

func getPayAmount() int {
	if isDead {
		return deadAmount()
	}

	if isSeparated {
		return separatedAmount()
	}

	if isRetired {
		return retiredAmount()
	}

	return normalPayAmount()
}

type Instrument struct {
	capital          int
	interestRate     int
	duration         int
	income           int
	adjustmentFactor int
}

func adjustedCapital(anInstrument Instrument) int {
	if anInstrument.capital <= 0 || anInstrument.interestRate <= 0 || anInstrument.duration <= 0 {
		return 0
	}

	return (anInstrument.income / anInstrument.duration) * anInstrument.adjustmentFactor
}

func main() {
	getPayAmount()

	instrument := Instrument{10, 2, 5, 43, 8}
	adjustedCapital(instrument)
}
