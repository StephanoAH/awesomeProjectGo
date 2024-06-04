package main

import "fmt"

func Sum(array []int) (sum int) {
	var total int
	for _, value := range array {
		total += value
	}
	return total
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return sums
}

func main() {
	//fmt.Printf("The total is = %d", Sum([]int{1, 2, 3, 4, 5}))
	fmt.Printf("The total is = %d", SumAll([]int{1, 2, 9}, []int{8, 2}))
}
