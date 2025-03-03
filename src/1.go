package main

import "fmt"

func main() {
	// Declare an array of integers
	numbers := []int{1, 2, 3, 4, 5}

	// Use a for loop to iterate over the array
	for i := range numbers {
		fmt.Println(numbers[i])
	}
}
