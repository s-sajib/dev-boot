/**
We've been asked to "bucket" costs per day, in a given period.

Complete the getCostsByDay() function using the append() function. It accepts a slice of cost structs and returns a slice of float64s, where each float64 represents the total cost for a day.

The length of the daily costs slice should be the number of days contained in the costs slice, up to and including the last day. There should only be one "bucket" of costs per day. Be sure to include entries for intermediate days, even if they don't have any costs.

Days are numbered and start at 0.


Given this input:

[]cost{
    {0, 4.0},
    {1, 2.1},
    {5, 2.5},
    {1, 3.1},
}

[]float64{
    4.0, // first day
    5.2, // 2.1 + 3.1
    0.0, // intermediate days with no costs
    0.0, // ...
    0.0, // ...
    2.5, // last day
}

*/

package slices

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	length := len(costs)

	maxDayNum := 0

	for i := 0; i < length; i++ {
		if costs[i].day > maxDayNum {
			maxDayNum = costs[i].day
		}
	}

	result := make([]float64, maxDayNum+1)

	for i := 0; i < length; i++ {
		result[costs[i].day] = result[costs[i].day] + costs[i].value
	}
	return result
}
