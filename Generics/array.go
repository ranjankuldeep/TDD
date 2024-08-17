package main

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	sum := func(acc, x int) int {
		return acc + x
	}
	return Reduce(numbers, sum, 0)
}

// SumAllTails calculates the sums of all but the first number given a collection of slices.
func SumAllTails(numbersToSum ...[]int) []int {
	sumTails := func(acc, x []int) []int {
		if len(x) == 0 {
			acc = append(acc, 0)
		} else {
			tail := x[1:]
			acc = append(acc, Sum(tail))
		}
		return acc
	}
	return Reduce(numbersToSum, sumTails, []int{})
}

// Generic reduce function
func Reduce[T any](collection []T, f func(T, T) T, intialValue T) T {
	var result = intialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}
