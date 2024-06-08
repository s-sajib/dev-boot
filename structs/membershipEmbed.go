/**
Create a new struct called Membership, it should have:

A Type field with the pre-defined membershipType string type.
A MessageCharLimit integer field.
Update the User struct to embed a Membership.

Complete the newUser function. It should return a new User with all the fields set as you
would expect based on the inputs. If the user is premium, the MessageCharLimit should be 1000,
otherwise, it should only be 100.
*/

package structs

type membershipType string

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

// don't touch above this line

type Membership struct {
	Type             membershipType
	MessageCharLimit int
}

type User struct {
	Name string
	Membership
}

// func newUser(name string, membershipType membershipType) User {
// 	newMember := User{name, Membership{membershipType, 100}}
// 	if membershipType == TypePremium {
// 		newMember.Membership.MessageCharLimit = 1000
// 	}
// 	return newMember

// }
