package interfaces

// import (
// 	"fmt"
// 	"testing"
// )

// func TestAssignment2(t *testing.T) {
// 	fmt.Println("-------------------------------------------------------------------")
// 	fmt.Println("------------------- Testing Assignment 2 Interface  ---------------")
// 	fmt.Println("-------------------------------------------------------------------")
// 	type testCase struct {
// 		body           string
// 		isSubscribed   bool
// 		expectedCost   int
// 		expectedFormat string
// 	}
// 	tests := []testCase{
// 		{"hello there", true, 22, "'hello there' | Subscribed"},
// 		{"general kenobi", false, 70, "'general kenobi' | Not Subscribed"},
// 	}
// 	if withSubmit {
// 		tests = append(tests, []testCase{
// 			{"i hate sand", true, 22, "'i hate sand' | Subscribed"},
// 			{"it's coarse and rough and irritating", false, 180, "'it's coarse and rough and irritating' | Not Subscribed"},
// 			{"and it gets everywhere", true, 44, "'and it gets everywhere' | Subscribed"},
// 		}...)
// 	}

// 	for _, test := range tests {
// 		e := email{
// 			body:         test.body,
// 			isSubscribed: test.isSubscribed,
// 		}
// 		cost := e.cost()
// 		format := e.format()
// 		if format != test.expectedFormat || cost != test.expectedCost {
// 			t.Errorf(
// 				`Test Failed: (%v, %v) ->
// 	expected: (%v, %v)
// 	actual: (%v, %v)
// `,
// 				test.body,
// 				test.isSubscribed,
// 				test.expectedCost,
// 				test.expectedFormat,
// 				cost,
// 				format,
// 			)
// 		} else {
// 			fmt.Printf(
// 				`Test Passed: (%v, %v) ->
// 	expected: (%v, %v)
// 	actual: (%v, %v)
// `,
// 				test.body,
// 				test.isSubscribed,
// 				test.expectedCost,
// 				test.expectedFormat,
// 				cost,
// 				format,
// 			)
// 		}
// 	}
// }
