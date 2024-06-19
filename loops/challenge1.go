package loops

/**
Textio needs to divide messages up into data packets. Normally, packet sizes are uniform and divided
into bytes. An intern had the bright idea to customize packet sizes per message. Help textio test this
terrible, terrible idea.

ASSIGNMENT
Implement the getPacketSize function to take a string message and return an integer. Find the largest
packet size where the resulting number of packets is a composite number, ie: not 1 or prime. Use the
provided isPrime function within getPacketSize. If no such packet size exists, return 0.
**/

func getPacketSize(message string) int {
	lengthOfMessage := len(message)

	packetSize := 0

	for i := 1; i <= lengthOfMessage; i++ {

		if lengthOfMessage%i == 0 {

			result := lengthOfMessage / i

			if !isPrime(i) && i != 1 {
				if result > packetSize {
					packetSize = result
				}
			}
		}
	}
	return packetSize
}

func isPrime(num int) bool {
	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
