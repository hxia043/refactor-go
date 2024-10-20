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
	result := 0

	if isDead {
		result = deadAmount()
	} else {
		if isSeparated {
			result = separatedAmount()
		} else {
			if isRetired {
				result = retiredAmount()
			} else {
				result = normalPayAmount()
			}
		}
	}

	return result
}

type Instrument struct {
	capital          int
	interestRate     int
	duration         int
	income           int
	adjustmentFactor int
}

func adjustedCapital(anInstrument Instrument) int {
	result := 0
	if anInstrument.capital > 0 {
		if anInstrument.interestRate > 0 && anInstrument.duration > 0 {
			result = (anInstrument.income / anInstrument.duration) * anInstrument.adjustmentFactor
		}
	}

	return result
}

func main() {
	getPayAmount()

	instrument := Instrument{10, 2, 5, 43, 8}
	adjustedCapital(instrument)
}
