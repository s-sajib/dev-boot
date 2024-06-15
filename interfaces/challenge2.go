package interfaces

/**
Implement the importance methods for each message type. They should return the importance score for each message type.

For a directMessage the importance score is based on if the message isUrgent or not. If it is urgent,
the importance score is 50 otherwise the importance score is equal to the DM's priorityLevel.

For a groupMessage the importance score is based on the messages priorityLevel

All systemAlert messages should return a 100 importance score.

Complete the processNotification function. It should identify the type and return different values for each type

For a directMessage, return the sender's username and importance score.
For a groupMessage, return the group's name and the importance score.
For an systemAlert, return the alert code and the importance score.
If the notification does not match any known type, return an empty string and a score of 0.
*/

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	}
	return d.priorityLevel
}
func (g groupMessage) importance() int {

	return g.priorityLevel
}
func (s systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch messageType := n.(type) {
	case directMessage:
		return messageType.senderUsername, messageType.importance()
	case groupMessage:
		return messageType.groupName, messageType.importance()
	case systemAlert:
		return messageType.alertCode, messageType.importance()
	default:
		return "", 0
	}

}
