package main

import "fmt"

func filterEven(numbers []int) []int {
	result := []int{}

	for _, num := range numbers {
		if num%2 == 0 {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	var to_be_filtered = [14]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	//var filter = []int{}
	//filter = filterEven(to_be_filtered[:])

	fmt.Printf("Filtered items: %d ", filterEven(to_be_filtered[:]))
}
