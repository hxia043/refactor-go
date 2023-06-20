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

func volumeCreditsFor(performance map[string]interface{}) float64 {
	result := 0.0

	audience, ok := performance["audience"].(int)
	if !ok {
		fmt.Println("no audience in performance")
		return 0
	}

	result += math.Max(float64(audience-30), 0)

	playType, ok := playFor(performance)["type"].(string)
	if !ok {
		fmt.Println("no play type in performance")
		return 0
	}

	if playType == "comedy" {
		result += math.Floor(float64(audience) / 5)
	}

	return result
}

func usd(amount int) string {
	return fmt.Sprintf("%d", amount)
}

func totalVolumeCredits(performances []map[string]interface{}) float64 {
	result := 0.0
	for _, performance := range performances {
		result += volumeCreditsFor(performance)
	}

	return result
}

func totalAmount(performances []map[string]interface{}) int {
	result := 0
	for _, performance := range performances {
		result += amountFor(performance)
	}

	return result
}

func statement(invoice, plays map[string]interface{}) string {
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
		// print line for this order
		playID, ok := playFor(performance)["name"].(string)
		if !ok {
			return ""
		}

		audience, ok := performance["audience"].(int)
		if !ok {
			return ""
		}

		result += "    " + playID + ": " + usd(amountFor(performance)/100) + " " + fmt.Sprintf("%d", audience) + "seats" + "\n"
	}

	result += fmt.Sprintf("Amount owed is %s\n", usd(totalAmount(performances)/100))
	result += fmt.Sprintf("You earned %f credits", totalVolumeCredits(performances))
	return result
}

func main() {
	fmt.Println(statement(invoice[0], plays))
}
