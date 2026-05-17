package main

import (
	"fmt"
)

func binary_search(array []int, itemToSearch int) (int, error) {
	low := 0
	high := len(array) - 1 // Example : 0 to 3 ( array length is 4)

	// 15, 30 , 45, 60
	// I want to search index where value is 45, so index 2
	for {
		mid := (low + high) / 2
		// Starting from mid value
		possibleGuess := array[mid]

		// If mid value is the item that i want to search, I return its index
		if itemToSearch == possibleGuess {
			// Returning the index where value is 45
			return mid, nil
		}

		// If mid value is lower that itemToSearch ( 30 and 45 for example)
		// then i want to see in other bisection
		if possibleGuess > itemToSearch {
			high = mid - 1
		} else {
			low = mid + 1
		}

		//  when pointers cross each other, the item is not present
		if low > high {
			return -1, fmt.Errorf("item %d not found in array", itemToSearch)
		}
	}
}

func main() {
	fmt.Println("--- Testing Binary Search with Error Handling ---")

	arr := []int{15, 30, 45, 60}

	// Test 1: Element present in the upper half
	result1, err1 := binary_search(arr, 45)
	if err1 != nil {
		fmt.Println("Test 1 Error:", err1)
	} else {
		fmt.Printf("Searching for 45. Expected index: 2 | Got: %d\n", result1)
	}

	// Test 2: Element present at the beginning
	result2, err2 := binary_search(arr, 15)
	if err2 != nil {
		fmt.Println("Test 2 Error:", err2)
	} else {
		fmt.Printf("Searching for 15. Expected index: 0 | Got: %d\n", result2)
	}

	// Test 3: Element not present
	result3, err3 := binary_search(arr, 100)
	if err3 != nil {
		fmt.Printf("Searching for 100. Expected index: -1 | Got Index: %d, Got Error: \"%v\"\n", result3, err3)
	} else {
		fmt.Printf("Searching for 100. Got index: %d\n", result3)
	}
}