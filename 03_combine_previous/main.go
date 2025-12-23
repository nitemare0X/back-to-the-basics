package main

import (
	"fmt"
)

func filter(numbers []int, condition func(int) bool) []int {
	result := []int{}

	for _, num := range numbers {
		if condition(num) {
			result = append(result, num)
		}
	}

	return result
}

func isEven(number int) bool {
	return number%2 == 0
}

func mapper(numbers []int, transform func(int) int) []int {
	result := []int{}

	for _, num := range numbers {
		result = append(result, transform(num))
	}

	return result
}

func square(number int) int {
	return number * number
}

func main() {
	var toBeFiltered = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Output: ", mapper(filter(toBeFiltered, isEven), square))
}
