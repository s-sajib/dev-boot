package slices

func getMessageCosts(messages []string) []float64 {
	messageLength := len(messages)
	costs := make([]float64, messageLength)

	for i := 0; i < messageLength; i++ {
		costs[i] = float64(len(messages[i])) * 0.01
	}
	return costs
}
