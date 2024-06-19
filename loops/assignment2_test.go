package loops

import (
	"fmt"
	"testing"
)

func TestAssignment2(t *testing.T) {
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("------------------- Assignment 2 Test Start ---------------------")
	fmt.Println("-----------------------------------------------------------------")
	type testCase struct {
		costMultiplier   float64
		maxCostInPennies int
		expected         int
	}
	tests := []testCase{
		{1.1, 5, 17},
		{1.3, 10, 9},
		{1.35, 25, 11},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{1.2, 15, 15},
			{1.3, 20, 12},
		}...)
	}

	for _, test := range tests {
		if output := getMaxMessagesToSend(test.costMultiplier, test.maxCostInPennies); output != test.expected {
			t.Errorf("Test Failed: (%v, %v) -> expected: %v actual: %v", test.costMultiplier, test.maxCostInPennies, test.expected, output)
		} else {
			fmt.Printf("Test Passed: (%v, %v) -> expected: %v actual: %v\n", test.costMultiplier, test.maxCostInPennies, test.expected, output)
		}
	}
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("------------------- Assignment 2 Test End ---------------------")
	fmt.Println("-----------------------------------------------------------------")
}
