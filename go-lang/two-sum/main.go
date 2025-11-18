package main

import (
	"log"
)

// main function – entry point of the program
func main() {
	// Print a simple message to show what problem we are solving
	println("Solving problem... (1. Two Sum)")

	// Define an integer slice (array)
	num := []int{2, 11, 7, 15}

	// Target value we want to reach by adding two numbers from the array
	t := 9

	// Log the array and target value
	log.Println("Array:", num, "Target:", t)

	// Call the twoSome function to find the indices
	res := twoSome(num, t)

	// Log the result (the two indices that add up to target)
	log.Println("Indices of numbers that sum to", t, ":", res)
}

// twoSome finds two numbers in the array whose sum equals the target
func twoSome(nums []int, target int) []int {

	// Create a map to store numbers we've seen so far
	// The map's key:   the number
	// The map's value: the index of that number in the array
	m := make(map[int]int)

	// Loop through each number in the slice
	for i, num := range nums {

		// Calculate the complement (the number we need to reach the target)
		// Example: target=9, num=2 → complement=7
		c := target - num

		// Check if the complement already exists in the map
		// idx = index of the complement stored earlier
		// f   = boolean flag (true if found, false if not)
		if idx, f := m[c]; f {
			// If found, return the pair of indices
			// The complement's index (idx) and current index (i)
			return []int{idx, i}
		}

		// Otherwise, store the current number with its index in the map
		// This means: we have seen "num" at index "i"
		m[num] = i
	}

	// If no matching pair was found, return nil
	return nil
}
