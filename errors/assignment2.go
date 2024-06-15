package errors

import "fmt"

/**
We need better error logs for our backend developers to help them debug their code.

Complete the getSMSErrorString() function. It should return a string with this format:

SMS that costs $COST to be sent to 'RECIPIENT' can not be sent

COST is the cost of the SMS, always showing the price formatted to 2 decimal places.
RECIPIENT is the stringified representation of the recipient's phone number
Be sure to include the $ symbol and the single quotes
*/

func getSMSErrorString(cost float64, recipient string) string {
	formattedString := fmt.Sprintf("SMS that costs $%0.2f to be sent to '%v' can not be sent", cost, recipient)

	return formattedString
}
