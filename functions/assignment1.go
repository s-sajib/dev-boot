/**
## Instructions:

### BILLING_SYSTEM
Textio is making some changes to how they bill users for bulk messages.
The system should now calculate the bill based on:

The number of messages sent
The cost per message
Thresholds for discounts

ASSIGNMENT
Complete the calculateFinalBill and the calculateDiscount
functions.

#### CALCULATE_FINAL_BILL

Use the calculateBaseBill function to get the cost for the
messages sent. Then, use the calculateDiscount function to get
the discount rate. Finally, calculate the final bill by
applying the discount to the base bill and return the result.

The discount is a percentage represented by a float. 10% = .1

CALCULATE_DISCOUNT
This function should take the number of messages sent, and
return the relevant discount based on the following discount
rates:

10% for more than 1000 messages.
20% for more than 5000 messages.
0% for anything less.
*/

package functions

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	baseCost := calculateBaseBill(costPerMessage, numMessages)
	discountRate := calculateDiscount(numMessages)
	finalBill := baseCost * (1 - discountRate)
	return finalBill
}

func calculateDiscount(messagesSent int) float64 {

	if messagesSent > 5000 {
		return .2
	}
	if messagesSent > 1000 {
		return .1
	}
	return 0.0
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}
