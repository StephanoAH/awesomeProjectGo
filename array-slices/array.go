package main

import "fmt"

func Sum(array []int) (sum int) {
	var total int
	for _, value := range array {
		total += value
	}
	return total
}

func SumAll() []int {
	var total int

}

func main() {
	fmt.Printf("The total is = %d", Sum([]int{1, 2, 3, 4, 5}))
}
