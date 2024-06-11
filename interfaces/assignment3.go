/**
Implement the getExpenseReport function.

If the expense is an email then it should return the email's toAddress and the cost of the email.
If the expense is an sms then it should return the sms's toPhoneNumber and its cost.
If the expense has any other underlying type, just return an empty string and 0.0 for the cost.
*/

package interfaces

// func getExpenseReport(e expense) (string, float64) {
// 	isEmail, ok := e.(email)
// 	isSms, SmsOk := e.(sms)
// 	if ok {
// 		return isEmail.toAddress, isEmail.cost()
// 	} else if SmsOk {
// 		return isSms.toPhoneNumber, isSms.cost()
// 	}
// 	return "", 0.0
// }

//alternative approach using type switch
func getExpenseReport(e expense) (string, float64) {
	switch messageType := e.(type) {
	case email:
		return messageType.toAddress, e.cost()
	case sms:
		return messageType.toPhoneNumber, e.cost()
	default:
		return "", 0.0
	}
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}
