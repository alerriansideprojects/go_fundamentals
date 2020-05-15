package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// Alternate way to move a variable for return
//  Expects that return will be a variable called sums
//  Allows for a little cleaner approach to returns and initializes sums early
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// numbers[1:] here is calling everything on a slice/array
			//  until the end of the slice/array
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return
}
