package main

import "fmt"

func filter(numbers []int, condition func(int) bool) []int {
	result := []int{}

	for _, num := range numbers {
		if condition(num) {
			result = append(result, num)
		}
	}

	return result
}

func isDivisibleBy3(number int) bool {
	//if number%3 == 0 {
	//	return true
	//}
	//return false
	// alternatively for ease
	return number%3 == 0
}

func isEven(number int) bool {
	return number%2 == 0
}

func main() {
	var to_be_filtered = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	fmt.Printf("Filtered items: %d ", filter(to_be_filtered, isDivisibleBy3))
	fmt.Println("Even numbers:", filter(to_be_filtered, isEven))
}
