package main

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	return calculateBaseBill(costPerMessage, numMessages) * calculateDiscountRate(numMessages)
}

func calculateDiscountRate(messagesSent int) float64 {
	if messagesSent <= 1000 {
		return 1
	}
	if (1000 < messagesSent) && (messagesSent <= 5000) {
		return 0.9
	}
	if 5000 < messagesSent {
		return 0.8
	}
	return -1
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}
