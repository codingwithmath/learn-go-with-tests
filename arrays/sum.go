package main

func Sum(numbers []int) int {
	sum := 0

	// for i := 0; i < len(numbers); i++ {
	// 	sum += numbers[i]
	// }

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	//the above ... means that this is a variadic function that can take a variable number of arguments

	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
