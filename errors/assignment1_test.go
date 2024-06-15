package errors

import (
	"fmt"
	"testing"
)

func TestAssignment1(t *testing.T) {

	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("------------------- Assignment 1 Test Start ---------------------")
	fmt.Println("-----------------------------------------------------------------")

	type testCase struct {
		msgToCustomer string
		msgToSpouse   string
		expectedCost  int
		expectedErr   error
	}
	tests := []testCase{
		{"Thanks for coming in to our flower shop today!", "We hope you enjoyed your gift.", 0, fmt.Errorf("can't send texts over 25 characters")},
		{"Thanks for joining us!", "Have a good day.", 76, nil},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{"Thank you.", "Enjoy!", 32, nil},
			{"We loved having you in!", "We hope the rest of your evening is fantastic.", 0, fmt.Errorf("can't send texts over 25 characters")},
		}...)
	}

	for _, test := range tests {
		cost, err := sendSMSToCouple(test.msgToCustomer, test.msgToSpouse)
		if cost != test.expectedCost || (err != nil && test.expectedErr != nil && err.Error() != test.expectedErr.Error()) {
			t.Errorf(`Test Failed: (%v, %v) ->
	expected: (%v, %v)
	actual: (%v, %v)
`, test.msgToCustomer, test.msgToSpouse, test.expectedCost, test.expectedErr, cost, err)
		} else {
			fmt.Printf(`Test Passed: (%v, %v) ->
	expected: (%v, %v)
	actual: (%v, %v)
`, test.msgToCustomer, test.msgToSpouse, test.expectedCost, test.expectedErr, cost, err)
		}
	}

	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("------------------- Assignment 1 Test End ---------------------")
	fmt.Println("-----------------------------------------------------------------")
}

var withSubmit = true
