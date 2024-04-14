package main

import "fmt"

func main() {
	// Creating an empty slice
	var slice []int

	// Appending elements to the slice
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 6)
	slice = append(slice, -5)

	// Printing the slice
	fmt.Println("Slice:", slice) // Output: Slice: [1 2 3]
}
