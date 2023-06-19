package main

import (
	"fmt"
	"math"
)

var invoice = []map[string]interface{}{
	{
		"customer": "BigCo",
		"performances": []map[string]interface{}{
			{
				"playID":   "hamlet",
				"audience": 55,
			},
			{
				"playID":   "as-like",
				"audience": 35,
			},
			{
				"playID":   "othello",
				"audience": 40,
			},
		},
	},
}

var plays = map[string]interface{}{
	"hamlet": map[string]interface{}{
		"name": "Hamlet",
		"type": "tragedy",
	},
	"as-like": map[string]interface{}{
		"name": "As You like It",
		"type": "comedy",
	},
	"othello": map[string]interface{}{
		"name": "Othello",
		"type": "tragedy",
	},
}

func statement(invoice, plays map[string]interface{}) string {
	totalAmount := 0
	volumeCredits := float64(0)

	customerValue, ok := invoice["customer"].(string)
	if !ok {
		return ""
	}
	result := "Statement for " + customerValue + "\n"

	performancesValue, ok := invoice["performances"].([]map[string]interface{})
	if !ok {
		return ""
	}

	for _, perf := range performancesValue {
		playID, ok := perf["playID"].(string)
		if !ok {
			return ""
		}

		playValue, ok := plays[playID].(map[string]interface{})
		thisAmount := 0

		switch playValue["type"].(string) {
		case "tragedy":
			thisAmount = 40000
			audience, ok := perf["audience"].(int)
			if !ok {
				return ""
			}

			if audience > 30 {
				thisAmount += 1000 * (audience - 30)
			}

			break
		case "comedy":
			thisAmount = 30000
			audience, ok := perf["audience"].(int)
			if !ok {
				return ""
			}
			if audience > 20 {
				thisAmount += 10000 + 500*(audience-20)
			}
			thisAmount += 300 * audience
			break
		default:
			return "unknown error"
		}

		// add volume credits
		audience, ok := perf["audience"].(int)
		if !ok {
			return ""
		}
		volumeCredits += math.Max(float64(audience-30), 0)

		playType, ok := playValue["type"].(string)
		if !ok {
			return ""
		}

		if playType == "comedy" {
			volumeCredits += math.Floor(float64(audience) / 5)
		}

		// print line for this order
		result += playID + ": " + fmt.Sprintf("%d", thisAmount/100) + " " + fmt.Sprintf("%d", audience) + "seats" + "\n"
		totalAmount += thisAmount
	}

	result += "Amount owed is " + fmt.Sprintf("%d", totalAmount/100) + "\n"
	result += "You earned " + fmt.Sprintf("%f", volumeCredits) + "credits" + "\n"
	return result
}

func main() {
	fmt.Println(statement(invoice[0], plays))
}
