package statement

import "fmt"

func usd(amount int) string {
	return fmt.Sprintf("%d", amount)
}

func RenderPlainText(data *statementData) string {
	result := "Statement for " + data.customer + "\n"

	for _, performance := range data.performances {
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

	result += fmt.Sprintf("Amount owed is %s$\n", usd(data.totalAmount/100))
	result += fmt.Sprintf("You earned %f credits", data.totalVolumeCredits)
	return result
}

func RenderHTML(data statementData) string {
	result := "<h1>Statement for " + data.customer + "</h1>\n"

	for _, performance := range data.performances {
		// print line for this order
		playID, ok := playFor(performance)["name"].(string)
		if !ok {
			return ""
		}

		audience, ok := performance["audience"].(int)
		if !ok {
			return ""
		}

		result += "<td>    " + playID + ": " + usd(amountFor(performance)/100) + " " + fmt.Sprintf("%d", audience) + "seats" + "</td>\n"
	}

	result += fmt.Sprintf("<p>Amount owed is %s</p>\n", usd(data.totalAmount/100))
	result += fmt.Sprintf("<p>You earned %f credits</p>", data.totalVolumeCredits)

	return result
}
