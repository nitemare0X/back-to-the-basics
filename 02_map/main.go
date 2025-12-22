package main

import (
	"fmt"
)

func mapper(numbers []int, transform func(int) int) []int {
	result := []int{}

	for _, num := range numbers {
		result = append(result, transform(num))
	}

	return result
}

func double(number int) int {
	return number * 2
}

func addOne(number int) int {
	return number + 1
}

func square(number int) int {
	return number * number
}

func main() {
	var toBeDoubled = []int{1, 2, 3, 4}
	var toAddOne = []int{5, 10, 15}
	var toBeSquared = []int{2, 4, 6}

	fmt.Println("Doubled items: ", mapper(toBeDoubled, double))
	fmt.Println("Add One: ", mapper(toAddOne, addOne))
	fmt.Println("Squared numbers: ", mapper(toBeSquared, square))
}
