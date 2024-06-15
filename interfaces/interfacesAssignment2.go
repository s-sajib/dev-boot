/**
Complete the required methods so that the email type implements both the expense and
formatter interfaces.

COST()
If the email is not "subscribed", then the cost is 5 cents for each character in the body. If it is,
then the cost is 2 cents per character.

Return the total cost of the entire email in cents.

FORMAT()
The format method should return a string in this format:

'CONTENT' | Subscribed

If the email is not subscribed, change the second part to "Not Subscribed":

'CONTENT' | Not Subscribed

The single quotes are included in the string, and CONTENT is the email's body. For example:

'Hello, World!' | Subscribed

Note: you may want to import the fmt package and use Sprintf.
*/

package interfaces

import "fmt"

// func (e email) cost() int {
// 	length := len(e.body)
// 	if e.isSubscribed {
// 		return length * 2
// 	}
// 	return length * 5
// }

func (e email) format() string {
	if e.isSubscribed {
		return fmt.Sprintf("'%s' | Subscribed", e.body)
	}
	return fmt.Sprintf("'%s' | Not Subscribed", e.body)
}

// type expense interface {
// 	cost() int
// }

type formatter interface {
	format() string
}

// type email struct {
// 	isSubscribed bool
// 	body         string
// }
