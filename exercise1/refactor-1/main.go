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

func playFor(performance map[string]interface{}) map[string]interface{} {
	playID, ok := performance["playID"].(string)
	if !ok {
		return nil
	}

	play, ok := plays[playID].(map[string]interface{})
	if !ok {
		return nil
	}

	return play
}

func amountFor(performance map[string]interface{}) int {
	result := 0

	playType, ok := playFor(performance)["type"].(string)
	if !ok {
		return 0
	}

	switch playType {
	case "tragedy":
		result = 40000
		audience, ok := performance["audience"].(int)
		if !ok {
			return 0
		}

		if audience > 30 {
			result += 1000 * (audience - 30)
		}

		break
	case "comedy":
		result = 30000
		audience, ok := performance["audience"].(int)
		if !ok {
			return 0
		}
		if audience > 20 {
			result += 10000 + 500*(audience-20)
		}
		result += 300 * audience
		break
	default:
		return 0
	}

	return result
}

func statement(invoice, plays map[string]interface{}) string {
	totalAmount := 0
	volumeCredits := float64(0)

	customer, ok := invoice["customer"].(string)
	if !ok {
		return ""
	}
	result := "Statement for " + customer + "\n"

	performances, ok := invoice["performances"].([]map[string]interface{})
	if !ok {
		return ""
	}

	for _, performance := range performances {
		// add volume credits
		audience, ok := performance["audience"].(int)
		if !ok {
			return ""
		}
		volumeCredits += math.Max(float64(audience-30), 0)

		playType, ok := playFor(performance)["type"].(string)
		if !ok {
			return ""
		}
		if playType == "comedy" {
			volumeCredits += math.Floor(float64(audience) / 5)
		}

		// print line for this order
		playID, ok := playFor(performance)["name"].(string)
		if !ok {
			return ""
		}
		result += "    " + playID + ": " + fmt.Sprintf("%d", amountFor(performance)/100) + " " + fmt.Sprintf("%d", audience) + "seats" + "\n"
		totalAmount += amountFor(performance)
	}

	result += "Amount owed is " + fmt.Sprintf("%d", totalAmount/100) + "\n"
	result += "You earned " + fmt.Sprintf("%f", volumeCredits) + "credits"
	return result
}

func main() {
	fmt.Println(statement(invoice[0], plays))
}
