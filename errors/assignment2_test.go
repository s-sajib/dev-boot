package errors

import (
	"fmt"
	"testing"
)

func TestAssignment2(t *testing.T) {
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println("------------------- Error Assignment 2 Test Start ---------------------")
	fmt.Println("-----------------------------------------------------------------------")
	type testCase struct {
		cost      float64
		recipient string
		expected  string
	}
	tests := []testCase{
		{1.4, "+1 (435) 555 0923", "SMS that costs $1.40 to be sent to '+1 (435) 555 0923' can not be sent"},
		{2.1, "+2 (702) 555 3452", "SMS that costs $2.10 to be sent to '+2 (702) 555 3452' can not be sent"},
	}
	if withSubmit {
		tests = append(tests, []testCase{{32.1, "+1 (801) 555 7456", "SMS that costs $32.10 to be sent to '+1 (801) 555 7456' can not be sent"}, {14.4, "+1 (234) 555 6545", "SMS that costs $14.40 to be sent to '+1 (234) 555 6545' can not be sent"}}...)
	}

	for _, test := range tests {
		if output := getSMSErrorString(test.cost, test.recipient); output != test.expected {
			t.Errorf(`Test Failed: (%v, %v) ->
	expected: %v
	actual: %v
`, test.cost, test.recipient, test.expected, output)
		} else {
			fmt.Printf(`Test Passed: (%v, %v) ->
	expected: %v
	actual: %v
`, test.cost, test.recipient, test.expected, output)
		}
	}
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println("------------------- Error Assignment 2 Test End ---------------------")
	fmt.Println("-----------------------------------------------------------------------")
}
