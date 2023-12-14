package arrays_and_slices

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum

}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, slice := range numbersToSum {
		sums = append(sums, Sum(slice))
	}
	return sums
}

func SumTails(tailsToSum ...[]int) []int {
	var sums []int

	for _, slice := range tailsToSum {
		if len(slice) > 0 {
			sums = append(sums, Sum(slice[1:]))
		} else {
			sums = append(sums, 0)
		}
	}
	return sums
}
