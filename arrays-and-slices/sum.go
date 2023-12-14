package arrays_and_slices

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum

}

func SumAll(numbersToSum ...[]int) []int {
	length := len(numbersToSum)
	sums := make([]int, length)

	for i, slice := range numbersToSum {
		sums[i] = Sum(slice)
	}
	return sums
}
