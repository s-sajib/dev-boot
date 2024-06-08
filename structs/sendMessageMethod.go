package structs

/**
Create a SendMessage method for the User struct.

It should take a message string and messageLength int as inputs.

If the message length is within the user's message character limit,
return the original message and true (indicating the message can be sent),
otherwise, return an empty string and false.
*/

func (user User) SendMessage(message string, messageLength int) (string, bool) {
	if user.MessageCharLimit >= messageLength {
		return message, true
	}
	return "", false
}

// don't touch below this line

func newUser(name string, membershipType membershipType) User {
	membership := Membership{Type: membershipType}
	if membershipType == TypeStandard {
		membership.MessageCharLimit = 100
	} else if membershipType == TypePremium {
		membership.MessageCharLimit = 1000
	}
	return User{Name: name, Membership: membership}
}
