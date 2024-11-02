package main

import "fmt"

func main() {
	// Create an empty slice
	aSlice := []float64{}
	// Both length and capacity are 0 because the slice is empty
	fmt.Println(aSlice, len(aSlice), cap(aSlice))

	// Add elements to a slice
	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, -34.0)
	fmt.Println(aSlice, "with length", len(aSlice))

	// A slice with a length of 4
	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4
	// t[4] = -5 this doesn't work because index is now out of range
	// Therefore, to add more items, we need to use append
	t = append(t, -5)
	fmt.Println(t)

	// A 2D slice
	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}
	// Visiting all elements of a 2D slice
	// with a double for loop
	for _, i := range twoD {
		for _, k := range i {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}

	// Creating a 2D slice variable named twoD
	make2D := make([][]int, 2)
	fmt.Println(make2D)
	make2D[0] = []int{1, 2, 3, 4}
	make2D[1] = []int{-1, -2, -3, -4}
	fmt.Println(make2D)
}
