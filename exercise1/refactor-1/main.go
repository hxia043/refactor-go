package main

import (
	"fmt"
	"statement/statement"
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

func main() {
	statementData := statement.New(invoice[0], plays)
	fmt.Println(statement.RenderPlainText(statementData))
}
