package statement

import (
	"fmt"
	"statement/performance"
)

var invoice = make(map[string]interface{})
var plays = make(map[string]interface{})

type statementData struct {
	customer           string
	performances       []map[string]interface{}
	totalAmount        int
	totalVolumeCredits float64
}

func getInvoice() map[string]interface{} {
	return invoice
}

func setInvoice(data map[string]interface{}) {
	invoice = data
}

func getPlays() map[string]interface{} {
	return plays
}

func setPlays(data map[string]interface{}) {
	plays = data
}

func totalAmount(performances []map[string]interface{}) int {
	result := 0
	for _, performance := range performances {
		result += amountFor(performance)
	}

	return result
}

func totalVolumeCredits(performances []map[string]interface{}) float64 {
	result := 0.0
	for _, performance := range performances {
		result += volumeCreditsFor(performance)
	}

	return result
}

func amountFor(performanceData map[string]interface{}) int {
	playType, ok := playFor(performanceData)["type"].(string)
	if !ok {
		return 0
	}

	if calculator := performance.CreatePerformanceCalculator(performanceData, playType); calculator != nil {
		return calculator.Amount()
	}

	return 0
}

func volumeCreditsFor(performanceData map[string]interface{}) float64 {
	playType, ok := playFor(performanceData)["type"].(string)
	if !ok {
		fmt.Println("no play type in performance")
		return 0.0
	}

	if calculator := performance.CreatePerformanceCalculator(performanceData, playType); calculator != nil {
		return calculator.VolumeCredits()
	}

	return 0.0
}

func playFor(performance map[string]interface{}) map[string]interface{} {
	playID, ok := performance["playID"].(string)
	if !ok {
		return nil
	}

	play, ok := getPlays()[playID].(map[string]interface{})
	if !ok {
		return nil
	}

	return play
}

func New(invoiceData, playsData map[string]interface{}) *statementData {
	setInvoice(invoiceData)
	setPlays(playsData)

	customer, ok := getInvoice()["customer"].(string)
	if !ok {
		return nil
	}

	performances, ok := getInvoice()["performances"].([]map[string]interface{})
	if !ok {
		return nil
	}

	return &statementData{
		customer:           customer,
		performances:       performances,
		totalAmount:        totalAmount(performances),
		totalVolumeCredits: totalVolumeCredits(performances),
	}
}
